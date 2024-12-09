package main

import (
	"fmt"
	"math"
)


func safeOrUnsafe(report []int) bool {
	previous_number := 0
	pp_number := 0
	increase := 0
	decrease := 0
	no_change := 0
	tm_change := 0


	for _, i := range report {
		if previous_number == 0 {
			pp_number = i
			previous_number = i
		} else {
			absolute := math.Abs(float64(previous_number - i))
			//absolute_pp := math.Abs(float64(pp_number - i))


			if previous_number < i {
				if pp_number > i {
					previous_number = pp_number
					decrease += 1
				} else if absolute > 3 {
					pp_number = previous_number
					previous_number = i
					tm_change += 1
				} else {
					pp_number = previous_number
					previous_number = i
					increase += 1
				}
			} else if previous_number > i {
				if pp_number < i {
					previous_number = pp_number
					increase += 1
				} else if absolute > 3 {
					pp_number = previous_number
					previous_number = i
					tm_change += 1
				} else {
					pp_number = previous_number
					previous_number = i
					decrease += 1
				}
			} else {
				if pp_number != i {
					previous_number = pp_number
					if previous_number < i {
						if pp_number > i {
							previous_number = pp_number
							decrease += 1
						} else if absolute > 3 {
							pp_number = previous_number
							previous_number = i
							tm_change += 1
						} else {
							pp_number = previous_number
							previous_number = i
							increase += 1
						}
					} else if previous_number > i {
						if pp_number < i {
							previous_number = pp_number
							increase += 1
						} else if absolute > 3 {
							pp_number = previous_number
							previous_number = i
							tm_change += 1
						} else {
							pp_number = previous_number
							previous_number = i
							decrease += 1
						}
					}
				} else {
					pp_number = previous_number
					previous_number = i
					no_change += 1
				}
			}
		}


	}


	if increase == len(report) - 1 || decrease == len(report) - 1 {
		fmt.Println(len(report))
		fmt.Println(increase)
		fmt.Println(decrease)
		return true
	} else {
		fmt.Println(len(report))
		fmt.Println(increase)
		fmt.Println(decrease)
		return false
	}

}


func main() {


	report1 := []int{7, 6, 4, 2, 1}
	report2 := []int{1, 2, 7, 8, 9}
	report3 := []int{9, 7, 6, 2, 1}
	report4 := []int{1, 3, 2, 4, 5}
	report5 := []int{8, 6, 4, 4, 1}
	report6 := []int{1, 3, 6, 7, 9}


	fmt.Println("Hello World!")
	fmt.Println(report1, report2, report3, report4, report5, report6)
	fmt.Println(safeOrUnsafe(report1))
	fmt.Println(safeOrUnsafe(report2))
	fmt.Println(safeOrUnsafe(report3))
	fmt.Println(safeOrUnsafe(report4))
	fmt.Println(safeOrUnsafe(report5))
	fmt.Println(safeOrUnsafe(report6))
}
