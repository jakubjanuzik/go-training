package main

var name = "Jan"

func main4() {
	var languages [5]string

	languages[0] = "GO"
	languages[1] = "GO"
	languages[2] = "POWER"
	languages[3] = "RANGERS"

	for index, language := range languages {
		println(index, language)
	}

	for _, language := range languages {
		println(language)
	}

	numbers := [4]int{10, 20, 30, 40}

	for index := 0; index < len(numbers); index++ {
		if numbers[index] < 30 {
			println(numbers[index])
		}

		if number := numbers[index]; number < 30 {
			println(number)
		}
	}
}
