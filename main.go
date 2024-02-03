package main

import (
	"fmt"

	"github.com/redgoose/pizza-day/excel"
)

func main() {
	processedRows, err := excel.ProcessFile("pizza.xlsx", "1. Pizza Lunch Original")
	if err != nil {
		panic(err)
	}

	fmt.Println(processedRows)

}
