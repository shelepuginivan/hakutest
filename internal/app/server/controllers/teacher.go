package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

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

// TeacherController is a controller that provides handlers for teacher routes.
type TeacherController struct {
	cfg *config.Config
}

// NewTeacher returns a new instance of [StudentController].
func NewTeacher(cfg *config.Config) *TeacherController {
	return &TeacherController{cfg: cfg}
}

// Index is a handler for the `GET /teacher` route.
//
// It redirects request to `/teacher/dashboard`.
func (co *TeacherController) Index(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/teacher/dashboard")
}

// Dashboard is a handler for the `GET /teacher/dashboard` route.
//
// It renders HTML template of the dashboard.
func (co *TeacherController) Dashboard(c *gin.Context) {
	localIP, err := network.GetLocalIP()
	if err == nil {
		localIP = fmt.Sprintf("http://%s:%d", localIP, co.cfg.General.Port)
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

// Tests is a handler for the `GET /teacher/tests` route.
//
// It renders HTML template of a test menu.
func (co *TeacherController) Tests(c *gin.Context) {
	c.HTML(http.StatusOK, "tests.gohtml", gin.H{
		"Tests": test.GetList(),
	})
}

// DownloadSelected is a handler for the `GET /teacher/tests/selected` route.
//
// It writes `.zip` archive with selected test to the response.
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

// DeleteSelected is a handler for the `POST /teacher/tests/selected` route.
//
// It accepts form (`application/x-www-form-urlencoded`) and deletes selected
// tests. The request is redirected to the `/teacher/tests` page.
func (co *TeacherController) DeleteSelected(c *gin.Context) {
	selected := c.PostFormArray("tests")
	test.DeleteMany(selected...)
	c.Redirect(http.StatusSeeOther, "/teacher/tests")
}

// ImportTests is a handler for the `POST /teacher/tests/import` route.
//
// It accepts form (`multipart/form-data`) and imports uploaded tests to the
// tests directory. The request is redirected to the `/teacher/tests` page.
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

// DownloadTest is a handler for the `GET /teacher/tests/action/:test` route.
//
// It writes the test as attachment in JSON format.
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

// DeleteTest is a handler for the `POST /teacher/tests/action/:test` route.
//
// It accepts form (`application/x-www-form-urlencoded`) and deletes specified
// test. The request is redirected to `/teacher/tests` page.
func (co *TeacherController) DeleteTest(c *gin.Context) {
	testName := c.Param("test")
	test.DeleteMany(testName)
	c.Redirect(http.StatusSeeOther, "/teacher/tests")
}

// TestEditor is a handler for the `GET /teacher/tests/edit` route.
//
// It renders HTML template of a test editor.
func (co *TeacherController) TestEditor(c *gin.Context) {
	testName, ok := c.GetQuery("name")
	t, err := test.GetByName(testName)
	if !ok || err != nil {
		t = &test.Test{}
	}

	c.HTML(http.StatusOK, "editor.gohtml", t)
}

// SubmitTest is a handler for `POST /teacher/tests/edit` route.
//
// It accepts JSON (`application/json`) and adds test bound to it.
//
// If test is valid and added sucessfully, SubmitTest responds with status `201
// Created`, otherwise it returns JSON with field `message` explaining occurred
// error.
func (co *TeacherController) SubmitTest(c *gin.Context) {
	var t test.Test

	err := c.BindJSON(&t)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": i18n.Get("err.unprocessable.text"),
		})
		return
	}

	err = test.Save(&t)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": i18n.Get("err.write.text"),
		})
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}

// Statistics is a handler for the `GET /teacher/statistics` route.
//
// If search parameter `q` is present, it renders HTML template of the results
// statistics, otherwise statistics menu is rendered.
func (co *TeacherController) Statistics(c *gin.Context) {
	resultName, ok := c.GetQuery("q")
	if !ok {
		c.HTML(http.StatusOK, "statistics_menu.gohtml", gin.H{
			"AvailableResults": result.AvailableResults(),
		})
		return
	}

	stats, err := statistics.NewFromSaved(resultName)
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

// StatisticsExport is a handler for the `GET /teacher/statistics/export`
// route.
//
// If search parameter `name` is present, it exports statistics in the format
// `format` (JSON by default) and writes it as response, otherwise
// StatisticsExport redirects request to the `/teacher/statistics`.
func (co *TeacherController) StatisticsExport(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/teacher/statistics")
		return
	}

	stats, err := statistics.NewFromSaved(name)
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

// SettingsPage is a handler for the `GET /teacher/settings` route.
//
// It renders HTML template of a settings page.
func (co *TeacherController) SettingsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.gohtml", gin.H{
		"Config":         co.cfg,
		"SupportedLangs": i18n.SupportedLangs(),
	})
}

// SettingsUpdate is a handler for the `POST /teacher/settings` route.
//
// It accepts JSON (`application/json`) and updates the application settings
// with the values provided in JSON format.
//
// If configuration is updated successfully, it responds with status `201
// Created`, otherwise it returns JSON with field `message` explaining occurred
// error.
func (co *TeacherController) SettingsUpdate(c *gin.Context) {
	var fields config.Fields
	err := c.BindJSON(&fields)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": i18n.Get("err.unprocessable.text"),
		})
		return
	}

	err = co.cfg.Update(func(_ config.Fields) config.Fields {
		return fields
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": i18n.Get("err.write.text"),
		})
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
