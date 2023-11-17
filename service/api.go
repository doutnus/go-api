package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// An array of type Album (struct) to store the data in memory found in the Album (struct)
// Exemplary and initial data to start API testing.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more.
// Indented JSON serializes our Struct data structure into a JSON format.
// We pass an http status that we want to return to the client and then we pass our array containing our structure
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// We created a variable that is linked to the Album struct, that is, we can use all the fields in the struct
// We check if an error occurs with BindJSON and the response is different from nil, it just returns.
// We add new data (album) to the album and then serialize this new album again.
func SetAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
We search for the user-specified id with a repetition structure
The repetition structure together with another variable receives a "range" from the variable
albums that allows you to browse all the data inside it.const
If the variable that received a range from our structure that contains the data finds an ID equal to
when specified by the user, it will serialize the request and return.
*/
func GetSpecificyAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, v := range albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// We receive the ID as a parameter
// We convert the ID that comes as 'string' to 'int'
// We check if there is any error with the convertion and continue
// We pass a 'for' along with a variable that receives a range to cycle through the album
// If the ID passed by the user exists, the specified album will be removed
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	for _, v := range albums {
		if v.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusAccepted, v)
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to delete album"})
}

/*
Note on decrement:

Every time we perform a search for the ID, our code
It does not lead directly to the album ID, but rather its index in Slice.

Considering this, let's say you want to change the album ID 2,
But in reality you end up changing the 3, because the entire index
Slice starts at 0, so the correct option would be to decrement it, staying on ID 1
This way, getting the position in the index correct and the correct album.
------------------------------------------------------------------------------------
1. We start by getting the id parameter
2. We convert ID to Integer type
3. We create a new variable containing the data from the Album struct
4. We collect the JSON from the request body and attach it to the variable
5. We use the decrement operator to actually set the album we want
*/
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	var updateAlbum Album

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}

	i--
	for _, v := range albums {
		if v.ID == id {
			albums[i].Title = updateAlbum.Title
			albums[i].Artist = updateAlbum.Artist
			albums[i].Price = updateAlbum.Price
			c.IndentedJSON(http.StatusAccepted, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to update album"})
}
