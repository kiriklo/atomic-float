// Atomic Float implemetation using mutex.
// For performance comparison.

package atomic_float

import (
	"sync"
)

type atomicFloat32Mutex struct {
	mu  sync.RWMutex
	f32 float32
}

func newAtomicFloat32Mutex(initial float32) *atomicFloat32Mutex {
	return &atomicFloat32Mutex{f32: initial}
}

// Add attempts to add delta to the value stored in the atomic float and return
// the new value.
func (a *atomicFloat32Mutex) add(delta float32) float32 {
	a.mu.Lock()
	a.f32 += delta
	new := a.f32
	a.mu.Unlock()
	return new
}

// Load atomically loads the current atomic float value.
func (a *atomicFloat32Mutex) load() float32 {
	a.mu.RLock()
	f := a.f32
	a.mu.RUnlock()
	return f
}

// Store atomically stores new into the atomic float.
func (a *atomicFloat32Mutex) store(new float32) {
	a.mu.Lock()
	a.f32 = new
	a.mu.Unlock()
}

// Swap atomically stores new and returns the previous value.
func (a *atomicFloat32Mutex) swap(new float32) float32 {
	a.mu.Lock()
	old := a.f32
	a.f32 = new
	a.mu.Unlock()
	return old
}

// CAS is an atomic compare-and-swap. It returns true if the current value equals old.
func (a *atomicFloat32Mutex) cas(old, new float32) bool {
	a.mu.Lock()
	var e bool
	if a.f32 == old {
		a.f32 = new
		e = true
	} else {
		e = false
	}
	a.mu.Unlock()
	return e
}

type atomicFloat64Mutex struct {
	mu  sync.RWMutex
	f64 float64
}

func newAtomicFloat64Mutex(initial float64) *atomicFloat64Mutex {
	return &atomicFloat64Mutex{f64: initial}
}

// Add attempts to add delta to the value stored in the atomic float and return
// the new value.
func (a *atomicFloat64Mutex) add(delta float64) float64 {
	a.mu.Lock()
	a.f64 += delta
	new := a.f64
	a.mu.Unlock()
	return new
}

// Load atomically loads the current atomic float value.
func (a *atomicFloat64Mutex) load() float64 {
	a.mu.RLock()
	f := a.f64
	a.mu.RUnlock()
	return f
}

// Store atomically stores new into the atomic float.
func (a *atomicFloat64Mutex) store(new float64) {
	a.mu.Lock()
	a.f64 = new
	a.mu.Unlock()
}

// Swap atomically stores new and returns the previous value.
func (a *atomicFloat64Mutex) swap(new float64) float64 {
	a.mu.Lock()
	old := a.f64
	a.f64 = new
	a.mu.Unlock()
	return old
}

// CAS is an atomic compare-and-swap. It returns true if the current value equals old.
func (a *atomicFloat64Mutex) cas(old, new float64) bool {
	a.mu.Lock()
	var e bool
	if a.f64 == old {
		a.f64 = new
		e = true
	} else {
		e = false
	}
	a.mu.Unlock()
	return e
}
