package main


import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	//"sort"
	"strconv"
	"strings"
)


func main() {
	// Intake File
	file, err := os.Open("day_two_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	sc := bufio.NewScanner(file)
	lines := make([]string, 0)


	// Each line will be an item in the array.
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}


	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}


	safe := 0
	not_safe := 0

	
	// Takes a line, which is a string array and puts it in the holder array via strings.Split().
	for _, j := range lines {
		var holder []string
		var holder_int []int
		holder = strings.Split(j, " ")


		// DEBUG Strings
		//fmt.Printf("Holder: %T\n", holder[1])
		//fmt.Printf("Holder: %d\n", holder[1])
		//fmt.Println(holder)


		// Iterates through holder array and converts item to integer.
		// Appends integer to holder_int array.
		for _, num := range holder {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			holder_int = append(holder_int, i)
		}


		// DEBUG Strings
		//fmt.Printf("HolderInt: %T\n", holder_int[1])
		//fmt.Printf("HolderInt: %d\n", holder_int[1])
		//fmt.Println(holder_int)


		// Tracking variables for when comparing two numbers in holder_int array.
		previous_number := 0
		increase := 0
		decrease := 0
		no_change := 0
		too_much_change := 0


		for _, j := range holder_int {
			if previous_number == 0 {
				previous_number = j
			} else {
				absolute_change := math.Abs(float64(previous_number - j))

				if absolute_change <= 3 {
					if previous_number == j {
						no_change += 1
					} else if previous_number < j {
						previous_number = j
						increase += 1
					} else if previous_number > j {
						previous_number = j
						decrease += 1
					}
				} else {
					too_much_change += 1
				}
			}
		}


		// Checks if report is safe or unsafe.
		if increase == 0 && no_change == 0 && too_much_change == 0 && decrease > 0 {
			safe += 1
		} else if decrease == 0 && no_change == 0 && too_much_change == 0 && increase > 0 {
			safe += 1
		} else {
			not_safe += 1
		}


	}


	fmt.Printf("Number of SAFE reports: %d\n", safe)
	fmt.Printf("Number of UNSAFE reports: %d\n", not_safe)
}


// TO-DO
// - How to resolve tolerating a single bad level.
