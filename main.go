package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/redgoose/pizza-day/excel"
	"github.com/redgoose/pizza-day/order"
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

	validClassCodes := map[string]bool{}
	for _, class := range conf.Classes {
		validClassCodes[class.Code] = true
	}

	processedRows, err := excel.ProcessFile(conf.File.Name, conf.File.SheetName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Orders to process:", len(processedRows))

	var ordersByClass = make(map[string][]order.Order)

	for _, row := range processedRows {
		// fmt.Println(row)

		// verify class codes
		classCode := ""
		if validClassCodes[row[3]] {
			classCode = row[3]
		} else {
			panic(fmt.Errorf("unknown class code encountered: %s", row[3]))
		}

		order := order.ParseOrder(row[11])
		order.Name = row[1]
		ordersByClass[classCode] = append(ordersByClass[classCode], order)
	}

	o, _ := json.MarshalIndent(ordersByClass, "", "\t")
	fmt.Println(string(o))

	SLICES_PER_PIZZA := conf.Pizza.SlicesPerPizza
	EXTRA_CHEESE_SLICES := conf.Pizza.ExtraCheeseSlices

	orderTotalsByClass, orderTotals := order.GetOrderTotals(ordersByClass, SLICES_PER_PIZZA, EXTRA_CHEESE_SLICES)

	otc, _ := json.MarshalIndent(orderTotalsByClass, "", "\t")
	fmt.Println(string(otc))

	ot, _ := json.MarshalIndent(orderTotals, "", "\t")
	fmt.Println(string(ot))
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
