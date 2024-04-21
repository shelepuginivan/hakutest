package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/app/server/controllers"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/shelepuginivan/hakutest/web"
)

func registerStatic(e *gin.Engine) {
	e.StaticFS("/static", http.FS(web.Static))
}

func registerTemplates(e *gin.Engine) {
	tmpl := template.Must(template.ParseFS(web.Templates, "templates/*.tmpl"))
	e.SetHTMLTemplate(tmpl)
}

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

func NewRouter(t *test.TestService, r *results.ResultsService) *gin.Engine {
	setMode(config.New().Server.Mode)

	router := gin.New()

	registerStatic(router)
	registerTemplates(router)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	editor := controllers.NewEditorController()
	search := controllers.NewSearchController(t)
	test := controllers.NewTestController(t, r)

	router.GET("/", search.SearchPage)
	router.GET("/editor/upload", editor.ChooseTest)
	router.GET("/editor/edit", editor.NewTest)
	router.POST("/editor/edit", editor.UploadTest)
	router.POST("/editor/create", editor.CreateTest)
	router.GET("/:test", test.GetTest)
	router.POST("/:test", test.SubmitTest)

	return router
}
