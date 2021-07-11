package controllers

import (
	"course-go/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Articles struct {
}

type createFormArticle struct {
	Title string                `form:"title" binding:"required"`
	Body  string                `form:"body" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

var articles []models.Article = []models.Article{
	{ID: 1, Title: "Title#1", Body: "Body#1"},
	{ID: 2, Title: "Title#2", Body: "Body#2"},
	{ID: 3, Title: "Title#3", Body: "Body#3"},
	{ID: 4, Title: "Title#4", Body: "Body#4"},
	{ID: 5, Title: "Title#5", Body: "Body#5"},
}

func (a *Articles) FindAll(ctx *gin.Context) {
	result := articles
	if limit := ctx.Query("limit"); limit != "" {
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
			return
		}
		result = result[:limit]
	}
	ctx.JSON(http.StatusOK, gin.H{"articles": result})
}

func (a *Articles) FindOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return
	}
	for _, article := range articles {
		if article.ID == uint(id) {
			ctx.JSON(http.StatusOK, gin.H{"article": article})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Article Not Found!"})
}

func (a *Articles) Create(ctx *gin.Context) {
	var form createFormArticle
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newArticle := models.Article{
		ID:    uint(len(articles) + 1),
		Title: form.Title,
		Body:  form.Body,
	}

	// Get image for Form request
	image, _ := ctx.FormFile("image")

	// Set path uploads file
	path := "uploads/articles/" + strconv.Itoa(int(newArticle.ID))
	os.MkdirAll(path, 0755)
	fileName := path + "/" + image.Filename

	// Save file
	if err := ctx.SaveUploadedFile(image, fileName); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newArticle.Image = os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + fileName

	articles = append(articles, newArticle)
	ctx.JSON(http.StatusOK, gin.H{"articles": articles})
}
