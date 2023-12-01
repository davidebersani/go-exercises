package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func concat2(s1, s2 string) string {
	return s1 + s2
}

// Le funzioni sono tipi di primo ordine, possono essere passati come parametri.

func main() {
	fmt.Println(concat("Lane,", " happy birthday!"))
	fmt.Println(concat("Elon,", " hope you're well!"))
	fmt.Println(concat("Go,", " is fantastic!"))
}

// Variables by value
