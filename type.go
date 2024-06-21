package atomic_float

// Compatable with src/runtime/internal/atomic/types.go

// An Float32 is an atomic float32. The zero value is zero.
type Float32 struct {
	_ noCopy
	v float32
}

// Load atomically loads and returns the value stored in x.
//
//go:nosplit
func (x *Float32) Load() float32 { return LoadFloat32(&x.v) }

// Store atomically stores val into x.
//
//go:nosplit
func (x *Float32) Store(val float32) { StoreFloat32(&x.v, val) }

// Swap atomically stores new into x and returns the previous value.
//
//go:nosplit
func (x *Float32) Swap(new float32) (old float32) { return SwapFloat32(&x.v, new) }

// CompareAndSwap executes the compare-and-swap operation for x.
//
//go:nosplit
func (x *Float32) CompareAndSwap(old, new float32) (swapped bool) {
	return CompareAndSwapFloat32(&x.v, old, new)
}

// Add atomically adds delta to x and returns the new value.
//
//go:nosplit
func (x *Float32) Add(delta float32) (new float32) { return AddFloat32(&x.v, delta) }

// Float64 is an atomically accessed float64 value.
//
// 8-byte aligned on all platforms, unlike a regular float64.
//
// An Float64 must not be copied.
type Float64 struct {
	noCopy noCopy
	_      align64
	v      float64
}

// Load atomically loads and returns the value stored in x.
//
//go:nosplit
func (x *Float64) Load() float64 { return LoadFloat64(&x.v) }

// Store atomically stores val into x.
//
//go:nosplit
func (x *Float64) Store(val float64) { StoreFloat64(&x.v, val) }

// Swap atomically stores new into x and returns the previous value.
//
//go:nosplit
func (x *Float64) Swap(new float64) (old float64) { return SwapFloat64(&x.v, new) }

// CompareAndSwap executes the compare-and-swap operation for x.
//
//go:nosplit
func (x *Float64) CompareAndSwap(old, new float64) (swapped bool) {
	return CompareAndSwapFloat64(&x.v, old, new)
}

// Add atomically adds delta to x and returns the new value.
//
//go:nosplit
func (x *Float64) Add(delta float64) (new float64) { return AddFloat64(&x.v, delta) }

// Copied from src/runtime/internal/atomic/types.go

// noCopy may be added to structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
//
// Note that it must not be embedded, due to the Lock and Unlock methods.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// align64 may be added to structs that must be 64-bit aligned.
// This struct is recognized by a special case in the compiler
// and will not work if copied to any other package.
type align64 struct{}
