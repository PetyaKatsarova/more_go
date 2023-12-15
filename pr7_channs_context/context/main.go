package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default() // for router

	r.GET("/hello", func (ctx *gin.Context)  {
		
		// the context stops the process once the timout is reached
	timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second)
	defer cancel()

											//ctx.Request.Context()
		req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://yahoo.com", nil)
		if err != nil { panic(err) }
		
		res, err := http.DefaultClient.Do(req)
		if err != nil { panic(err) }
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil { panic(err) }

		ctx.Data(200, "text/html", data)
	})

	r.Run()
}