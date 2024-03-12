package order

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Order struct {
	Raw                       string
	Name                      string
	Drinks                    int
	CheeseSlices              int
	PepperoniSlices           int
	DairyFreeCheeseSlices     int
	GlutenFreeCheeseSlices    int
	GlutenFreePepperoniSlices int
}

type OrderTotal struct {
	Drinks                    int
	CheeseSlices              int
	ExtraCheeseSlices         int
	PepperoniSlices           int
	DairyFreeCheeseSlices     int
	GlutenFreeCheeseSlices    int
	GlutenFreePepperoniSlices int
	CheesePizzas              int
	PepperoniPizzas           int
	DairyFreeCheesePizzas     int
	GlutenFreeCheesePizzas    int
	GlutenFreePepperoniPizzas int
}

func GetOrderTotalsByRoom(ordersByRoom map[string][]Order, SLICES_PER_PIZZA int, EXTRA_CHEESE_SLICES int) map[string]*OrderTotal {
	orderTotalsByRoom := make(map[string]*OrderTotal)

	for roomNumber, orders := range ordersByRoom {
		orderTotalsByRoom[roomNumber] = &OrderTotal{}
		for _, order := range orders {
			orderTotalsByRoom[roomNumber].CheeseSlices += order.CheeseSlices
			orderTotalsByRoom[roomNumber].PepperoniSlices += order.PepperoniSlices
			orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices += order.DairyFreeCheeseSlices
			orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices += order.GlutenFreeCheeseSlices
			orderTotalsByRoom[roomNumber].GlutenFreePepperoniSlices += order.GlutenFreePepperoniSlices
			orderTotalsByRoom[roomNumber].Drinks += order.Drinks
		}

		orderTotalsByRoom[roomNumber].CheeseSlices += EXTRA_CHEESE_SLICES

		if orderTotalsByRoom[roomNumber].CheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByRoom[roomNumber].CheesePizzas, orderTotalsByRoom[roomNumber].CheeseSlices = SlicesToWholePizzas(orderTotalsByRoom[roomNumber].CheeseSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByRoom[roomNumber].PepperoniSlices >= SLICES_PER_PIZZA {
			orderTotalsByRoom[roomNumber].PepperoniPizzas, orderTotalsByRoom[roomNumber].PepperoniSlices = SlicesToWholePizzas(orderTotalsByRoom[roomNumber].PepperoniSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByRoom[roomNumber].DairyFreeCheesePizzas, orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices = SlicesToWholePizzas(orderTotalsByRoom[roomNumber].DairyFreeCheeseSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByRoom[roomNumber].GlutenFreeCheesePizzas, orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices = SlicesToWholePizzas(orderTotalsByRoom[roomNumber].GlutenFreeCheeseSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByRoom[roomNumber].GlutenFreePepperoniSlices >= SLICES_PER_PIZZA {
			orderTotalsByRoom[roomNumber].GlutenFreePepperoniPizzas, orderTotalsByRoom[roomNumber].GlutenFreePepperoniSlices = SlicesToWholePizzas(orderTotalsByRoom[roomNumber].GlutenFreePepperoniSlices, SLICES_PER_PIZZA)
		}
	}

	return orderTotalsByRoom
}

func GetOrderTotals(ordersByRoom map[string][]Order, SLICES_PER_PIZZA int, EXTRA_CHEESE_SLICES int) OrderTotal {
	orderTotals := OrderTotal{}

	for _, orders := range ordersByRoom {
		for _, order := range orders {
			orderTotals.CheeseSlices += order.CheeseSlices
			orderTotals.PepperoniSlices += order.PepperoniSlices
			orderTotals.DairyFreeCheeseSlices += order.DairyFreeCheeseSlices
			orderTotals.GlutenFreeCheeseSlices += order.GlutenFreeCheeseSlices
			orderTotals.GlutenFreePepperoniSlices += order.GlutenFreePepperoniSlices
			orderTotals.Drinks += order.Drinks
		}
		orderTotals.CheeseSlices += EXTRA_CHEESE_SLICES
		orderTotals.ExtraCheeseSlices += EXTRA_CHEESE_SLICES // keep count for pdf
	}

	orderTotals.CheesePizzas = PizzasToOrder(orderTotals.CheeseSlices, SLICES_PER_PIZZA)
	orderTotals.PepperoniPizzas = PizzasToOrder(orderTotals.PepperoniSlices, SLICES_PER_PIZZA)
	orderTotals.DairyFreeCheesePizzas = PizzasToOrder(orderTotals.DairyFreeCheeseSlices, SLICES_PER_PIZZA)
	orderTotals.GlutenFreeCheesePizzas = PizzasToOrder(orderTotals.GlutenFreeCheeseSlices, SLICES_PER_PIZZA)
	orderTotals.GlutenFreePepperoniPizzas = PizzasToOrder(orderTotals.GlutenFreePepperoniSlices, SLICES_PER_PIZZA)

	return orderTotals
}

func SlicesToWholePizzas(slices int, slicesPerPizza int) (pizzas int, remainingSlices int) {
	pizzas = int(math.Floor(float64(slices) / float64(slicesPerPizza)))
	remainingSlices = slices - (pizzas * slicesPerPizza)
	return pizzas, remainingSlices
}

func PizzasToOrder(slices int, slicesPerPizza int) int {
	if slices >= 1 {
		pizzasToOrder := math.Max(float64(slices), float64(slicesPerPizza)) / float64(slicesPerPizza)
		pizzasToOrder = math.Ceil(pizzasToOrder)
		return int(pizzasToOrder)
	}
	return 0
}

func ParseOrder(orderStr string) Order {

	orders := strings.Split(orderStr, ",")
	for i, order := range orders {
		orders[i] = strings.TrimSpace(order)
	}

	parsedOrder := Order{
		Raw: orderStr,
	}

	for _, order := range orders {

		numSlices := 0
		if strings.HasPrefix(order, "Additional") {
			numSlices = 1
		} else {
			var err error
			numSlices, err = strconv.Atoi(order[0:1])
			if err != nil {
				panic(fmt.Errorf("could not determine number of slices: %s", order))
			}
		}

		if strings.Contains(order, "Dairy Free Cheese Pizza") {
			parsedOrder.DairyFreeCheeseSlices += numSlices
		} else if strings.Contains(order, "Cheese Gluten Free Pizza") {
			parsedOrder.GlutenFreeCheeseSlices += numSlices
		} else if strings.Contains(order, "Cheese Pizza") {
			parsedOrder.CheeseSlices += numSlices
		} else if strings.Contains(order, "Pepperoni Gluten Free Pizza") {
			parsedOrder.GlutenFreePepperoniSlices += numSlices
		} else if strings.Contains(order, "Pepperoni Pizza") {
			parsedOrder.PepperoniSlices += numSlices
		} else {
			panic(fmt.Errorf("order could not be parsed: %s", order))
		}

		if strings.Contains(order, "Drink") {
			parsedOrder.Drinks++
		}
	}

	return parsedOrder
}
