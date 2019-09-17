package main

/*
#include "skynet.h"
extern int gos_cb(struct skynet_context* context, void* ptr, int typ, int session, int source, char* msg, int sz);

static int gos_cb_wrapper(struct skynet_context * context, void *ud, int type, int session, uint32_t source , const void * msg, size_t sz) {
	gos_cb(context, ud, type, session, (int)source, (char*)msg, (int)sz);
}

static inline void set_gos_callback(struct skynet_context * context, uintptr_t ud) {
	skynet_callback(context, (void*)ud, gos_cb_wrapper);
}
*/
import "C"

import (
	"log"
	"sync"
	"sync/atomic"
	"unsafe"
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

//export gos_cb
func gos_cb(context *C.struct_skynet_context, ptr unsafe.Pointer, typ C.int, session C.int, source C.int, s *C.char, sz C.int) C.int {
	msg := C.GoStringN(s, sz)
	log.Printf("gos cb: from %d, session: %d, msg:%s", source, session, msg)
	return 0
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
	C.set_gos_callback(ctx, C.uintptr_t(ptr))
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
