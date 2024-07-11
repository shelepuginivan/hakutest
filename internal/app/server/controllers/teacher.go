package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/network"
	"github.com/shelepuginivan/hakutest/internal/pkg/uptime"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/statistics"
	"github.com/shelepuginivan/hakutest/pkg/test"
	"github.com/shelepuginivan/hakutest/pkg/version"
)

// TeacherController is a [github.com/gin-gonic/gin] controller that provides
// handlers for teacher routes.
type TeacherController struct {
	cfg *config.Config
}

// NewTeacher returns a new instance of [StudentController].
func NewTeacher(cfg *config.Config) *TeacherController {
	return &TeacherController{cfg: cfg}
}

// Index is a [github.com/gin-gonic/gin] handler for the `GET /teacher` route.
// It redirects request to `/teacher/dashboard`.
func (co *TeacherController) Index(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/teacher/dashboard")
}

// Dashboard is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/dashboard` route. It renders teacher dashboard HTML page.
func (co *TeacherController) Dashboard(c *gin.Context) {
	localIP, err := network.GetLocalIP()
	if err == nil {
		localIP = fmt.Sprintf("http://%s:%d", localIP, co.cfg.Port)
	}

	u := uptime.Uptime()

	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"AddressAvailable": err == nil,
		"Address":          localIP,
		"Version":          version.Version,
		"Uptime": gin.H{
			"Hours":   int(u.Hours()),
			"Minutes": int(u.Minutes()) % 60,
			"Seconds": int(u.Seconds()) % 60,
		},
	})
}

// Tests is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/tests` route. It renders tests menu HTML page.
func (co *TeacherController) Tests(c *gin.Context) {
	c.HTML(http.StatusOK, "tests.gohtml", gin.H{
		"Tests": test.GetList(),
	})
}

// DownloadSelected is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/tests/selected` route. It writes `.zip` archive with each selected
// test.
func (co *TeacherController) DownloadSelected(c *gin.Context) {
	selected := c.QueryArray("tests")
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=hakutest.zip")

	err := test.WriteZip(c.Writer, selected...)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Error(err)
	}
}

// DeleteSelected is a [github.com/gin-gonic/gin] handler for the `POST
// /teacher/tests/selected` route. It deletes every selected test and redirects
// request to the `/teacher/tests` page.
func (co *TeacherController) DeleteSelected(c *gin.Context) {
	selected := c.PostFormArray("tests")
	test.DeleteMany(selected...)
	c.Redirect(http.StatusSeeOther, "/teacher/tests")
}

// ImportTests is a [github.com/gin-gonic/gin] handler for the `POST
// /teacher/tests/import` route. It imports uploaded tests to the tests
// directory and redirects request to the `/teacher/tests` page.
func (co *TeacherController) ImportTests(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
			"Title":   i18n.Get("err.unprocessable.title"),
			"Text":    i18n.Get("err.unprocessable.text"),
			"Code":    http.StatusUnprocessableEntity,
			"Message": "failed to get uploaded files",
			"Error":   err.Error(),
		})
		return
	}

	files := form.File["files"]

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			continue
		}
		defer f.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, f); err != nil {
			continue
		}

		test.Import(buf.Bytes())
	}

	c.Redirect(http.StatusSeeOther, "/teacher/tests")
}

// DownloadSelected is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/tests/action/:test` route. It writes the test as attachment in JSON
// format.
func (co *TeacherController) DownloadTest(c *gin.Context) {
	testName := c.Param("test")

	c.Header("Content-Type", "application/json")
	c.Header(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=%s.json", testName),
	)

	err := test.WriteJSON(c.Writer, testName)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Error(err)
	}
}

