package main

import "database/sql"

type Waitress struct {
	db *sql.DB
}

func NewWaitress(db *sql.DB) *Waitress {
	return &Waitress{db: db}
}

func (w Waitress) ListMenu() (Meals, error) {
	return w.getMeals()
}

func (w Waitress) getMeals() (Meals, error) {
	stmt, err := w.db.Prepare(`
	SELECT m.id, m.name, m.description, m.drink_id, d.name, d.description
	FROM meals AS m
		INNER JOIN drinks AS d ON d.id = m.drink_id`,
	)
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
		if err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.DrinkID, &m.Drink.Name, &m.Drink.Description); err != nil {
			return nil, err
		}

		meals = append(meals, m)
	}

	return meals, nil
}
