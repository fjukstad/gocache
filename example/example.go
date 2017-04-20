package main

import (
	"fmt"

	"github.com/fjukstad/gocache"
)

func main() {

	url := "http://blog.golang.org/error-handling-and-go"

	_, err := gocache.Get(url)

	if err != nil {
		fmt.Println("Get failed: ", err)
		return
	}

}
