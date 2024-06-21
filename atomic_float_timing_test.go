package atomic_float

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAddFloat32(b *testing.B) {
	var x Float32
	var y float32 = 2.5
	for i := 0; i < b.N; i++ {
		x.Add(y)
	}
}

func BenchmarkAddFloat32_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	var y float32 = 2.5
	for i := 0; i < b.N; i++ {
		mf.add(y)
	}
}

func BenchmarkAddFloat32Parallel(b *testing.B) {
	var x Float32
	var delta float32 = 2.5
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Add(delta)
		}
	})
}

func BenchmarkAddFloat32Parallel_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	var delta float32 = 2.5
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.add(delta)
		}
	})
}

func BenchmarkAddInt32Parallel(b *testing.B) {
	var x atomic.Int32
	var delta int32 = 1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Add(delta)
		}
	})
}

func BenchmarkStoreFloat32(b *testing.B) {
	var x Float32
	for i := 0; i < b.N; i++ {
		x.Store(float32(i))
		if res := x.Load(); res != float32(i) {
			b.Errorf("Expected %v, got %v", float32(i), res)
		}
	}
}

func BenchmarkStoreFloat32_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	for i := 0; i < b.N; i++ {
		mf.store(float32(i))
		if res := mf.load(); res != float32(i) {
			b.Errorf("Expected %v, got %v", float32(i), res)
		}
	}
}

func BenchmarkStoreFloat32Parallel(b *testing.B) {
	var x Float32
	var delta float32 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Store(delta)
		}
	})
}

func BenchmarkStoreFloat32Parallel_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	var delta float32 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.store(delta)
		}
	})
}

func BenchmarkSwapFloat32(b *testing.B) {
	var x Float32
	for i := 1; i < b.N; i++ {
		if result := x.Swap(float32(i)); result != float32(i-1) {
			b.Errorf("Expected %v, got %v", float32(i-1), result)
		}
		if res := x.Load(); res != float32(i) {
			b.Errorf("Expected %v, got %v", float32(i), res)
		}
	}
}

func BenchmarkSwapFloat32_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	for i := 1; i < b.N; i++ {
		if result := mf.swap(float32(i)); result != float32(i-1) {
			b.Errorf("Expected %v, got %v", float32(i-1), result)
		}
		if res := mf.load(); res != float32(i) {
			b.Errorf("Expected %v, got %v", float32(i), res)
		}
	}
}

func BenchmarkSwapFloat32Parallel(b *testing.B) {
	var x Float32
	var delta float32 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Swap(delta)
		}
	})
}

func BenchmarkSwapFloat32Parallel_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	var delta float32 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.swap(delta)
		}
	})
}

func BenchmarkCASFloat32(b *testing.B) {
	var x Float32
	for i := 1; i < b.N; i++ {
		if result := x.CompareAndSwap(float32(i-1), float32(i)); result != true {
			b.Errorf("Expected %v, got %v", true, result)
		}
		if res := x.Load(); res != float32(i) {
			b.Errorf("Expected %v, got %v", float32(i), res)
		}
	}
}

func BenchmarkCASFloat32_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)

	for i := 1; i < b.N; i++ {
		if result := mf.cas(float32(i-1), float32(i)); result != true {
			b.Errorf("Expected %v, got %v", true, result)
		}
		if res := mf.load(); res != float32(i) {
			b.Errorf("Expected %v, got %v", float32(i), res)
		}
	}
}

func BenchmarkCASFloat32Parallel(b *testing.B) {
	var x Float32
	var delta float32 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.CompareAndSwap(0.0, delta)
		}
	})
}

func BenchmarkCASFloat32Parallel_Mutex(b *testing.B) {
	var x float32
	mf := newAtomicFloat32Mutex(x)
	var delta float32 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.cas(0.0, delta)
		}
	})
}
func BenchmarkAddFloat64(b *testing.B) {
	var x Float64
	var y float64 = 2.5
	for i := 0; i < b.N; i++ {
		x.Add(y)
	}
}

func BenchmarkAddFloat64_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	var y float64 = 2.5
	for i := 0; i < b.N; i++ {
		mf.add(y)
	}
}

func BenchmarkAddFloat64Parallel(b *testing.B) {
	var x Float64
	var delta float64 = 2.5
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Add(delta)
		}
	})
}

func BenchmarkAddFloat64Parallel_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	var delta float64 = 2.5
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.add(delta)
		}
	})
}

func BenchmarkStoreFloat64(b *testing.B) {
	var x Float64
	for i := 0; i < b.N; i++ {
		x.Store(float64(i))
		if res := x.Load(); res != float64(i) {
			b.Errorf("Expected %v, got %v", float64(i), res)
		}
	}
}

func BenchmarkStoreFloat64_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	for i := 0; i < b.N; i++ {
		mf.store(float64(i))
		if res := mf.load(); res != float64(i) {
			b.Errorf("Expected %v, got %v", float64(i), res)
		}
	}
}

func BenchmarkStoreFloat64Parallel(b *testing.B) {
	var x Float64
	var delta float64 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Store(delta)
		}
	})
}

func BenchmarkStoreFloat64Parallel_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	var delta float64 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.store(delta)
		}
	})
}

func BenchmarkSwapFloat64(b *testing.B) {
	var x Float64
	for i := 1; i < b.N; i++ {
		if result := x.Swap(float64(i)); result != float64(i-1) {
			b.Errorf("Expected %v, got %v", float64(i-1), result)
		}
		if res := x.Load(); res != float64(i) {
			b.Errorf("Expected %v, got %v", float64(i), res)
		}
	}
}

func BenchmarkSwapFloat64_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	for i := 1; i < b.N; i++ {
		if result := mf.swap(float64(i)); result != float64(i-1) {
			b.Errorf("Expected %v, got %v", float64(i-1), result)
		}
		if res := mf.load(); res != float64(i) {
			b.Errorf("Expected %v, got %v", float64(i), res)
		}
	}
}

func BenchmarkSwapFloat64Parallel(b *testing.B) {
	var x Float64
	var delta float64 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.Swap(delta)
		}
	})
}

func BenchmarkSwapFloat64Parallel_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	var delta float64 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.swap(delta)
		}
	})
}

func BenchmarkCASFloat64(b *testing.B) {
	var x Float64
	for i := 1; i < b.N; i++ {
		if result := x.CompareAndSwap(float64(i-1), float64(i)); result != true {
			b.Errorf("Expected %v, got %v", true, result)
		}
		if res := x.Load(); res != float64(i) {
			b.Errorf("Expected %v, got %v", float64(i), res)
		}
	}
}

func BenchmarkCASFloat64_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)

	for i := 1; i < b.N; i++ {
		if result := mf.cas(float64(i-1), float64(i)); result != true {
			b.Errorf("Expected %v, got %v", true, result)
		}
		if res := mf.load(); res != float64(i) {
			b.Errorf("Expected %v, got %v", float64(i), res)
		}
	}
}

func BenchmarkCASFloat64Parallel(b *testing.B) {
	var x Float64
	var delta float64 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x.CompareAndSwap(0.0, delta)
		}
	})
}

func BenchmarkCASFloat64Parallel_Mutex(b *testing.B) {
	var x float64
	mf := newAtomicFloat64Mutex(x)
	var delta float64 = 1.1
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mf.cas(0.0, delta)
		}
	})
}
