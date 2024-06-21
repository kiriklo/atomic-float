# atomic-float

Package atomic-float provides an implementation of an atomic floating point number in Go.
The package exposes a high-level API similar to the standard sync/atomic package, allowing to load and store values atomically without needing to manage memory alignment or unsafe pointers directly.

Interesting part is that there is no Float64 type in sync/atomic package, but Float64 type exists in
src/runtime/internal/atomic/types.go and it looks like this:

```GO
// Float64 is an atomically accessed float64 value.
//
// 8-byte aligned on all platforms, unlike a regular float64.
//
// A Float64 must not be copied.
type Float64 struct {
	// Inherits noCopy and align64 from Uint64.
	u Uint64
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (f *Float64) Load() float64 {
	r := f.u.Load()
	return *(*float64)(unsafe.Pointer(&r))
}

// Store updates the value atomically.
//
//go:nosplit
func (f *Float64) Store(value float64) {
	f.u.Store(*(*uint64)(unsafe.Pointer(&value)))
}
```

So as you can see it implements only Load() and Store().

Performance benchmarks show that this implementation may outperform others in certain scenarios,
especially when working with large CPU numbers.
For example performance comparison with locking float using mutex (see mutex_float.go):

Results for 8 cores, amd64:

Add operations:
Parallel add is difficult to implement for float64 so it's slower than mutex.
```
BenchmarkAddFloat32-8                          280463023                4.288 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat32_Mutex-8                    150758780                7.927 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat32Parallel-8                   13970506                87.95 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat32Parallel_Mutex-8             21598959                53.17 ns/op            0 B/op          0 allocs/op
```

Store operations:
~5x times faster than mutex.
```
BenchmarkStoreFloat32-8                        435084490                2.754 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat32_Mutex-8                   97964332                11.94 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat32Parallel-8                 72069214                17.23 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat32Parallel_Mutex-8           23748750                51.77 ns/op            0 B/op          0 allocs/op
```

Swap operations:
~2x times faster than mutex.
```
BenchmarkSwapFloat32-8                         367457792                3.319 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat32_Mutex-8                    99337452                11.85 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat32Parallel-8                  74272150                16.01 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat32Parallel_Mutex-8            34221607                34.11 ns/op            0 B/op          0 allocs/op
```

Comare and swap operations:
~2x times faster than mutex.
```
BenchmarkCASFloat32-8                  351014774                3.426 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat32_Mutex-8             99425815                11.91 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat32Parallel-8           75087339                15.90 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat32Parallel_Mutex-8     36394281                33.53 ns/op            0 B/op          0 allocs/op
```

Here results for float32 16 cores, amd64:
```
BenchmarkAddFloat32-16                         128748429                9.163 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat32_Mutex-16                    51499935                20.37 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat32Parallel-16                  27619159                49.33 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat32Parallel_Mutex-16            20564847                55.89 ns/op            0 B/op          0 allocs/op

BenchmarkStoreFloat32-16                       122806131                9.146 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat32_Mutex-16                  36710494                30.19 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat32Parallel-16                73159482                17.12 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat32Parallel_Mutex-16          22394076                55.67 ns/op            0 B/op          0 allocs/op

BenchmarkSwapFloat32-16                        128837408                9.206 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat32_Mutex-16                   35437849                30.69 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat32Parallel-16                 79185967                16.19 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat32Parallel_Mutex-16           23610846                51.16 ns/op            0 B/op          0 allocs/op

BenchmarkCASFloat32-16                         125738266                9.631 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat32_Mutex-16                    38472019                30.52 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat32Parallel-16                  81562186                17.25 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat32Parallel_Mutex-16            23268123                51.63 ns/op            0 B/op          0 allocs/op
```

Results for float64 16 cores, amd64:
```
BenchmarkAddFloat64-16                         129843850                9.414 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat64_Mutex-16                    50157577                22.27 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat64Parallel-16                  24768976                53.35 ns/op            0 B/op          0 allocs/op
BenchmarkAddFloat64Parallel_Mutex-16            21684469                54.17 ns/op            0 B/op          0 allocs/op

BenchmarkStoreFloat64-16                       132773221                9.021 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat64_Mutex-16                  38601219                29.91 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat64Parallel-16                84785281                15.80 ns/op            0 B/op          0 allocs/op
BenchmarkStoreFloat64Parallel_Mutex-16          20736561                53.34 ns/op            0 B/op          0 allocs/op

BenchmarkSwapFloat64-16                        129053419                9.344 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat64_Mutex-16                   38986734                30.65 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat64Parallel-16                 72557970                16.08 ns/op            0 B/op          0 allocs/op
BenchmarkSwapFloat64Parallel_Mutex-16           22597039                51.21 ns/op            0 B/op          0 allocs/op

BenchmarkCASFloat64-16                         128451848                9.275 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat64_Mutex-16                    38984960                30.23 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat64Parallel-16                  88792860                16.05 ns/op            0 B/op          0 allocs/op
BenchmarkCASFloat64Parallel_Mutex-16            19706762                52.48 ns/op            0 B/op          0 allocs/op
```