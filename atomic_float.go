package atomic_float

// See src/runtime/internal/atomic/types.go

//go:nosplit
//go:noinline
func LoadFloat32(ptr *float32) float32 {
	return *ptr
}

//go:noescape
func AddFloat32(ptr *float32, delta float32) float32

//go:noescape
func StoreFloat32(ptr *float32, delta float32)

//go:noescape
func SwapFloat32(ptr *float32, delta float32) float32

//go:noescape
func CompareAndSwapFloat32(ptr *float32, old float32, new float32) bool

//go:nosplit
//go:noinline
func LoadFloat64(ptr *float64) float64 {
	return *ptr
}

//go:noescape
func AddFloat64(ptr *float64, delta float64) float64

//go:noescape
func StoreFloat64(ptr *float64, delta float64)

//go:noescape
func SwapFloat64(ptr *float64, delta float64) float64

//go:noescape
func CompareAndSwapFloat64(ptr *float64, old float64, new float64) bool
