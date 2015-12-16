package main

import (
	"fmt"

	"github.com/fjukstad/gocache"
)

func main() {

	gocache.SetInvalidationTime("1m")

	url := "http://blog.golang.org/error-handling-and-go"

	_, err := gocache.Get(url)

	if err != nil {
		fmt.Println("Get failed: ", err)
		return
	}

}
