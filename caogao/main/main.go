package main

import "fmt"

func main() {
	s := "dfghj"
	for _, v := range s {
		fmt.Println(string(v))
	}
}
