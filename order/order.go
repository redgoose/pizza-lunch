package order

import (
	"errors"
	"math"
	"strconv"
	"strings"
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

func GetOrderTotals(ordersByClass map[string][]Order, SLICES_PER_PIZZA int, EXTRA_CHEESE_SLICES int) (map[string]*OrderTotal, OrderTotal) {
	orderTotalsByClass := make(map[string]*OrderTotal)
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
			orderTotalsByClass[classCode].CheesePizzas, orderTotalsByClass[classCode].CheeseSlices = SlicesToWholePizzas(orderTotalsByClass[classCode].CheeseSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByClass[classCode].PepperoniSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].PepperoniPizzas, orderTotalsByClass[classCode].PepperoniSlices = SlicesToWholePizzas(orderTotalsByClass[classCode].PepperoniSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByClass[classCode].DairyFreeCheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].DairyFreeCheesePizzas, orderTotalsByClass[classCode].DairyFreeCheeseSlices = SlicesToWholePizzas(orderTotalsByClass[classCode].DairyFreeCheeseSlices, SLICES_PER_PIZZA)
		}

		if orderTotalsByClass[classCode].GlutenFreeCheeseSlices >= SLICES_PER_PIZZA {
			orderTotalsByClass[classCode].GlutenFreeCheesePizzas, orderTotalsByClass[classCode].GlutenFreeCheeseSlices = SlicesToWholePizzas(orderTotalsByClass[classCode].GlutenFreeCheeseSlices, SLICES_PER_PIZZA)
		}
	}

	orderTotals.CheesePizzas = PizzasToOrder(orderTotals.CheeseSlices, SLICES_PER_PIZZA)
	orderTotals.PepperoniPizzas = PizzasToOrder(orderTotals.PepperoniSlices, SLICES_PER_PIZZA)
	orderTotals.DairyFreeCheesePizzas = PizzasToOrder(orderTotals.DairyFreeCheeseSlices, SLICES_PER_PIZZA)
	orderTotals.GlutenFreeCheesePizzas = PizzasToOrder(orderTotals.GlutenFreeCheeseSlices, SLICES_PER_PIZZA)

	return orderTotalsByClass, orderTotals
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
