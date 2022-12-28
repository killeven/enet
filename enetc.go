package enet

// #cgo CFLAGS: -Ienet/include/
// #cgo linux LDFLAGS: -Lenet/ -lenet
// #cgo windows LDFLAGS: -Lenet/ -lenet -lWs2_32 -lWinmm
// #include <enet/enet.h>
import "C"

// Initialize enet
func Initialize() {
	C.enet_initialize()
}

// Deinitialize enet
func Deinitialize() {
	C.enet_deinitialize()
}
