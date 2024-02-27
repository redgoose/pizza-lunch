package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/redgoose/pizza-day/excel"
	"gopkg.in/yaml.v3"
)

type Order struct {
	Raw                    string
	Name                   string
	Drinks                 int
	CheeseSlices           int
	PepperoniSlices        int
	DairyFreeCheeseSlices  int
	GlutenFreeCheeseSlices int
}

type OrderTotal struct {
	Drinks                 int
	CheeseSlices           int
	PepperoniSlices        int
	DairyFreeCheeseSlices  int
	GlutenFreeCheeseSlices int
	CheesePizzas           int
	PepperoniPizzas        int
	DairyFreeCheesePizzas  int
	GlutenFreeCheesePizzas int
}

func main() {
	execute()
}

func execute() {
	conf, err := readConfig("pizza-day.yml")
	if err != nil {
		panic(err)
	}

	validClassCodes := map[string]bool{}
	for _, class := range conf.Classes {
		validClassCodes[class.Code] = true
	}

	processedRows, err := excel.ProcessFile(conf.File.Name, conf.File.SheetName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Orders to process:", len(processedRows))

	var ordersByClass = make(map[string][]Order)

	for _, row := range processedRows {
		// fmt.Println(row)

		// verify class codes
		classCode := ""
		if validClassCodes[row[3]] {
			classCode = row[3]
		} else {
			panic(fmt.Errorf("unknown class code encountered: %s", row[3]))
		}

		order := parseOrder(row[11])
		order.Name = row[1]
		ordersByClass[classCode] = append(ordersByClass[classCode], order)
	}

	o, _ := json.MarshalIndent(ordersByClass, "", "\t")
	fmt.Println(string(o))

	orderTotalsByClass := make(map[string]*OrderTotal)

	SLICES_PER_PIZZA := conf.Pizza.SlicesPerPizza
	EXTRA_CHEESE_SLICES := conf.Pizza.ExtraCheeseSlices

	orderTotals := OrderTotal{}

	for classCode, orders := range ordersByClass {
		orderTotalsByClass[classCode] = &OrderTotal{}
		for _, order := range orders {
			orderTotalsByClass[classCode].CheeseSlices += order.CheeseSlices
			orderTotalsByClass[classCode].PepperoniSlices += order.PepperoniSlices
			orderTotalsByClass[classCode].DairyFreeCheeseSlices += order.DairyFreeCheeseSlices
			orderTotalsByClass[classCode].GlutenFreeCheeseSlices += order.GlutenFreeCheeseSlices
			orderTotalsByClass[classCode].Drinks += order.Drinks

			orderTotals.CheeseSlices += order.CheeseSlices
			orderTotals.PepperoniSlices += order.PepperoniSlices
			orderTotals.DairyFreeCheeseSlices += order.DairyFreeCheeseSlices
			orderTotals.GlutenFreeCheeseSlices += order.GlutenFreeCheeseSlices
			orderTotals.Drinks += order.Drinks
		}

		orderTotalsByClass[classCode].CheeseSlices += EXTRA_CHEESE_SLICES

		if orderTotalsByClass[classCode].CheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].CheesePizzas = int(math.Floor(float64(orderTotalsByClass[classCode].CheeseSlices) / float64(SLICES_PER_PIZZA)))
			orderTotalsByClass[classCode].CheeseSlices = orderTotalsByClass[classCode].CheeseSlices - (orderTotalsByClass[classCode].CheesePizzas * SLICES_PER_PIZZA)
		}

		if orderTotalsByClass[classCode].PepperoniSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].PepperoniPizzas = int(math.Floor(float64(orderTotalsByClass[classCode].PepperoniSlices) / float64(SLICES_PER_PIZZA)))
			orderTotalsByClass[classCode].PepperoniSlices = orderTotalsByClass[classCode].PepperoniSlices - (orderTotalsByClass[classCode].PepperoniSlices * SLICES_PER_PIZZA)
		}

		if orderTotalsByClass[classCode].DairyFreeCheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].DairyFreeCheesePizzas = int(math.Floor(float64(orderTotalsByClass[classCode].DairyFreeCheeseSlices) / float64(SLICES_PER_PIZZA)))
			orderTotalsByClass[classCode].DairyFreeCheeseSlices = orderTotalsByClass[classCode].DairyFreeCheeseSlices - (orderTotalsByClass[classCode].DairyFreeCheeseSlices * SLICES_PER_PIZZA)
		}

		if orderTotalsByClass[classCode].GlutenFreeCheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].GlutenFreeCheesePizzas = int(math.Floor(float64(orderTotalsByClass[classCode].GlutenFreeCheeseSlices) / float64(SLICES_PER_PIZZA)))
			orderTotalsByClass[classCode].GlutenFreeCheeseSlices = orderTotalsByClass[classCode].GlutenFreeCheeseSlices - (orderTotalsByClass[classCode].GlutenFreeCheeseSlices * SLICES_PER_PIZZA)
		}
	}

	otc, _ := json.MarshalIndent(orderTotalsByClass, "", "\t")
	fmt.Println(string(otc))

	ot, _ := json.MarshalIndent(orderTotals, "", "\t")
	fmt.Println(string(ot))

}

func parseOrder(orderStr string) Order {

	orders := strings.Split(orderStr, ",")
	for i, order := range orders {
		orders[i] = strings.TrimSpace(order)
	}

	parsedOrder := Order{
		Raw: orderStr,
	}

	for _, order := range orders {

		numSlices, err := strconv.Atoi(order[0:1])
		if err != nil {
			panic(err)
		}

		if strings.Contains(order, "Dairy Free Cheese Pizza") {
			parsedOrder.DairyFreeCheeseSlices += numSlices
		} else if strings.Contains(order, "Cheese Gluten Free Pizza") {
			parsedOrder.GlutenFreeCheeseSlices += numSlices
		} else if strings.Contains(order, "Cheese Pizza") {
			parsedOrder.CheeseSlices += numSlices
		} else if strings.Contains(order, "Pepperoni Pizza") {
			parsedOrder.PepperoniSlices += numSlices
		} else {
			panic(errors.New("order could not be parsed"))
		}

		parsedOrder.Drinks++
	}

	return parsedOrder
}

type config struct {
	File    file    `yaml:"file"`
	Pizza   pizza   `yaml:"pizza"`
	Classes []class `yaml:"classes"`
}

type file struct {
	Name      string `yaml:"name"`
	SheetName string `yaml:"sheetName"`
}

type pizza struct {
	SlicesPerPizza    int `yaml:"slicesPerPizza"`
	ExtraCheeseSlices int `yaml:"extraCheeseSlices"`
}

type class struct {
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
