package main
// https://www.youtube.com/watch?v=NCaKbDVogDI&t=220s

import (
	// "app/pr5_orders-api/application"
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/PetyaKatsarova/more_go/pr5_orders-api/application"
	// "github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("hello world :)")
	app := application.New()

	ctx, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelFunc() // defer means complete after curr func is done

	err := app.Start(ctx)
	if err != nil {	fmt.Println("failed to start app:", err) }
}

// go mod tidy     after adding the github link in this file          
// to run the app on windows 11:
/*
0.0 go get github.com/redis/go-redis/v9 // ram db
0.1 install docker, start it from the app docker
0.3 in terminal: docker ps // output: container id image command created status ports names
0.4 sc stop redis // sc is service control ? didnt get confirmation: successfully stopped redis
1. windows: command prompt: docker run -p 6379:6379 redis:latest
2. in terminal: go run main.go
3.  curl -X POST localhost:3000 -v 
*/

/*
attempt 2: 
1. download docker desktop app, bash for windows 11
2. in bash: docker pull redis
3.Microsoft Windows [Version 10.0.22621.2428] in command prompt
(c) Microsoft Corporation. All rights reserved.

C:\Users\petya.katsarova>docker start 908acd543874
908acd543874

C:\Users\petya.katsarova>docker exec -it 908acd543874 redis-cli
127.0.0.1:6379>
-- if u have only one usage of each socket address(protocol/network addr/port) is
normally permitted: in command prompt run:
netstat -ano | findstr :3000 to find and after kill the process
taskkill /F /PID <PID>
--  go get github.com/google/uuid
*/