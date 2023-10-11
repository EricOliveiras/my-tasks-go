package main

import (
	application "github/ericoliveiras/basic-crud-go"
	"github/ericoliveiras/basic-crud-go/internal/config"
)

func main() {
	cgf := config.NewConfig()

	application.Start(cgf)
}
