package main

import (
	"fmt"
	"strconv"
)

// var pelajaran = "PHP"

func main() {
	// sekolah := "sd negeri"

	// fmt.Println("halo semuanya saya sedang bahasa " + sekolah)

	// umur := 74

	// fmt.Println("umur saya", umur)

	// var total int
	// a := 10
	// b := 70

	// total = a + b

	// fmt.Println(total)

	// gaji := 10000
	// gaji2 := strconv.Itoa(gaji)

	// fmt.Println("gaji saat ini adalah " + gaji2)

	// fmt.Println(getBiography(20, "eka pratama", "programmer"))

	daftarAngka := [...]string{"10", "20", "30", "40"}

	fmt.Println(daftarAngka)
}

func getBiography(age int, name string, status string) string {

	currentAge := strconv.Itoa(age)

	return name + " adalah seoran " + status + " berumur " + currentAge
}
