//go:build ignore

package test

import "fmt"

func OnExitPrintf1(call fmt.CallContext, n int, err error) {
	println("Exiting hook1....")
	call.SetReturnVal(0, 1024)
	v := call.GetData().(int)
	println(v)
}

type any = interface{}

func OnEnterPrintf1(call fmt.CallContext, format string, arg ...any) {
	println("Entering hook1....")
	call.SetData(555)
	call.SetParam(0, "olleH%s\n")
	p1 := call.GetParam(1).([]any)
	p1[0] = "goodcatch"
}

func OnEnterPrintf2(call fmt.CallContext, format interface{}, arg ...interface{}) {
	println("hook2")
	for i := 0; i < 10; i++ {
		if i == 5 {
			panic("deliberately")
		}
	}
}

func onEnterSprintf1(call fmt.CallContext, format string, arg ...any) {
	print("a1")
}

func onExitSprintf1(call fmt.CallContext, s string) {
	print("b1")
}

func onEnterSprintf2(call fmt.CallContext, format string, arg ...any) {
	print("a2")
	_ = call.IsSkipCall()
}

func onExitSprintf2(call fmt.CallContext, s string) {
	println("b2")
}

func onEnterSprintf3(call fmt.CallContext, format string, arg ...any) {
	println("a3")
}

func onExitSprintf3(call fmt.CallContext, s string) {
	print("b3")
}
