package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/redgoose/pizza-day/config"
	"github.com/redgoose/pizza-day/excel"
	"github.com/redgoose/pizza-day/order"
	"github.com/redgoose/pizza-day/pdf"
)

func main() {
	execute()
}

func execute() {
	conf, err := config.ReadConfig("pizza-day.yml")
	if err != nil {
		panic(err)
	}

	var roomInfo = make(map[string]config.Room)
	var roomNumbers = []string{}

	for _, room := range conf.Rooms {
		roomCodes := strings.Split(room.Code, "|")
		for _, roomCode := range roomCodes {
			roomInfo[roomCode] = room
		}

		roomNumbers = append(roomNumbers, room.Room)
	}

	sort.Strings(roomNumbers)

	processedRows, err := excel.ProcessFile(conf.File.Name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Processing %d orders...\n", len(processedRows))

	var ordersByRoom = make(map[string][]order.Order)

	for _, row := range processedRows {
		// fmt.Println(row)

		// verify room code exists in config
		roomCode := ""
		if _, ok := roomInfo[row[3]]; ok {
			roomCode = row[3]
		} else {
			panic(fmt.Errorf("unexpected room code: %s", row[3]))
		}

		order := order.ParseOrder(row[11])
		order.Name = row[1]

		roomNumber := roomInfo[roomCode].Room
		ordersByRoom[roomNumber] = append(ordersByRoom[roomNumber], order)
	}

	// o, _ := json.MarshalIndent(ordersByRoom, "", "\t")
	// fmt.Println(string(o))

	SLICES_PER_PIZZA := conf.Pizza.SlicesPerPizza
	EXTRA_CHEESE_SLICES := conf.Pizza.ExtraCheeseSlices

	orderTotalsByRoom := order.GetOrderTotalsByRoom(ordersByRoom, SLICES_PER_PIZZA, EXTRA_CHEESE_SLICES)
	orderTotals := order.GetOrderTotals(ordersByRoom, SLICES_PER_PIZZA, EXTRA_CHEESE_SLICES)

	// otr, _ := json.MarshalIndent(orderTotalsByRoom, "", "\t")
	// fmt.Println(string(otr))

	// for _, roomNumber := range roomNumbers {
	// 	fmt.Println(roomNumber, "\t", orderTotalsByRoom[roomNumber].CheesePizzas*8+orderTotalsByRoom[roomNumber].CheeseSlices-2)
	// }

	// ot, _ := json.MarshalIndent(orderTotals, "", "\t")
	// fmt.Println(string(ot))

	pdf.BuildPDF(roomNumbers, roomInfo, ordersByRoom, orderTotalsByRoom, orderTotals)
	fmt.Println("PDF created :)")
}
