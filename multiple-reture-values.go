package main

import "fmt"

func vals() (int, string) { // don't forget comma
	return 1, "freestyle"
}

func main() {

	num, str := vals()
	fmt.Println(num)
	fmt.Println(str)

	_, res := vals()
	fmt.Println(res)

}
