package i18n

import (
	"os"
	"path/filepath"

	"github.com/shelepuginivan/hakutest/internal/pkg/directories"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type ServerI18n struct {
	StopTitle   string `yaml:"stop_title" mapstructure:"stop_title"`
	StopTooltip string `yaml:"stop_tooltip" mapstructure:"stop_tooltip"`
}

type StatsAppI18n struct {
	LabelTest      string `yaml:"label_test" mapstructure:"label_test"`
	LabelFormat    string `yaml:"label_format" mapstructure:"label_format"`
	LabelDirectory string `yaml:"label_directory" mapstructure:"label_directory"`
	SubmitText     string `yaml:"submit_text" mapstructure:"submit_text"`
	CancelText     string `yaml:"cancel_text" mapstructure:"cancel_text"`
	SelectText     string `yaml:"select_text" mapstructure:"select_text"`
	SuccessText    string `yaml:"success_text" mapstructure:"success_text"`
	ErrorPrefix    string `yaml:"error_prefix" mapstructure:"error_prefix"`
}

type StatsExcelI18n struct {
	TestResultsSheet    string `yaml:"test_results_sheet" mapstructure:"test_results_sheet"`
	TestStatisticsSheet string `yaml:"statistics_sheet" mapstructure:"statistics_sheet"`
	HeaderStudent       string `yaml:"header_student" mapstructure:"header_student"`
	HeaderPoints        string `yaml:"header_points" mapstructure:"header_points"`
	HeaderPercentage    string `yaml:"header_percentage" mapstructure:"header_percentage"`
}

type StatsImageI18n struct {
	Title  string `yaml:"title" mapstructure:"title"`
	LabelX string `yaml:"label_x" mapstructure:"label_x"`
	LabelY string `yaml:"label_y" mapstructure:"label_y"`
}

type StatsI18n struct {
	App   StatsAppI18n   `yaml:"app" mapstructure:"app"`
	Excel StatsExcelI18n `yaml:"excel" mapstructure:"excel"`
	Image StatsImageI18n `yaml:"image" mapstructure:"image"`
}

type WebEditorI18n struct {
	Header                   string `yaml:"header" mapstructure:"header"`
	LabelTitle               string `yaml:"label_title" mapstructure:"label_title"`
	LabelDescription         string `yaml:"label_description" mapstructure:"label_description"`
	LabelSubject             string `yaml:"label_subject" mapstructure:"label_subject"`
	LabelAuthor              string `yaml:"label_author" mapstructure:"label_author"`
	LabelTarget              string `yaml:"label_target" mapstructure:"label_target"`
	LabelInstitution         string `yaml:"label_institution" mapstructure:"label_institution"`
	LabelExpiresIn           string `yaml:"label_expires_in" mapstructure:"label_expires_in"`
	LabelAddTask             string `yaml:"label_add_task" mapstructure:"label_add_task"`
	LabelTaskHeader          string `yaml:"label_task_header" mapstructure:"label_task_header"`
	LabelTaskType            string `yaml:"label_task_type" mapstructure:"label_task_type"`
	LabelTaskTypeSingle      string `yaml:"label_task_type_single" mapstructure:"label_task_type_single"`
	LabelTaskTypeMultiple    string `yaml:"label_task_type_multiple" mapstructure:"label_task_type_multiple"`
	LabelTaskTypeOpen        string `yaml:"label_task_type_open" mapstructure:"label_task_type_open"`
	LabelTaskText            string `yaml:"label_task_text" mapstructure:"label_task_text"`
	LabelTaskAnswer          string `yaml:"label_task_answer" mapstructure:"label_task_answer"`
	LabelTaskOptions         string `yaml:"label_task_options" mapstructure:"label_task_options"`
	LabelTaskAddOption       string `yaml:"label_task_add_option" mapstructure:"label_task_add_option"`
	LabelAddAttachment       string `yaml:"label_add_attachment" mapstructure:"label_add_attachment"`
	LabelAttachmentName      string `yaml:"label_attachment_name" mapstructure:"label_attachment_name"`
	LabelAttachmentType      string `yaml:"label_attachment_type" mapstructure:"label_attachment_type"`
	LabelAttachmentTypeFile  string `yaml:"label_attachment_type_file" mapstructure:"label_attachment_type_file"`
	LabelAttachmentTypeImage string `yaml:"label_attachment_type_image" mapstructure:"label_attachment_type_image"`
	LabelAttachmentTypeVideo string `yaml:"label_attachment_type_video" mapstructure:"label_attachment_type_video"`
	LabelAttachmentTypeAudio string `yaml:"label_attachment_type_audio" mapstructure:"label_attachment_type_audio"`
	LabelAttachmentSrc       string `yaml:"label_attachment_src" mapstructure:"label_attachment_src"`
	LabelUploadTestInput     string `yaml:"label_upload_test_input" mapstructure:"label_upload_test_input"`
	LabelUploadTestButton    string `yaml:"label_upload_test_button" mapstructure:"label_upload_test_button"`
	LabelNewTest             string `yaml:"label_new_test" mapstructure:"label_new_test"`
	LabelDownloadTest        string `yaml:"label_download_test" mapstructure:"label_download_test"`
}

type WebErrorI18n struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Details string `yaml:"details" mapstructure:"details"`
}

type WebExpiredI18n struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Message string `yaml:"message" mapstructure:"message"`
}

