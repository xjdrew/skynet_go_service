package main

// #include "skynet.h"
import "C"

import (
	"log"
	"sync"
	"sync/atomic"
)

type gosEnv struct {
	handle uint32
	args   string
}

var (
	nextHandle uint32
	handleMap  sync.Map
)

func getEnv(ptr uintptr) *gosEnv {
	env, ok := handleMap.Load(ptr)
	if !ok {
		return nil
	}
	return env.(*gosEnv)
}

//export gos_create
func gos_create() uintptr {
	handle := atomic.AddUint32(&nextHandle, 1)
	env := &gosEnv{
		handle: handle,
	}
	ptr := uintptr(handle)
	handleMap.Store(ptr, env)
	log.Print("gos create:", handle)
	return ptr
}

//export gos_init
func gos_init(ptr uintptr, ctx *C.struct_skynet_context, args *C.char) {
	env := getEnv(ptr)
	env.args = C.GoString(args)
	log.Print("gos init:", env.handle, env.args)
	return
}

//export gos_release
func gos_release(ptr uintptr) {
	env := getEnv(ptr)
	if env == nil {
		return
	}

	log.Print("gos create:", env.handle)
	handleMap.Delete(ptr)
}

//export gos_signal
func gos_signal(ctx uintptr, signal C.int) {
}

func main() {}
