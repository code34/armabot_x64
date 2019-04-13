package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
	"bytes"
	"encoding/json"
	"net/http"
)

//export RVExtensionVersion
func RVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("Version 1.0")
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export RVExtensionArgs
func RVExtensionArgs(output *C.char, outputsize C.size_t, input *C.char, argv **C.char, argc C.int) {
	var offset = unsafe.Sizeof(uintptr(0))
	var out []string
	for index := C.int(0); index < argc; index++ {
		out = append(out, C.GoString(*argv))
		argv = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + offset))
	}
	temp := fmt.Sprintf("Function: %s nb params: %d params: %s!", C.GoString(input), argc,  out)

	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export RVExtension
func RVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	parameters := strings.Split(fmt.Sprintf(C.GoString(input)), ";")
	temp := execWH(parameters)

	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

/* Execute Webhook
https://discordapp.com/developers/docs/resources/webhook#get-webhook
*/
func execWH(parameters []string) string {
	type Query struct {
		Varname string `json:"content"`
		Varvalue string `json:"username"`
	}

	url := parameters[0]
	content := parameters[1]
	username := parameters [2]

	u := Query{Varname: content, Varvalue: username}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	_, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return "false"
	} else {
		return "true"
	}
}

func main() {}