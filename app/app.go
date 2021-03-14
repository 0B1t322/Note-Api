package app

import (
	"github.com/0B1t322/Note-API/models/note"
	"fmt"
	"net/http"

	"github.com/0B1t322/Note-API/handlers"
	"github.com/0B1t322/Note-API/routers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type app struct {
	r *mux.Router
	port string
	db *gorm.DB
}

type App interface {
	Start() error
}

func New(db *gorm.DB, port string) App {
	return &app{
		r: routers.New(handlers.New(db)),
		port: port,
		db: db,
	}
}

func (a *app) Start() error {
	if err := a.init(); err != nil {
		return err
	}
	
	return http.ListenAndServe(fmt.Sprintf(":%s", a.port), a.r)
}

func (a *app) init() error {
	return note.AutoMigrate(a.db)
}