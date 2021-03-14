package main

import (
	"time"

	"github.com/0B1t322/Note-API/app"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)

func main() {
	m := db.NewManager("root", "root", "note.db:3306", 20 *time.Second)
	db, err := m.OpenDataBase("notes_service")
	if err != nil {
		panic(err)
	}

	a := app.New(db, "8080")

	if err := a.Start(); err != nil {
		panic(err)
	}
}