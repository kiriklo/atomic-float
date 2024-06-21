package atomic_float

import (
	"math"
	"math/big"
	"runtime"
	"testing"
)

func TestConvert(t *testing.T) {
	var f float64 = 10.0123
	bf := big.NewRat(1, 1)
	bf.SetFloat64(f)
}

// TestAddFloat32_Positive checks if adding a positive number to a Float32 works as expected.
func TestAddFloat32_Positive(t *testing.T) {
	var f Float32
	f.Add(1.2)
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	f.Add(2.3)
	if result := f.Load(); result != 3.5 {
		t.Errorf("Expected %v, got %v", 3.5, result)
	}

}

// TestAddFloat32_Neg checks if adding a negative number to a Float32 works as expected.
func TestAddFloat32_Negative(t *testing.T) {
	var f Float32
	f.Add(-1.2)
	if result := f.Load(); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	f.Add(-2.3)
	if result := f.Load(); result != -3.5 {
		t.Errorf("Expected %v, got %v", -3.5, result)
	}
}

// TestAddFloat32_Zero checks if adding zero to a Float32 returns the same value.
func TestAddFloat32_Zero(t *testing.T) {
	var f Float32
	f.Add(0)
	if result := f.Load(); result != 0 {
		t.Errorf("Expected %v, got %v", 0, result)
	}
}

// TestAddFloat32_Large checks if adding a large number to Float32 works as expected does not overflow.
func TestAddFloat32_Large(t *testing.T) {
	var f Float32
	var largeNum float32 = 1e6
	f.Add(largeNum)
	if result := f.Load(); result != largeNum {
		t.Errorf("Expected %v, got %v", largeNum, result)
	}
}

// TestAddFloat32_Small checks if adding a small number to a Float32 works as expected and does not underflow.
func TestAddFloat32_Small(t *testing.T) {
	var f Float32
	var smallNum float32 = 1e-6
	f.Add(smallNum)
	if result := f.Load(); result != smallNum {
		t.Errorf("Expected %v, got %v", smallNum, result)
	}
}

// TestAdd32_Infinity if adding infinity to a Float32.
func TestAddFloat32_Infinity(t *testing.T) {
	var f Float32
	inf := float32(math.Inf(1))
	f.Add(inf)
	if result := f.Load(); result != inf {
		t.Errorf("Expected %v, got %v", inf, result)
	}
}

func TestStoreFloat32_Positive(t *testing.T) {
	var f Float32
	f.Store(1.2)
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	f.Store(2.3)
	if result := f.Load(); result != 2.3 {
		t.Errorf("Expected %v, got %v", 2.3, result)
	}
}

func TestStoreFloat32_Negative(t *testing.T) {
	var f Float32
	f.Store(-1.2)
	if result := f.Load(); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	f.Store(-2.3)
	if result := f.Load(); result != -2.3 {
		t.Errorf("Expected %v, got %v", -2.3, result)
	}
}

func TestSwapFloat32_Positive(t *testing.T) {
	var f Float32
	if result := f.Swap(1.2); result != 0.0 {
		t.Errorf("Expected %v, got %v", 0.0, result)
	}
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.Swap(2.3); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.Load(); result != 2.3 {
		t.Errorf("Expected %v, got %v", 2.3, result)
	}
}

func TestSwapFloat32_Negative(t *testing.T) {
	var f Float32
	if result := f.Swap(-1.2); result != 0.0 {
		t.Errorf("Expected %v, got %v", 0.0, result)
	}
	if result := f.Load(); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	if result := f.Swap(-2.3); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	if result := f.Load(); result != -2.3 {
		t.Errorf("Expected %v, got %v", -2.3, result)
	}
}

func TestCompareAndSwapFloat32_Positive(t *testing.T) {
	var f Float32
	if result := f.CompareAndSwap(0.0, 1.2); result != true {
		t.Errorf("Expected %v, got %v", true, result)
	}
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.CompareAndSwap(0.0, 1.2); result != false {
		t.Errorf("Expected %v, got %v", false, result)
	}
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.CompareAndSwap(1.2, 2.3); result != true {
		t.Errorf("Expected %v, got %v", true, result)
	}
	if result := f.Load(); result != 2.3 {
		t.Errorf("Expected %v, got %v", 2.3, result)
	}
}

func TestAddFloat32Concurrent(t *testing.T) {
	const itemsCount = 2000
	const gorotines = 30
	var f Float32
	var delta float32 = 2.5

	done := make(chan bool)
	for i := 0; i < gorotines; i++ {
		go func() {
			for j := 0; j < itemsCount; j++ {
				f.Add(delta)
			}
			done <- true
		}()
	}
	for i := 0; i < gorotines; i++ {
		<-done
	}
	if result := f.Load(); result != float32(itemsCount*gorotines)*delta {
		t.Errorf("Expected %v, got %v", float32(itemsCount*gorotines)*delta, result)
	}
	runtime.GC()
}

