package main

import "fmt"

// Element Function
/*
parameter/input
proses
return value
*/

func main() {
	fmt.Println(sayHello("Nama Saya Nichola Saputra"))
	fmt.Println(sayHello("Nama Saya Joko Morro"))
	fmt.Println(sayHello("Nama Saya Budi Raharjo"))

}

func sayHello(hi string) string {
	return hi
}
