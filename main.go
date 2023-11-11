package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/google/uuid"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.MaxMultipartMemory = 1 << 20  // 1 MiB
    router.POST("/submit_contract", submitContract)

    router.Run("localhost:8081")
}


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

// getAlbums responds with the list of all albums as JSON.
func submitContract(c *gin.Context) {
    // single file
    file, _ := c.FormFile("file")
    log.Println(file.Filename)
    // Upload the file with random name
    targetPath := "submissions/" + uuid.New().String() + ".sol"
    c.SaveUploadedFile(file, targetPath)
    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

