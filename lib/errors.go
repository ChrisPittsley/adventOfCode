package lib

import (
	"fmt"
	"os"
)

func ErrorOut(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
