package main

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main8() {

	increment := wrapper()

	println(increment())
	println(increment())

	increment2 := wrapper()

	println(increment())
	println(increment2())
}
