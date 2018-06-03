package main

func main5() {
	switch name {
	case "Jan":
		println("hello")
		// fallthrough <-- domyślnie jest z breakiem
	case "Maria":
		println("Hello")
	default:
		println("Default")
	}
}
