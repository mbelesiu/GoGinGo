package server

import (
	"GoGinGo/internal/database"
	"GoGinGo/internal/store"
)

func Start() {

	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()

	router.Run(":8080")
}
