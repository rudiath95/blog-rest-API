package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiath95/blog-rest-API/ini"
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
	ini.SyncDatabases()
}

func main() {
	r := gin.Default()
	//>>Start Route

	//>>End Route
	r.Run()
}
