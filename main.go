package main

import (
  	"single-service/databases"
    "single-service/routes"
)

func main() {
	databases.ConnectDatabase()
    routes.Route()
}