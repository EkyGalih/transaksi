package main

import (
	"github.com/EkyGalih/transaksi/db"
	"github.com/EkyGalih/transaksi/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
