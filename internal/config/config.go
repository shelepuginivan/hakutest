package config

import (
	"os"
	"path/filepath"

	"github.com/shelepuginivan/hakutest/internal/pkg/utils"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type GeneralConfig struct {
	TestsDirectory   string `yaml:"tests_directory" mapstructure:"tests_directory"`
	ResultsDirectory string `yaml:"results_directory" mapstructure:"results_directory"`
	ShowResults      bool   `yaml:"show_results" mapstructure:"show_results"`
}

type ServerConfig struct {
	Port int    `yaml:"port" mapstructure:"port"`
	Mode string `yaml:"mode" mapstructure:"mode"`
}

type ExcelConfig struct {
	TestResultsSheet    string `yaml:"test_results_sheet" mapstructure:"test_results_sheet"`
	TestStatisticsSheet string `yaml:"statistics_sheet" mapstructure:"test_statistics_sheet"`
	HeaderStudent       string `yaml:"header_student" mapstructure:"header_student"`
	HeaderPoints        string `yaml:"header_points" mapstructure:"header_points"`
	HeaderPercentage    string `yaml:"header_percentage" mapstructure:"header_percentage"`
}

type ImageConfig struct {
	Title  string `yaml:"title" mapstructure:"title"`
	LabelX string `yaml:"label_x" mapstructure:"label_x"`
	LabelY string `yaml:"label_y" mapstructure:"label_y"`
}

type StatisticsConfig struct {
	Excel ExcelConfig `yaml:"excel" mapstructure:"excel"`
	Image ImageConfig `yaml:"image" mapstructure:"image"`
}

type UiEditorConfig struct {
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

type UiErrorConfig struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Details string `yaml:"details" mapstructure:"details"`
}

type UiExpiredConfig struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Message string `yaml:"message" mapstructure:"message"`
}

type UiSearchConfig struct {
	InputPlaceholder  string `yaml:"input_placeholder" mapstructure:"input_placeholder"`
	SearchButtonLabel string `yaml:"search_button_label" mapstructure:"search_button_label"`
}

type UiSubmittedConfig struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Message string `yaml:"message" mapstructure:"message"`
}

type UiTestConfig struct {
	StudentNameLabel  string `yaml:"student_name_label" mapstructure:"student_name_label"`
	OpenAnswerLabel   string `yaml:"open_answer_label" mapstructure:"open_answer_label"`
	SubmitButtonLabel string `yaml:"submit_button_label" mapstructure:"submit_button_label"`
}

type UiConfig struct {
	Editor    UiEditorConfig    `yaml:"editor" mapstructure:"editor"`
	Error     UiErrorConfig     `yaml:"error" mapstructure:"error"`
	Expired   UiExpiredConfig   `yaml:"expired" mapstructure:"expired"`
	Search    UiSearchConfig    `yaml:"search" mapstructure:"search"`
	Submitted UiSubmittedConfig `yaml:"submitted" mapstructure:"submitted"`
	Test      UiTestConfig      `yaml:"test" mapstructure:"test"`
}

type Config struct {
	General    GeneralConfig    `yaml:"general" mapstructure:"general"`
	Server     ServerConfig     `yaml:"server" mapstructure:"server"`
	Statistics StatisticsConfig `yaml:"stats" mapstructure:"stats"`
	Ui         UiConfig         `yaml:"ui" mapstructure:"ui"`
}

func getConfigDir() string {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "hakutest"
	}

	return filepath.Join(configDir, "hakutest")
}

func getViper() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(getConfigDir())
	v.AddConfigPath(utils.GetExecutablePath())
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	return v
}

func Default() Config {
	testsDirectory := "user_test"
	resultsDirectory := "user_results"

	cacheDir, err := os.UserCacheDir()

	if err == nil {
		testsDirectory = filepath.Join(cacheDir, "hakutest", "tests")
		resultsDirectory = filepath.Join(cacheDir, "hakutest", "results")
	}

	defaultConfig := Config{
		General: GeneralConfig{
			TestsDirectory:   testsDirectory,
			ResultsDirectory: resultsDirectory,
			ShowResults:      true,
		},
		Server: ServerConfig{
			Port: 8080,
			Mode: "release",
		},
		Statistics: StatisticsConfig{
			Excel: ExcelConfig{
				TestResultsSheet:    "Test Results",
				TestStatisticsSheet: "Test Statistics",
				HeaderStudent:       "Student",
				HeaderPoints:        "Points",
				HeaderPercentage:    "%",
			},
			Image: ImageConfig{
				Title:  "Student Performance",
				LabelX: "Points",
				LabelY: "Students",
			},
		},
		Ui: UiConfig{
			Editor: UiEditorConfig{
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
			Error: UiErrorConfig{
				Header:  "An error occurred!",
				Details: "Details",
			},
			Expired: UiExpiredConfig{
				Header:  "Test expired!",
				Message: "This test is no longer available",
			},
			Search: UiSearchConfig{
				InputPlaceholder:  "Search for a test",
				SearchButtonLabel: "Search",
			},
			Submitted: UiSubmittedConfig{
				Header:  "Submitted!",
				Message: "The test results are not displayed according to the system settings",
			},
			Test: UiTestConfig{
				StudentNameLabel:  "Your name:",
				OpenAnswerLabel:   "Answer:",
				SubmitButtonLabel: "Submit",
			},
		},
	}

	return defaultConfig
}

func New() Config {
	config := Default()
	configDir := getConfigDir()
	configPath := filepath.Join(configDir, "config.yaml")

	v := getViper()
	v.SetDefault("general", config.General)
	v.SetDefault("server", config.Server)
	v.SetDefault("stats", config.Statistics)
	v.SetDefault("ui", config.Ui)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err = os.MkdirAll(configDir, os.ModeDir|os.ModePerm); err != nil {
			panic(err)
		}

		file, err := os.Create(configPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		data, err := yaml.Marshal(config)

		if err != nil {
			data = []byte{}
		}

		file.Write(data)
	}

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
