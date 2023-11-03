package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/app/server/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "web/static")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	editor := new(controllers.EditorController)
	test := new(controllers.TestController)

	router.GET("/editor/upload", editor.ChooseTest)
	router.GET("/editor/edit", editor.NewTest)
	router.POST("/editor/edit", editor.UploadTest)
	router.POST("/editor/create", editor.CreateTest)
	router.GET("/:test", test.GetTest)
	router.POST("/:test", test.SubmitTest)

	return router
}
