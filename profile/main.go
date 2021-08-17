package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"

	"github.com/lemon-mint/slowtable"
)

var keys = func() []string {
	tmp := make([]string, 8192)
	for i := 0; i < 8192; i++ {
		tmp[i] = strconv.Itoa(i)
	}
	return tmp
}()

var keysBytes = func() [][]byte {
	tmp := make([][]byte, 8192)
	for i := 0; i < 8192; i++ {
		tmp[i] = []byte(strconv.Itoa(i))
	}
	return tmp
}()

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	table := slowtable.NewTable(nil, 65535)
	table.PoolPreload()
	for i := 0; i < 8192; i++ {
		table.SetS(keys[i], nil)
	}

	for j := 0; j < 1024*40; j++ {
		for i := 0; i < 8192; i++ {
			table.GetS(keys[i])
		}
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
