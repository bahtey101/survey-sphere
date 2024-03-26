package main

import (
	// "fmt"
	"src/database"
	"src/tests"
)

func main() {
	database.Init()

	tests.InitTestValues()
}
