package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SLICES_PER_PIZZA = 8
var EXTRA_CHEESE_SLICES = 2

func TestSlicesToWholePizzasWithExtraSlices(t *testing.T) {
	slices := 36 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices := SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 4, pizzas)
	assert.Equal(t, 6, remainingSlices)

	slices = 21 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 2, pizzas)
	assert.Equal(t, 7, remainingSlices)

	slices = 31 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 4, pizzas)
	assert.Equal(t, 1, remainingSlices)

	slices = 17 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 2, pizzas)
	assert.Equal(t, 3, remainingSlices)

	slices = 8 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 1, pizzas)
	assert.Equal(t, 2, remainingSlices)

	slices = 6 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 1, pizzas)
	assert.Equal(t, 0, remainingSlices)
}

func TestSlicesToWholePizzas(t *testing.T) {
	slices := 36
	pizzas, remainingSlices := SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 4, pizzas)
	assert.Equal(t, 4, remainingSlices)

	slices = 8
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 1, pizzas)
	assert.Equal(t, 0, remainingSlices)

	slices = 1
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 0, pizzas)
	assert.Equal(t, 1, remainingSlices)

	slices = 0
	pizzas, remainingSlices = SlicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 0, pizzas)
	assert.Equal(t, 0, remainingSlices)
}

func TestPizzasToOrder(t *testing.T) {

	slices := 594
	pizzas := PizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 75, pizzas)

	slices = 278
	pizzas = PizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 35, pizzas)

	slices = 12
	pizzas = PizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 2, pizzas)

	slices = 5
	pizzas = PizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 1, pizzas)

	slices = 0
	pizzas = PizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 0, pizzas)
}

func TestParseOrder(t *testing.T) {
	orderStr := "1 Slices Cheese Pizza and Drink"
	expectedOrder := Order{
		Raw:                    orderStr,
		CheeseSlices:           1,
		PepperoniSlices:        0,
		DairyFreeCheeseSlices:  0,
		GlutenFreeCheeseSlices: 0,
		Drinks:                 1,
	}

	order := ParseOrder(orderStr)
	assert.Equal(t, expectedOrder, order)

	orderStr = "3 Slices Pepperoni Pizza and Drink"
	expectedOrder = Order{
		Raw:                    orderStr,
		CheeseSlices:           0,
		PepperoniSlices:        3,
		DairyFreeCheeseSlices:  0,
		GlutenFreeCheeseSlices: 0,
		Drinks:                 1,
	}

	order = ParseOrder(orderStr)
	assert.Equal(t, expectedOrder, order)

	orderStr = "2 Slices Dairy Free Cheese Pizza and Drink"
	expectedOrder = Order{
		Raw:                    orderStr,
		CheeseSlices:           0,
		PepperoniSlices:        0,
		DairyFreeCheeseSlices:  2,
		GlutenFreeCheeseSlices: 0,
		Drinks:                 1,
	}

	order = ParseOrder(orderStr)
	assert.Equal(t, expectedOrder, order)

	orderStr = "3 Slices Cheese Gluten Free Pizza and Drink"
	expectedOrder = Order{
		Raw:                    orderStr,
		CheeseSlices:           0,
		PepperoniSlices:        0,
		DairyFreeCheeseSlices:  0,
		GlutenFreeCheeseSlices: 3,
		Drinks:                 1,
	}

	order = ParseOrder(orderStr)
	assert.Equal(t, expectedOrder, order)

	orderStr = "1 Slice Cheese Pizza and Drink, 1 Slice Pepperoni Pizza and Drink"
	expectedOrder = Order{
		Raw:                    orderStr,
		CheeseSlices:           1,
		PepperoniSlices:        1,
		DairyFreeCheeseSlices:  0,
		GlutenFreeCheeseSlices: 0,
		Drinks:                 2,
	}

	order = ParseOrder(orderStr)
	assert.Equal(t, expectedOrder, order)
}

func TestParseOrderPanics(t *testing.T) {
	// test non number prefix
	assertPanic(t, func() { ParseOrder("X Slices Cheese Pizza and Drink") })

	// test non pizza slice type
	assertPanic(t, func() { ParseOrder("2 foobar") })
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
