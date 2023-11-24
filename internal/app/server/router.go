package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/app/server/controllers"
	"github.com/shelepuginivan/hakutest/internal/config"
)

func setMode(mode string) {
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

func NewRouter() *gin.Engine {
	setMode(config.Init().Server.Mode)

	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "web/static")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	editor := new(controllers.EditorController)
	search := new(controllers.SearchController)
	test := new(controllers.TestController)

	router.GET("/", search.SearchPage)
	router.GET("/editor/upload", editor.ChooseTest)
	router.GET("/editor/edit", editor.NewTest)
	router.POST("/editor/edit", editor.UploadTest)
	router.POST("/editor/create", editor.CreateTest)
	router.GET("/:test", test.GetTest)
	router.POST("/:test", test.SubmitTest)

	return router
}
