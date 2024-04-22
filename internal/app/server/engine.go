package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/app/server/controllers"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
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

func NewEngine(app *application.App) *gin.Engine {
	setMode(app.Config.Server.Mode)

	engine := gin.New()

	registerStatic(engine)
	registerTemplates(engine)

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	testService := test.NewService(app)
	resultsService := results.NewService(app)

	editor := controllers.NewEditorController(app)
	search := controllers.NewSearchController(app, testService)
	test := controllers.NewTestController(app, testService, resultsService)

	engine.GET("/", search.SearchPage)
	engine.GET("/editor/upload", editor.ChooseTest)
	engine.GET("/editor/edit", editor.NewTest)
	engine.POST("/editor/edit", editor.UploadTest)
	engine.POST("/editor/create", editor.CreateTest)
	engine.GET("/:test", test.GetTest)
	engine.POST("/:test", test.SubmitTest)

	return engine
}
