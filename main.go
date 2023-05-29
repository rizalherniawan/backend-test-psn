package main

import (
	"github.com/rizalherniawan/backend-test-psn/config"
	"github.com/rizalherniawan/backend-test-psn/routes"
)

func main() {
	db := config.InitDB()
	routes.Handler(db)
}
