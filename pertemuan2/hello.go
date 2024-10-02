package main

import "fmt"

func main() {
	var greetings = "Selamat Datang di Dunia DAP"
	var a, b int

	fmt.Println(greetings)
	fmt.Print("Masukkan angka pertama ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua ")
	fmt.Scanln(&b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
