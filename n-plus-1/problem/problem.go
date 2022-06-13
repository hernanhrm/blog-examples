package main

import (
	"database/sql"
)

type Waitress struct {
	db *sql.DB
}

func NewWaitress(db *sql.DB) *Waitress {
	return &Waitress{db: db}
}

func (b Waitress) ListMenu() (Meals, error) {
	// the `1` of our `N+1`
	meals, err := b.getMeals()
	if err != nil {
		return nil, err
	}

	// here is our waitress going back and forth from floor 1 to 10,
	// will be the `N` of our `N+1`
	for i, meal := range meals {
		drink, err := b.getDrinkByID(meal.DrinkID)
		if err != nil {
			return nil, err
		}

		meals[i].Drink = drink
	}

	return meals, nil
}

func (b Waitress) getMeals() (Meals, error) {
	stmt, err := b.db.Prepare(`SELECT id, name, description, drink_id FROM meals LIMIT 20`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals Meals
	for rows.Next() {
		m := Meal{}
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.DrinkID); err != nil {
			return nil, err
		}

		meals = append(meals, m)
	}

	return meals, nil
}

func (b Waitress) getDrinkByID(ID uint) (Drink, error) {
	stmt, err := b.db.Prepare(`SELECT id, name, description FROM drinks WHERE id = $1`)
	if err != nil {
		return Drink{}, err
	}
	defer stmt.Close()

	m := Drink{}
	if err := stmt.QueryRow(ID).Scan(&m.ID, &m.Name, &m.Description); err != nil {
		return Drink{}, err
	}

	return m, nil
}
