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

func NewEngine(t *test.TestService, r *results.ResultsService) *gin.Engine {
	setMode(config.New().Server.Mode)

	engine := gin.New()

	registerStatic(engine)
	registerTemplates(engine)

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	editor := controllers.NewEditorController()
	search := controllers.NewSearchController(t)
	test := controllers.NewTestController(t, r)

	engine.GET("/", search.SearchPage)
	engine.GET("/editor/upload", editor.ChooseTest)
	engine.GET("/editor/edit", editor.NewTest)
	engine.POST("/editor/edit", editor.UploadTest)
	engine.POST("/editor/create", editor.CreateTest)
	engine.GET("/:test", test.GetTest)
	engine.POST("/:test", test.SubmitTest)

	return engine
}
