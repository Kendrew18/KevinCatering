package main

import (
	"KevinCatering/db"
	"KevinCatering/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
