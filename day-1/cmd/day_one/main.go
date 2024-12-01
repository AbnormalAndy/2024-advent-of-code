package main


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
	// Intake File
	file, err := os.Open("day_one_input.txt")
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


	var list_one []string
	var list_two []string


	var list_one_number []int
	var list_two_number []int


	// Split the numbers in the line into different arrays.
	for _, j := range lines {
		var holder []string
		holder = strings.Split(j, "   ")
		//fmt.Println(holder[2])
		list_one = append(list_one, holder[0])
		list_two = append(list_two, holder[1])
	}


	// Convert string arrays into integer arrays.
	for _, j := range list_one {
		num, err := strconv.Atoi(j)

		if err != nil {
			fmt.Println("Error converting string:", num)
		} else {
			list_one_number = append(list_one_number, num)
		}
	}


	// Convert string arrays into integer arrays.
	for _, j := range list_two {
		num, err := strconv.Atoi(j)

		if err != nil {
			fmt.Println("Error converting string:", num)
		} else {
			list_two_number = append(list_two_number, num)
		}
	}


	// Debug information to check type and if number was sorted.
	fmt.Printf("L1: %T\n", list_one[2])
	fmt.Printf("L1N: %T\n", list_one_number[2])
	fmt.Printf("L1NS: %d\n", list_one_number[2])
	fmt.Printf("L2: %T\n", list_two[2])
	fmt.Printf("L2N: %T\n", list_two_number[2])
	fmt.Printf("L2NS: %d\n", list_two_number[2])


	// Sort array from lowest to highest.
	sort.Slice(list_one_number, func(i, j int) bool {
		return list_one_number[i] < list_one_number[j]
	})


	// Sort array from lowest to highest.
	sort.Slice(list_two_number, func(i, j int) bool {
		return list_two_number[i] < list_two_number[j]
	})


	// Debug information to check type and if number was sorted.
	fmt.Printf("L1: %T\n", list_one[2])
	fmt.Printf("L1N: %T\n", list_one_number[2])
	fmt.Printf("L1NS: %d\n", list_one_number[2])
	fmt.Printf("L2: %T\n", list_two[2])
	fmt.Printf("L2N: %T\n", list_two_number[2])
	fmt.Printf("L2NS: %d\n", list_two_number[2])
	

	sum_part_one := 0


	// Compare numbers in each array and calculate difference.
	for i, j := range list_one_number {
		if j != list_two_number[i] && j > list_two_number[i] {
			total := j - list_two_number[i]
			sum_part_one += total
		} else if j != list_two_number[i] && j < list_two_number[i] {
			total := list_two_number[i] - j
			sum_part_one += total
		} else {
			sum_part_one += 0
		}
	}


	fmt.Println("Part One Sum:", sum_part_one)


	freq := make(map[int]int)


	// Count frequency of number in list one appearing in list two.
	// Adds number and frequency of that number to a map.
	for _, j := range list_one_number {
		for _, num := range list_two_number {
			if j == num {
				freq[j] = freq[j] + 1
			}
		}
	}


	sum_part_two := 0


	// Takes key from map (the number) times the value in the map (the frequency).
	for k, v := range freq {
		total := k * v
		sum_part_two += total
	}


	fmt.Println("Part Two Sum:", sum_part_two)
}
