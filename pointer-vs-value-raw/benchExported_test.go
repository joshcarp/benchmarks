/* Note: This is originally from @err0r500 https://github.com/err0r500/golang-pointer-vs-value-benchmark */
package pointer_vs_value_raw

import "testing"

func BenchmarkNewValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewValue()
	}
}

func BenchmarkNewPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewPointer()
	}
}

func BenchmarkFuncValue(b *testing.B) {
	big := NewValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DummyFuncValue(big)
	}
}

func BenchmarkFuncPointer(b *testing.B) {
	big := NewPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DummyFuncPointer(big)
	}
}

func BenchmarkMethodValue(b *testing.B) {
	big := NewValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.DummyMethodValue()
	}
}

func BenchmarkMethodPointer(b *testing.B) {
	big := NewPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.DummyMethodPointer()
	}
}

func BenchmarkFuncValueNoRef(b *testing.B) {
	big := NewValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DummyFuncNoRefValue(big)
	}
}

func BenchmarkFuncPointerNoRef(b *testing.B) {
	big := NewPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DummyFuncNoRefPointer(big)
	}
}

func BenchmarkMethodValueNoRef(b *testing.B) {
	big := NewValue()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.DummyMethodNoRefValue()
	}
}

func BenchmarkMethodPointerNoRef(b *testing.B) {
	big := NewPointer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.DummyMethodNoRefPointer()
	}
}
