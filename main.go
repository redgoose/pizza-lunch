package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/redgoose/pizza-day/excel"
	"github.com/redgoose/pizza-day/order"
	"github.com/redgoose/pizza-day/pdf"
	"gopkg.in/yaml.v3"
)

func main() {
	execute()
}

func execute() {
	conf, err := readConfig("pizza-day.yml")
	if err != nil {
		panic(err)
	}

	var roomInfo = make(map[string]room)
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

	// ot, _ := json.MarshalIndent(orderTotals, "", "\t")
	// fmt.Println(string(ot))

	pdf.BuildPDF(roomNumbers, orderTotalsByRoom, orderTotals)
	fmt.Println("PDF created :)")
}

type config struct {
	File  file   `yaml:"file"`
	Pizza pizza  `yaml:"pizza"`
	Rooms []room `yaml:"rooms"`
}

type file struct {
	Name string `yaml:"name"`
}

type pizza struct {
	SlicesPerPizza    int `yaml:"slicesPerPizza"`
	ExtraCheeseSlices int `yaml:"extraCheeseSlices"`
}

type room struct {
	Teacher string `yaml:"teacher"`
	Room    string `yaml:"room"`
	Class   string `yaml:"class"`
	Code    string `yaml:"code"`
}

func readConfig(configPath string) (config, error) {
	f, err := os.Open(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return config{}, fmt.Errorf("no such file %s", configPath)
	}
	if err != nil {
		return config{}, fmt.Errorf("failed to open config file %s: %w", configPath, err)
	}
	defer f.Close()

	var conf config
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return conf, fmt.Errorf("failed to read config file: %w", err)
	}

	return conf, nil
}
