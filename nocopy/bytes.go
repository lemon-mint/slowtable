package nocopy

import (
	"reflect"
	"unsafe"
)

func StringToBytes(s string) []byte {
	var bytes []byte
	*(*string)(unsafe.Pointer(&bytes)) = s
	(*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Cap = len(s)
	return bytes
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Alias
var S2B = StringToBytes
var B2S = BytesToString
