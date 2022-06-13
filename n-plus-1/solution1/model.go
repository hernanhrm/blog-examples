package main

// Meal model for a record of meals table
type Meal struct {
	ID          uint
	Name        string
	Description string
	DrinkID     uint
	Drink       Drink
}

type Meals []Meal

// Drink model for a record of drinks table
type Drink struct {
	ID          uint
	Name        string
	Description string
}

type Drinks []Drink