// TestAddFloat64_Positive checks if adding a positive number to a Float64 works as expected.
func TestAddFloat64_Positive(t *testing.T) {
	var f Float64
	f.Add(1.2)
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	f.Add(2.3)
	if result := f.Load(); result != 3.5 {
		t.Errorf("Expected %v, got %v", 3.5, result)
	}

}

// TestAddFloat64_Neg checks if adding a negative number to a Float64 works as expected.
func TestAddFloat64_Negative(t *testing.T) {
	var f Float64
	f.Add(-1.2)
	if result := f.Load(); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	f.Add(-2.3)
	if result := f.Load(); result != -3.5 {
		t.Errorf("Expected %v, got %v", -3.5, result)
	}
}

// TestAddFloat64_Zero checks if adding zero to a Float64 returns the same value.
func TestAddFloat64_Zero(t *testing.T) {
	var f Float64
	f.Add(0)
	if result := f.Load(); result != 0 {
		t.Errorf("Expected %v, got %v", 0, result)
	}
}

// TestAddFloat64_Large checks if adding a large number to Float64 works as expected does not overflow.
func TestAddFloat64_Large(t *testing.T) {
	var f Float64
	var largeNum float64 = 1e12
	f.Add(largeNum)
	if result := f.Load(); result != largeNum {
		t.Errorf("Expected %v, got %v", largeNum, result)
	}
}

// TestAddFloat64_Small checks if adding a small number to a Float64 works as expected and does not underflow.
func TestAddFloat64_Small(t *testing.T) {
	var f Float64
	var smallNum float64 = 1e-12
	f.Add(smallNum)
	if result := f.Load(); result != smallNum {
		t.Errorf("Expected %v, got %v", smallNum, result)
	}
}

// TestAdd32_Infinity if adding infinity to a Float32.
func TestAddFloat64_Infinity(t *testing.T) {
	var f Float64
	inf := float64(math.Inf(1))
	f.Add(inf)
	if result := f.Load(); result != inf {
		t.Errorf("Expected %v, got %v", inf, result)
	}
}

func TestStoreFloat64_Positive(t *testing.T) {
	var f Float64
	f.Store(1.2)
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	f.Store(2.3)
	if result := f.Load(); result != 2.3 {
		t.Errorf("Expected %v, got %v", 2.3, result)
	}
}

func TestStoreFloat64_Negative(t *testing.T) {
	var f Float64
	f.Store(-1.2)
	if result := f.Load(); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	f.Store(-2.3)
	if result := f.Load(); result != -2.3 {
		t.Errorf("Expected %v, got %v", -2.3, result)
	}
}

func TestSwapFloat64_Positive(t *testing.T) {
	var f Float64
	if result := f.Swap(1.2); result != 0.0 {
		t.Errorf("Expected %v, got %v", 0.0, result)
	}
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.Swap(2.3); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.Load(); result != 2.3 {
		t.Errorf("Expected %v, got %v", 2.3, result)
	}
}

func TestSwapFloat64_Negative(t *testing.T) {
	var f Float64
	if result := f.Swap(-1.2); result != 0.0 {
		t.Errorf("Expected %v, got %v", 0.0, result)
	}
	if result := f.Load(); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	if result := f.Swap(-2.3); result != -1.2 {
		t.Errorf("Expected %v, got %v", -1.2, result)
	}
	if result := f.Load(); result != -2.3 {
		t.Errorf("Expected %v, got %v", -2.3, result)
	}
}

func TestCompareAndSwapFloat64_Positive(t *testing.T) {
	var f Float64
	if result := f.CompareAndSwap(0.0, 1.2); result != true {
		t.Errorf("Expected %v, got %v", true, result)
	}
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.CompareAndSwap(0.0, 1.2); result != false {
		t.Errorf("Expected %v, got %v", false, result)
	}
	if result := f.Load(); result != 1.2 {
		t.Errorf("Expected %v, got %v", 1.2, result)
	}
	if result := f.CompareAndSwap(1.2, 2.3); result != true {
		t.Errorf("Expected %v, got %v", true, result)
	}
	if result := f.Load(); result != 2.3 {
		t.Errorf("Expected %v, got %v", 2.3, result)
	}
}

func TestAddFloat64Concurrent(t *testing.T) {
	const itemsCount = 10000
	const gorotines = 10
	var f Float64
	var delta float64 = 2.5

	done := make(chan bool)
	for i := 0; i < gorotines; i++ {
		go func() {
			for j := 0; j < itemsCount; j++ {
				f.Add(delta)
			}
			done <- true
		}()
	}
	for i := 0; i < gorotines; i++ {
		<-done
	}
	if result := f.Load(); result != float64(itemsCount*gorotines)*delta {
		t.Errorf("Expected %v, got %v", float64(itemsCount*gorotines)*delta, result)
	}
	runtime.GC()
}
