package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getOxygenRating(nums []uint16, mask uint16) uint16 {
	if len(nums) == 1 {
		return nums[0]
	}
	var set, unset []uint16
	for _, num := range nums {
		if num&mask == mask {
			set = append(set, num)
		} else {
			unset = append(unset, num)
		}
	}
	if len(set) >= len(unset) {
		return getOxygenRating(set, mask>>1)
	} else {
		return getOxygenRating(unset, mask>>1)
	}
}

func getCo2Rating(nums []uint16, mask uint16) uint16 {
	if len(nums) == 1 {
		return nums[0]
	}
	var set, unset []uint16
	for _, num := range nums {
		if num&mask == mask {
			set = append(set, num)
		} else {
			unset = append(unset, num)
		}
	}
	if len(unset) <= len(set) {
		return getCo2Rating(unset, mask>>1)
	} else {
		return getCo2Rating(set, mask>>1)
	}
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var numbers []uint16
	for line, text := range strings.Split(string(data), "\n") {
		if text == "" {
			continue
		}
		num, err := strconv.ParseUint(text, 2, 16)
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error on line %d: %v", line, err))
		}
		numbers = append(numbers, uint16(num))
	}
	oxygen := getOxygenRating(numbers, 2048)
	co2 := getCo2Rating(numbers, 2048)
	fmt.Printf("%d\n", uint32(oxygen)*uint32(co2))
	os.Exit(0)
}
