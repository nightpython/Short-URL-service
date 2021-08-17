package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/shortener"
	"github.com/my/repo/store"
	"net/http"
)
//модель запроса
type UrlCreationRequest struct {
	LongUrl	string	`json:"long_url"`
	UserId string  `json:"user_id"`
}
//парсим JSON, генерируем сокращенный URL, сохраняем сопоставление короткого и длинного URL
func CreateShortUrl(c *gin.Context)  {
	var creationRequest UrlCreationRequest
	if err:=c.ShouldBindJSON(&creationRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	shortUrl :=shortener.UrlGenerator(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl,creationRequest.LongUrl,creationRequest.UserId)
	host:="http://localhost:9808/"
	c.JSON(200, gin.H{
		"message" : "Short url created, congratulations!",
		"short url": host + shortUrl,
	})
}


func RedirectLongUrl(c *gin.Context)  {
	shortUrl := c.Param("shortUrl")
	longUrl := store.GetInitialUrl(shortUrl)
	c.Redirect(302, longUrl)
}