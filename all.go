//age on mars
package main

import "fmt"

func mars_age(age int) int {
	age *= 365
	age /= 687
	return age
}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}



//debug & fix
package main

import "fmt"

func main() {
	// outputs GO 3 times
	fmt.Println("GO")
	fmt.Println("GO")
	fmt.Println("GO")
}



//downloader 
package main

import "fmt"

//define the download() function

func download(s int, ch chan int) {
	sum := 0
	for i := 1; i <= s; i++ {
		sum += i
	}
	ch <- sum
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var s1, s2, s3 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	go download(s1, ch1)
	go download(s2, ch2)
	go download(s3, ch3)

	//output the sum of all results
	fmt.Println(<-ch1 + <-ch2 + <-ch3)
}



//match results
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



//say the numbers
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		var x int
		fmt.Scanln(&x)
		if x >= 0 && x <= 10 {
			switch x {
			case 0:
				fmt.Print("Zero")
			case 1:
				fmt.Print("One")
			case 2:
				fmt.Print("Two")
			case 3:
				fmt.Print("Three")
			case 4:
				fmt.Print("Four")
			case 5:
				fmt.Print("Five")
			case 6:
				fmt.Print("Six")
			case 7:
				fmt.Print("Seven")
			case 8:
				fmt.Print("Eight")
			case 9:
				fmt.Print("Nine")
			case 10:
				fmt.Print("Ten")
			default:
				fmt.Println("This is not a number between 0 and 10")
			}
		}
	}
}



//ticking timer
package main

import "fmt"

type Timer struct {
	value string
	id    int
}

func (x *Timer) tick() {
	x.id++
	fmt.Println(x.id)
}

func main() {
	var x int
	fmt.Scanln(&x)

	t := Timer{"timer1", 0}

	for i := 0; i < x; i++ {
		t.tick()
	}
}