// Statistics is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/statistics` route. If search parameter `q` is present, it renders
// HTML view of the results statistics, otherwise it renders statistics menu.
func (co *TeacherController) Statistics(c *gin.Context) {
	resultName, ok := c.GetQuery("q")
	if !ok {
		c.HTML(http.StatusOK, "statistics_menu.gohtml", gin.H{
			"AvailableResults": result.AvailableResults(),
		})
		return
	}

	stats, err := statistics.NewFromName(resultName)
	if err != nil {
		c.HTML(http.StatusNotFound, "info.gohtml", gin.H{
			"Title": i18n.Get("err.not_found.title"),
			"Text":  i18n.Get("err.not_found.text"),
		})
		return
	}

	c.HTML(http.StatusOK, "statistics.gohtml", gin.H{
		"Stats": stats,
		"ExportFormats": map[string]string{
			statistics.FormatXLSX: statistics.DescriptionXLSX,
			statistics.FormatCSV:  statistics.DescriptionCSV,
			statistics.FormatJSON: statistics.DescriptionJSON,
		},
	})
}

// StatisticsExport is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/statistics/export` route.
//
// If search parameter `name` is not present, it redirects request to the
// `/teacher/statistics`.
//
// If search parameter `name` is present, it exports statistics in the format
// `format` (JSON by default) and writes it as response.
func (co *TeacherController) StatisticsExport(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/teacher/statistics")
		return
	}

	stats, err := statistics.NewFromName(name)
	if err != nil {
		c.HTML(http.StatusNotFound, "info.gohtml", gin.H{
			"Title": i18n.Get("err.not_found.title"),
			"Text":  i18n.Get("err.not_found.text"),
		})
		return
	}

	format, ok := c.GetQuery("format")
	if !ok {
		format = statistics.FormatJSON
	}

	c.Header(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=%s.%s", name, format),
	)

	switch format {
	case statistics.FormatCSV:
		c.Header("Content-Type", statistics.MimeCSV)
		err = stats.WriteCSV(c.Writer)
	case statistics.FormatXLSX:
		c.Header("Content-Type", statistics.MimeXLSX)
		err = stats.WriteXLSX(c.Writer)
	default:
		c.Header("Content-Type", statistics.MimeJSON)
		err = stats.WriteJSON(c.Writer)
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
}

// SettingsPage is a [github.com/gin-gonic/gin] handler for the `GET
// /teacher/settings` route. It renders settings HTML page.
func (co *TeacherController) SettingsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.gohtml", gin.H{
		"Config":         co.cfg,
		"SupportedLangs": i18n.SupportedLangs(),
	})
}

// SettingsPage is a [github.com/gin-gonic/gin] handler for the `POST
// /teacher/settings` route. It updates the application settings with the
// values provided in the form.
//
// The configuration is updated dynamically, i.e. values are applied in-place,
// application restart is not required. The only exception is the server port.
func (co *TeacherController) SettingsUpdate(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
			"Title":   i18n.Get("err.unprocessable.title"),
			"Text":    i18n.Get("err.unprocessable.text"),
			"Code":    http.StatusUnprocessableEntity,
			"Message": "failed to parse settings form",
			"Error":   err.Error(),
		})
		return
	}

	fields := co.cfg.Fields

	fields.Debug = c.PostForm("debug") == "on"
	fields.DisableTray = c.PostForm("disable_tray") == "on"
	fields.Lang = c.PostForm("lang")
	fields.OverwriteResults = c.PostForm("overwrite_results") == "on"
	fields.ShowResults = c.PostForm("show_results") == "on"
	fields.TestsDirectory = c.PostForm("tests_directory")
	fields.ResultsDirectory = c.PostForm("results_directory")

	port, err := strconv.Atoi(c.PostForm("port"))
	if err == nil {
		fields.Port = port
	}

	err = co.cfg.Update(fields)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{
			"Title":   i18n.Get("err.write.title"),
			"Text":    i18n.Get("err.write.text"),
			"Code":    http.StatusUnprocessableEntity,
			"Message": "failed to write config file",
			"Error":   err.Error(),
		})
		return
	}

	c.HTML(http.StatusCreated, "settings.gohtml", gin.H{
		"Config":         co.cfg,
		"SupportedLangs": i18n.SupportedLangs(),
	})
}
