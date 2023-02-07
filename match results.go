package main

import "fmt"

func main() {

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var a int

	lastMatchResult := ""

	fmt.Scanln(&lastMatchResult)

	results = append(results, lastMatchResult)

	for _, v := range results {

		switch {

		case v == "w":

			a += 3

		case v == "l":

			a += 0

		case v == "d":

			a += 1

		}

	}

	fmt.Println(a)

}
