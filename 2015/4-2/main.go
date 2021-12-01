package main

import (
	"adventOfCode/lib"
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	num := 0
Iter:
	for ; ; num += 1 {
		str := fmt.Sprintf("%s%d", string(data), num)
		hash := md5.Sum([]byte(str))
		for x := range hash[0:3] {
			if hash[x] != 0 {
				continue Iter
			}
		}
		break
	}
	fmt.Printf("%d\n", num)
	os.Exit(0)
}
