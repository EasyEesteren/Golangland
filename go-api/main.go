package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type basicResponse struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}


var home = basicResponse{
	ID:   "1",
	Body: "Homepage",
}


var kenobi = basicResponse{
	ID:   "1",
	Body: "Hello There",
}

var warOfTheStars = []basicResponse{
	{
		ID: "1",
		Body: "Phantom Menance",
	},
	{
		ID: "2",
		Body: "Attack of the Clones",
	},
	{
		ID: "3",
		Body: "Revenge of the Sith",
	},
	{
		ID: "4",
		Body: "A New Hope",
	},
	{
		ID: "5",
		Body: "The Empire Strikes Back",
	},
	{
		ID: "6",
		Body: "Return of the Jedi",
	},
}

func getHome(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, home)
}

func getHelloThere(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, kenobi)
}

func getWarOfTheStars(context *gin.Context) {
	id := context.Param("id")
	if (id == "all"){
		context.IndentedJSON(http.StatusOK, warOfTheStars)
		return
	}
	for _, movie := range warOfTheStars {
        if movie.ID == id {
            context.IndentedJSON(http.StatusOK, movie)
            return
        }
    }
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "This is not the movie you are looking for"})
}

func main() {
	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/test", getHelloThere)
	router.GET("/starwars/:id", getWarOfTheStars)
	router.Run("localhost:9090")
}
