package main

import "fmt"

// Meal model for a record of meals table
type Meal struct {
	ID          uint
	Name        string
	Description string
	DrinkID     uint
	Drink       Drink
}

type Meals []Meal

func (m Meals) JoinDrinks(drink Drinks) {
	for i, meal := range m {
		m[i].Drink = drink.GetByID(meal.DrinkID)
	}
}

func (m Meals) GetUniqueDrinkIDs() []uint {
	var ids []uint
	drinks := make(map[uint]struct{}, 0)

	for _, v := range m {
		_, ok := drinks[v.DrinkID]
		if ok {
			continue
		}

		drinks[v.DrinkID] = struct{}{}
		ids = append(ids, v.DrinkID)
	}

	fmt.Println(ids)

	return ids
}

// Drink model for a record of drinks table
type Drink struct {
	ID          uint
	Name        string
	Description string
}

type Drinks []Drink

func (d Drinks) GetByID(ID uint) Drink {
	for _, drink := range d {
		if drink.ID == ID {
			return drink
		}
	}

	return Drink{}
}
