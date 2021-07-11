package main

import (
	"course-go/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Static("/uploads", "./uploads")
	uploadPath := []string{"articles", "users"}
	for _, dir := range uploadPath {
		os.MkdirAll("uploads/"+dir, 0755)
	}
	routes.Serve(r)

	r.Run()
}
