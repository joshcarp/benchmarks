/* Note: This is originally from @err0r500 https://github.com/err0r500/golang-pointer-vs-value-benchmark */
package pointer_vs_value_raw /*  */

type MyBigStruct struct {
	f0  int64
	f1  int64
	f2  int64
	f3  int64
	f4  int64
	f5  int64
	f6  int64
	f7  int64
	f8  int64
	f9  int64
	f10 int64
	f11 int64
	f12 int64
	f13 int64
	f14 int64
	f15 int64
	f16 int64
	f17 int64
	f18 int64
	f19 int64
}

func NewValue() MyBigStruct {
	return MyBigStruct{}
}

func NewPointer() *MyBigStruct {
	return &MyBigStruct{}
}

func DummyFuncValue(mbs MyBigStruct) {
	_ = mbs
}

func DummyFuncPointer(mbs *MyBigStruct) {
	_ = mbs
}

func (mbs MyBigStruct) DummyMethodValue() {
	_ = mbs
}

func (mbs *MyBigStruct) DummyMethodPointer() {
	_ = mbs
}

func DummyFuncNoRefValue(mbs MyBigStruct) {
	_ = mbs
}

func DummyFuncNoRefPointer(mbs *MyBigStruct) {
	_ = mbs
}

func (mbs MyBigStruct) DummyMethodNoRefValue() {
	_ = mbs
}

func (mbs *MyBigStruct) DummyMethodNoRefPointer() {
	_ = mbs
}
