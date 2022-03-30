package map_vs_switch

import (
	"testing"
)

type Foo struct {
	a          []byte
	b          string
	c          int
	d, e, f, g, h []byte
}

func (f *Foo) MutateFoo() {
	f.b = string(f.a)
	f.c = len(f.b)
}

func BenchmarkMutateFoo(b *testing.B) {
	f := Foo{a: []byte("15236781902312312152367819023123121523678190231231215236781902312312")}
	for i := 0; i < b.N; i++ {
		f.MutateFoo()
	}
}

func (f Foo) DontMutateFoo() Foo {
	f.b = string(f.a)
	f.c = len(f.b)
	return f
}

func BenchmarkDontMutateFoo(b *testing.B) {
	f := Foo{a: []byte("15236781902312312152367819023123121523678190231231215236781902312312")}
	for i := 0; i < b.N; i++ {
		f = f.DontMutateFoo()
	}
}
