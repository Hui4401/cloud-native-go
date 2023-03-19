package reflection

func ExampleLaw2() {
	Law2()
	// Output:
	// They are same!
}

func ExampleLaw3Fail() {
	// Law3Fail() // panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	// Output:
	//
}

func ExampleLaw3Success() {
	Law3Success()
	// Output: x = 7.1
}
