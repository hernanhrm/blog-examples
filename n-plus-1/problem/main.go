package main

import (
	"fmt"
	"log"

	n_plus_1 "github.com/hernanhrm/n1-problem"
)

func main() {
	db, err := n_plus_1.DBConnection()
	if err != nil {
		log.Fatalln(err)
	}
	waitress := NewWaitress(db)

	menu, err := waitress.ListMenu()
	if err != nil {
		log.Fatalln(err)
	}

	for _, meal := range menu {
		fmt.Println(meal)
	}
}
