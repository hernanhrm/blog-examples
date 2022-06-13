package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type Waitress struct {
	db *sql.DB
}

func NewWaitress(db *sql.DB) *Waitress {
	return &Waitress{db: db}
}

func (w Waitress) ListMenu() (Meals, error) {
	meals, err := w.getMeals()
	if err != nil {
		return nil, err
	}

	drinks, err := w.getDrinksByIDsIn(meals.GetUniqueDrinkIDs())
	if err != nil {
		return nil, err
	}

	meals.JoinDrinks(drinks)

	return meals, nil
}

func (w Waitress) getMeals() (Meals, error) {
	stmt, err := w.db.Prepare(`SELECT id, name, description, drink_id FROM meals LIMIT 20`)
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

func (w Waitress) getDrinksByIDsIn(ids []uint) (Drinks, error) {
	query := fmt.Sprintf(`SELECT id, name, description FROM drinks WHERE id IN (%s)`, buildIN(ids))
	fmt.Println(query)
	stmt, err := w.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drinks Drinks
	for rows.Next() {
		m := Drink{}
		if err := rows.Scan(&m.ID, &m.Name, &m.Description); err != nil {
			return nil, err
		}

		drinks = append(drinks, m)
	}

	return drinks, nil
}

func buildIN(ids []uint) string {
	args := ""
	for _, id := range ids {
		args += fmt.Sprintf("%d,", id)
	}

	return strings.TrimSuffix(args, ",")
}
