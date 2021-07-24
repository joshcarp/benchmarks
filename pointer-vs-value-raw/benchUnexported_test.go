/* Note: This is originally from @err0r500 https://github.com/err0r500/golang-pointer-vs-value-benchmark */
package pointer_vs_value_raw

import "testing"

func BenchmarkUnexportedNewValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = newValue()
	}
}

func BenchmarkUnexportedNewPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = newPointer()
	}
}

func BenchmarkUnexportedFuncValue(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncValue(big)
	}
}

func BenchmarkUnexportedFuncPointer(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncPointer(big)
	}
}

func BenchmarkUnexportedMethodValue(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodValue()
	}
}

func BenchmarkUnexportedMethodPointer(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodPointer()
	}
}

func BenchmarkUnexportedFuncValueNoRef(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncNoRefValue(big)
	}
}

func BenchmarkUnexportedFuncPointerNoRef(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dummyFuncNoRefPointer(big)
	}
}

func BenchmarkUnexportedMethodValueNoRef(b *testing.B) {
	big := newValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodNoRefValue()
	}
}

func BenchmarkUnexportedMethodPointerNoRef(b *testing.B) {
	big := newPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.dummyMethodNoRefPointer()
	}
}
