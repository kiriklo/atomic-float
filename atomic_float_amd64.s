#include "textflag.h"

// See src/runtime/internal/atomic/atomic_amd64.s

// Works but slow
// float32 AddFloat32(ptr *float32, delta float32)
// Atomically:
//	*ptr += delta;
//	return *ptr;
TEXT ·AddFloat32(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), BX
    MOVL    delta+8(FP), X0

    loop:
    MOVL    0(BX), AX
    MOVL    AX, X1
    ADDSS   X0, X1
    MOVL    X1, CX
    LOCK
    CMPXCHGL    CX, 0(BX)
    SETEQ   AX
    CMPB    AX, $1
    JNE loop
    MOVL    CX, ret+16(FP)
    RET

// float32 StoreFloat32(ptr *float32, new float32)
// Atomically:
//	*ptr = new;
TEXT ·StoreFloat32(SB), NOSPLIT, $0-20
  	MOVQ	ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    XCHGL   AX, 0(BX)
    RET

// float32 SwapFloat32(ptr *float32, new float32)
// Atomically:
//	old := *ptr;
//	*ptr = new;
//	return old;
TEXT ·SwapFloat32(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), BX
    MOVL    delta+8(FP), AX
    XCHGL   AX, 0(BX)
    MOVQ	AX, ret+16(FP)
    RET

// bool CompareAndSwapFloat32(float32 *val, float32 old, float32 new)
// Atomically:
//	if(*val == old){
//		*val = new;
//		return 1;
//	} else
//		return 0;
TEXT ·CompareAndSwapFloat32(SB),NOSPLIT,$0-17
	MOVQ	ptr+0(FP), BX
	MOVL	old+8(FP), AX
	MOVL	new+12(FP), CX
	LOCK
	CMPXCHGL	CX, 0(BX)
	SETEQ   AX
    MOVL    AX, ret+16(FP)
	RET

// Works but slow
// float64 AddFloat64(ptr *float64, delta float64)
// Atomically:
//	*ptr += delta;
//	return *ptr;
TEXT ·AddFloat64(SB), NOSPLIT, $0-24
    MOVQ	ptr+0(FP), BX
    MOVQ    delta+8(FP), X0

    loop:
    MOVQ    0(BX), AX
    MOVQ    AX, X1
    ADDSD   X0, X1
    MOVQ    X1, CX
    LOCK
    CMPXCHGQ    CX, 0(BX)
    SETEQ   AX
    CMPB    AX, $1
    JNE loop
    MOVQ    CX, ret+16(FP)
    RET

// float64 StoreFloat64(ptr *float64, new float64)
// Atomically:
//	*ptr = new;
TEXT ·StoreFloat64(SB), NOSPLIT, $0-24
    MOVQ	ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    XCHGQ   AX, 0(BX)
    RET

// float64 SwapFloat64(ptr *float64, new float64)
// Atomically:
//	old := *ptr;
//	*ptr = new;
//	return old;
TEXT ·SwapFloat64(SB), NOSPLIT, $0-24
    MOVQ	ptr+0(FP), BX
    MOVQ    delta+8(FP), AX
    XCHGQ   AX, 0(BX)
    MOVQ	AX, ret+16(FP)
    RET

// bool CompareAndSwapFloat64(float64 *val, float64 old, float64 new)
// Atomically:
//	if(*val == old){
//		*val = new;
//		return 1;
//	} else
//		return 0;
TEXT ·CompareAndSwapFloat64(SB),NOSPLIT,$0-25
    MOVQ	ptr+0(FP), BX
	MOVQ	old+8(FP), AX
	MOVQ    new+16(FP), CX
	LOCK
	CMPXCHGQ    CX, 0(BX)
	SETEQ   AX
    MOVQ    AX, ret+24(FP)
	RET