type WebSearchI18n struct {
	InputPlaceholder  string `yaml:"input_placeholder" mapstructure:"input_placeholder"`
	SearchButtonLabel string `yaml:"search_button_label" mapstructure:"search_button_label"`
}

type WebSubmittedI18n struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Message string `yaml:"message" mapstructure:"message"`
}

type WebTestI18n struct {
	StudentNameLabel  string `yaml:"student_name_label" mapstructure:"student_name_label"`
	OpenAnswerLabel   string `yaml:"open_answer_label" mapstructure:"open_answer_label"`
	SubmitButtonLabel string `yaml:"submit_button_label" mapstructure:"submit_button_label"`
}

type WebI18n struct {
	Editor    WebEditorI18n    `yaml:"editor" mapstructure:"editor"`
	Error     WebErrorI18n     `yaml:"error" mapstructure:"error"`
	Expired   WebExpiredI18n   `yaml:"expired" mapstructure:"expired"`
	Search    WebSearchI18n    `yaml:"search" mapstructure:"search"`
	Submitted WebSubmittedI18n `yaml:"submitted" mapstructure:"submitted"`
	Test      WebTestI18n      `yaml:"test" mapstructure:"test"`
}

type I18n struct {
	Server     ServerI18n `yaml:"server" mapstructure:"server"`
	Statistics StatsI18n  `yaml:"stats" mapstructure:"stats"`
	Web        WebI18n    `yaml:"web" mapstructure:"web"`
}

func getViper() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(directories.Executable())
	v.AddConfigPath(directories.Config())
	v.SetConfigType("yaml")
	v.SetConfigName("i18n")

	return v
}

func Default() I18n {
	return I18n{
		Server: ServerI18n{
			StopTitle:   "Stop Hakutest",
			StopTooltip: "Stop Hakutest server and quit",
		},
		Statistics: StatsI18n{
			App: StatsAppI18n{
				LabelTest:      "Test",
				LabelFormat:    "Format",
				LabelDirectory: "Export to",
				SubmitText:     "Export",
				CancelText:     "Cancel",
				SelectText:     "(Select one)",
				SuccessText:    "Statistics exported successfully!",
				ErrorPrefix:    "An error occurred! Detail:",
			},
			Excel: StatsExcelI18n{
				TestResultsSheet:    "Test Results",
				TestStatisticsSheet: "Test Statistics",
				HeaderStudent:       "Student",
				HeaderPoints:        "Points",
				HeaderPercentage:    "%",
			},
			Image: StatsImageI18n{
				Title:  "Student Performance",
				LabelX: "Points",
				LabelY: "Students",
			},
		},
		Web: WebI18n{
			Editor: WebEditorI18n{
				Header:                   "Test Editor",
				LabelTitle:               "Title:",
				LabelDescription:         "Description:",
				LabelSubject:             "Subject:",
				LabelAuthor:              "Author:",
				LabelTarget:              "Target audience:",
				LabelInstitution:         "Institution:",
				LabelExpiresIn:           "Expires in:",
				LabelAddTask:             "+ Add task",
				LabelTaskHeader:          "Task",
				LabelTaskType:            "Type:",
				LabelTaskTypeSingle:      "Single answer",
				LabelTaskTypeMultiple:    "Multiple answers",
				LabelTaskTypeOpen:        "Open question",
				LabelTaskText:            "Text:",
				LabelTaskAnswer:          "Answer:",
				LabelTaskOptions:         "Answer options",
				LabelTaskAddOption:       "+ Add option",
				LabelAddAttachment:       "Add attachment",
				LabelAttachmentName:      "Name:",
				LabelAttachmentType:      "Type:",
				LabelAttachmentTypeFile:  "File",
				LabelAttachmentTypeImage: "Image",
				LabelAttachmentTypeVideo: "Video",
				LabelAttachmentTypeAudio: "Audio",
				LabelAttachmentSrc:       "Source (URL):",
				LabelUploadTestInput:     "Upload test file",
				LabelUploadTestButton:    "Upload and edit",
				LabelNewTest:             "Create new test",
				LabelDownloadTest:        "Download test",
			},
			Error: WebErrorI18n{
				Header:  "An error occurred!",
				Details: "Details",
			},
			Expired: WebExpiredI18n{
				Header:  "Test expired!",
				Message: "This test is no longer available",
			},
			Search: WebSearchI18n{
				InputPlaceholder:  "Search for a test",
				SearchButtonLabel: "Search",
			},
			Submitted: WebSubmittedI18n{
				Header:  "Submitted!",
				Message: "The test results are not displayed according to the system settings",
			},
			Test: WebTestI18n{
				StudentNameLabel:  "Your name:",
				OpenAnswerLabel:   "Answer:",
				SubmitButtonLabel: "Submit",
			},
		},
	}
}

func createDefaultI18n() error {
	configDir := directories.Config()
	configPath := filepath.Join(configDir, "i18n.yaml")

	err := os.MkdirAll(configDir, os.ModeDir|os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.Create(configPath)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := yaml.Marshal(Default())

	if err != nil {
		data = []byte{}
	}

	_, err = file.Write(data)

	return err
}

func New() I18n {
	i18n := Default()

	v := getViper()
	v.SetDefault("stats", i18n.Statistics)
	v.SetDefault("web", i18n.Web)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}

		if err := createDefaultI18n(); err != nil {
			panic(err)
		}
	}

	if err := v.Unmarshal(&i18n); err != nil {
		panic(err)
	}

	return i18n
}
