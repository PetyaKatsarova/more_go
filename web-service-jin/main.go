package main

// go mod init example/web-service-gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// storing data in memory means that the set of albums will be
// lost each time you stop the server, then recreated when you start it.
/*
gin.Context is the most important part of Gin. It carries request details,
validates and serializes JSON, and more. (Despite the similar name, this is
     different from Go’s built-in context package.)
*/
func getAlbums(c *gin.Context) {
	//serialize the struct into JSON and add it to the response.
	// Context.JSON to send more compact JSON. indented form is much easier to work with when debugging and the size difference is usually small.
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum album // type album
    if err := c.BindJSON(&newAlbum); err != nil { return }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, c)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "album not found"})
}

func main() {
	fmt.Println("hello gin")
	router := gin.Default()

	router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
    router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080") //attach the r to an http.Server and start the it
}

// after go run . ,for GET req: do in another terminal: curl http://localhost:8080/albums

// for post req: 
/*

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'*/

    /*
    for get album by id in separate terminal after go run . 
    curl http://localhost:8080/albums/2
    */