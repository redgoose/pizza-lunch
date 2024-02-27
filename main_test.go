package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SLICES_PER_PIZZA = 8
var EXTRA_CHEESE_SLICES = 2

func TestSlicesWithExtraSlicesToWholePizzas(t *testing.T) {
	slices := 36 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices := slicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 4, pizzas)
	assert.Equal(t, 6, remainingSlices)

	slices = 21 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = slicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 2, pizzas)
	assert.Equal(t, 7, remainingSlices)

	slices = 31 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = slicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 4, pizzas)
	assert.Equal(t, 1, remainingSlices)

	slices = 17 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = slicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 2, pizzas)
	assert.Equal(t, 3, remainingSlices)

	slices = 8 + EXTRA_CHEESE_SLICES
	pizzas, remainingSlices = slicesToWholePizzas(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 1, pizzas)
	assert.Equal(t, 2, remainingSlices)
}

func TestPizzasToOrder(t *testing.T) {

	slices := 594
	pizzas := pizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 75, pizzas)

	slices = 278
	pizzas = pizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 35, pizzas)

	slices = 12
	pizzas = pizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 2, pizzas)

	slices = 5
	pizzas = pizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 1, pizzas)

	slices = 0
	pizzas = pizzasToOrder(slices, SLICES_PER_PIZZA)
	assert.Equal(t, 0, pizzas)
}
