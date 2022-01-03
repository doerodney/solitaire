package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello from solitaire.")
	pRow := flag.Uint("row", 0, "row of initially empty cell")
	pColumn := flag.Uint("column", 2, "column of initally empty cell")
	flag.Parse()

	fmt.Printf("row %d\n", *pRow)
	fmt.Printf("column: %d\n", *pColumn)

	cell := NewCell(*pColumn, *pRow, true, true)
	fmt.Printf("%v\n", *cell)
}
