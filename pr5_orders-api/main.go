package main

import (
	// "context"
	"app/pr5_orders-api/application"
	"context"
	"fmt"
	"github.com/PetyaKatsarova/more_go/pr5_orders-api/application"
)

func main() {
	fmt.Println("hello world :)")
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {	fmt.Println("failed to start app:", err) }
}

// go mod tidy       
// curl -X POST localhost:3000/hello -v         