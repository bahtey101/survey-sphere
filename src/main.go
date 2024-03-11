package main

import (
	"fmt"
	"src/database"
)

func main() {
	db := database.Init()
	fmt.Print(db)
}
