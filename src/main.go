package main

import (
	"src/database"
	"src/tests"
)

func main() {
	database.Init()

	database.CreateTables()

	tests.TEST_1()

	//database.DeleteUser(user1)
	//database.DeleteUser(user2)
}
