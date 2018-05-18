package main

import (
  "database/sql"

  "github.com/gorilla/mux"
   _ "github.com/lib/pq"
)

type App struct {
  Router *mux.Router
  DB *sql.DB
}

// initialize => create db connection && routes
func (a *App) Initialize(user, password, dbname string) { }

// start application
func (a *App) Run(addr string) { }
