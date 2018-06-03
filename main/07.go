package main

func swap(text, otherText string) (string, string) {
	return otherText, text
}

func getTime() (time string) {
	time = "13:43"
	return
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func execute(fn func()) {
	fn()
}

func create() func() {
	return func() {
		println("Anonymous Function")
	}
}

func main7() {
	var text, otherText = swap("1", "2")
	println(text, otherText)

	greet := func() {
		println("hello")
	}

	execute(greet)

	myFn := create()
	myFn()

	execute(create())

	var rest = func() string {
		println("SIF")
		return "result"
	}()

	println(rest)

}
