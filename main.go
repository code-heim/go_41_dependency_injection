package main

import "fmt"

type Oven interface {
	Heat() string
}

type Ingredient interface {
	Mix() string
}

type GasOven struct{}

func (gasOven GasOven) Heat() string {
	return "Heating with Gas Oven!"
}

type ElectricOven struct{}

func (electricOven ElectricOven) Heat() string {
	return "Heating with electric oven!"
}

type Flour struct{}

func (f Flour) Mix() string {
	return "Mixing flour!"
}

type Sugar struct{}

func (s Sugar) Mix() string {
	return "Mixing sugar!"
}

type Butter struct{}

func (b Butter) Mix() string {
	return "Adding butter!"
}

// Bakery is a struct that depends on the Oven and Ingredients.
type Bakery struct {
	oven        Oven
	ingredients []Ingredient
}

// Bake is a method that uses the oven and ingredient to bake a pastry.
func (b *Bakery) Bake() {
	fmt.Println(b.oven.Heat())
	for _, ingredient := range b.ingredients {
		fmt.Println(ingredient.Mix())
	}

	fmt.Println("Baking an awesome pastry!")
}

func main() {
	gasOven := GasOven{}
	electricOven := ElectricOven{}
	flour := Flour{}
	sugar := Sugar{}
	butter := Butter{}

	bakery := Bakery{oven: gasOven, ingredients: []Ingredient{flour, butter, sugar}}
	bakery.Bake()

	bakery = Bakery{oven: electricOven, ingredients: []Ingredient{sugar, butter}}
	bakery.Bake()
}
