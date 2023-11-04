package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/rodaine/table"
	"gopkg.in/yaml.v3"
)

const (
	port                           = "8080"
	studentNameLabel               = "Your name:"
	openAnswerLabel                = "Answer:"
	submitButtonLabel              = "Submit"
	errorHeaderLabel               = "An error occurred!"
	errorDetailsLabel              = "Details"
	editorHeader                   = "Test Editor"
	editorLabelTitle               = "Title:"
	editorLabelDescription         = "Description:"
	editorLabelSubject             = "Subject:"
	editorLabelTarget              = "Target audience:"
	editorLabelInstitution         = "Institution:"
	editorLabelExpiresIn           = "Expires in:"
	editorLabelAddTask             = "+ Add task"
	editorLabelTaskHeader          = "Task"
	editorLabelTaskType            = "Type:"
	editorLabelTaskTypeSingle      = "Single answer"
	editorLabelTaskTypeMultiple    = "Multiple answers"
	editorLabelTaskTypeOpen        = "Open question"
	editorLabelTaskText            = "Text:"
	editorLabelTaskAnswer          = "Answer:"
	editorLabelTaskOptions         = "Answer options"
	editorLabelTaskAddOption       = "+ Add option"
	editorLabelAddAttachment       = "Add attachment"
	editorLabelAttachmentName      = "Name:"
	editorLabelAttachmentType      = "Type:"
	editorLabelAttachmentTypeFile  = "File"
	editorLabelAttachmentTypeImage = "Image"
	editorLabelAttachmentTypeVideo = "Video"
	editorLabelAttachmentTypeAudio = "Audio"
	editorLabelAttachmentSrc       = "Source (URL):"
	editorLabelUploadTestInput     = "Upload test file"
	editorLabelUploadTestButton    = "Upload and edit"
	editorLabelNewTest             = "Create new test"
	editorLabelDownloadTest        = "Download test"
)

type GeneralConfig struct {
	TestsDirectory   string `yaml:"tests_directory"`
	ResultsDirectory string `yaml:"results_directory"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type UiEditorConfig struct {
	Header                   string `yaml:"header"`
	LabelTitle               string `yaml:"label_title"`
	LabelDescription         string `yaml:"label_description"`
	LabelSubject             string `yaml:"label_subject"`
	LabelTarget              string `yaml:"label_target"`
	LabelInstitution         string `yaml:"label_institution"`
	LabelExpiresIn           string `yaml:"label_expires_in"`
	LabelAddTask             string `yaml:"label_add_task"`
	LabelTaskHeader          string `yaml:"label_task_header"`
	LabelTaskType            string `yaml:"label_task_type"`
	LabelTaskTypeSingle      string `yaml:"label_task_type_single"`
	LabelTaskTypeMultiple    string `yaml:"label_task_type_multiple"`
	LabelTaskTypeOpen        string `yaml:"label_task_type_open"`
	LabelTaskText            string `yaml:"label_task_text"`
	LabelTaskAnswer          string `yaml:"label_task_answer"`
	LabelTaskOptions         string `yaml:"label_task_options"`
	LabelTaskAddOption       string `yaml:"label_task_add_option"`
	LabelAddAttachment       string `yaml:"label_add_attachment"`
	LabelAttachmentName      string `yaml:"label_attachment_name"`
	LabelAttachmentType      string `yaml:"label_attachment_type"`
	LabelAttachmentTypeFile  string `yaml:"label_attachment_type_file"`
	LabelAttachmentTypeImage string `yaml:"label_attachment_type_image"`
	LabelAttachmentTypeVideo string `yaml:"label_attachment_type_video"`
	LabelAttachmentTypeAudio string `yaml:"label_attachment_type_audio"`
	LabelAttachmentSrc       string `yaml:"label_attachment_src"`
	LabelUploadTestInput     string `yaml:"label_upload_test_input"`
	LabelUploadTestButton    string `yaml:"label_upload_test_button"`
	LabelNewTest             string `yaml:"label_new_test"`
	LabelDownloadTest        string `yaml:"label_download_test"`
}

type UiErrorConfig struct {
	ErrorHeaderLabel  string `yaml:"error_header_label"`
	ErrorDetailsLabel string `yaml:"error_details_label"`
}

type UiTestConfig struct {
	StudentNameLabel  string `yaml:"student_name_label"`
	OpenAnswerLabel   string `yaml:"open_answer_label"`
	SubmitButtonLabel string `yaml:"submit_button_label"`
}

type UiConfig struct {
	Editor UiEditorConfig `yaml:"editor"`
	Error  UiErrorConfig  `yaml:"error"`
	Test   UiTestConfig   `yaml:"test"`
}

type Config struct {
	General GeneralConfig `yaml:"general"`
	Server  ServerConfig  `yaml:"server"`
	Ui      UiConfig      `yaml:"ui"`
}

func getConfigPath() string {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "config.yaml"
	}

	return path.Join(configDir, "hakutest", "config.yaml")
}

func Init() Config {
	testsDirectory := "user_test"
	resultsDirectory := "user_results"

	configPath := getConfigPath()
	config := Config{}
	cacheDir, err := os.UserCacheDir()

	if err == nil {
		testsDirectory = path.Join(cacheDir, "hakutest", "tests")
		resultsDirectory = path.Join(cacheDir, "hakutest", "results")
	}

	defaultConfig := Config{
		General: GeneralConfig{
			TestsDirectory:   testsDirectory,
			ResultsDirectory: resultsDirectory,
		},
		Server: ServerConfig{
			Port: port,
		},
		Ui: UiConfig{
			Editor: UiEditorConfig{
				Header:                   editorHeader,
				LabelTitle:               editorLabelTitle,
				LabelDescription:         editorLabelDescription,
				LabelSubject:             editorLabelSubject,
				LabelTarget:              editorLabelTarget,
				LabelInstitution:         editorLabelInstitution,
				LabelExpiresIn:           editorLabelExpiresIn,
				LabelAddTask:             editorLabelAddTask,
				LabelTaskHeader:          editorLabelTaskHeader,
				LabelTaskType:            editorLabelTaskType,
				LabelTaskTypeSingle:      editorLabelTaskTypeSingle,
				LabelTaskTypeMultiple:    editorLabelTaskTypeMultiple,
				LabelTaskTypeOpen:        editorLabelTaskTypeOpen,
				LabelTaskText:            editorLabelTaskText,
				LabelTaskAnswer:          editorLabelTaskAnswer,
				LabelTaskOptions:         editorLabelTaskOptions,
				LabelTaskAddOption:       editorLabelTaskAddOption,
				LabelAddAttachment:       editorLabelAddAttachment,
				LabelAttachmentName:      editorLabelAttachmentName,
				LabelAttachmentType:      editorLabelAttachmentType,
				LabelAttachmentTypeFile:  editorLabelAttachmentTypeFile,
				LabelAttachmentTypeImage: editorLabelAttachmentTypeImage,
				LabelAttachmentTypeVideo: editorLabelAttachmentTypeVideo,
				LabelAttachmentTypeAudio: editorLabelAttachmentTypeAudio,
				LabelAttachmentSrc:       editorLabelAttachmentSrc,
				LabelUploadTestInput:     editorLabelUploadTestInput,
				LabelUploadTestButton:    editorLabelUploadTestButton,
				LabelNewTest:             editorLabelNewTest,
				LabelDownloadTest:        editorLabelDownloadTest,
			},
			Error: UiErrorConfig{
				ErrorHeaderLabel:  errorHeaderLabel,
				ErrorDetailsLabel: errorDetailsLabel,
			},
			Test: UiTestConfig{
				StudentNameLabel:  studentNameLabel,
				OpenAnswerLabel:   openAnswerLabel,
				SubmitButtonLabel: submitButtonLabel,
			},
		},
	}

	configFile, err := os.ReadFile(configPath)

	if err != nil {
		return defaultConfig
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return defaultConfig
	}

	if config.General.TestsDirectory == "" {
		config.General.TestsDirectory = testsDirectory
	}

	if config.General.ResultsDirectory == "" {
		config.General.ResultsDirectory = testsDirectory
	}

	if config.Server.Port == "" {
		config.Server.Port = port
	}

	if config.Ui.Editor.Header == "" {
		config.Ui.Editor.Header = editorHeader
	}

	if config.Ui.Editor.LabelTitle == "" {
		config.Ui.Editor.LabelTitle = editorLabelTitle
	}

	if config.Ui.Editor.LabelDescription == "" {
		config.Ui.Editor.LabelDescription = editorLabelDescription
	}

	if config.Ui.Editor.LabelSubject == "" {
		config.Ui.Editor.LabelSubject = editorLabelSubject
	}

	if config.Ui.Editor.LabelTarget == "" {
		config.Ui.Editor.LabelTarget = editorLabelTarget
	}

	if config.Ui.Editor.LabelInstitution == "" {
		config.Ui.Editor.LabelInstitution = editorLabelInstitution
	}

	if config.Ui.Editor.LabelExpiresIn == "" {
		config.Ui.Editor.LabelExpiresIn = editorLabelExpiresIn
	}

	if config.Ui.Editor.LabelAddTask == "" {
		config.Ui.Editor.LabelAddTask = editorLabelAddTask
	}

	if config.Ui.Editor.LabelTaskHeader == "" {
		config.Ui.Editor.LabelTaskHeader = editorLabelTaskHeader
	}

	if config.Ui.Editor.LabelTaskType == "" {
		config.Ui.Editor.LabelTaskType = editorLabelTaskType
	}

	if config.Ui.Editor.LabelTaskTypeSingle == "" {
		config.Ui.Editor.LabelTaskTypeSingle = editorLabelTaskTypeSingle
	}

	if config.Ui.Editor.LabelTaskTypeMultiple == "" {
		config.Ui.Editor.LabelTaskTypeMultiple = editorLabelTaskTypeMultiple
	}

	if config.Ui.Editor.LabelTaskTypeOpen == "" {
		config.Ui.Editor.LabelTaskTypeOpen = editorLabelTaskTypeOpen
	}

	if config.Ui.Editor.LabelTaskText == "" {
		config.Ui.Editor.LabelTaskText = editorLabelTaskText
	}

	if config.Ui.Editor.LabelTaskAnswer == "" {
		config.Ui.Editor.LabelTaskAnswer = editorLabelTaskAnswer
	}

	if config.Ui.Editor.LabelTaskOptions == "" {
		config.Ui.Editor.LabelTaskOptions = editorLabelTaskOptions
	}

	if config.Ui.Editor.LabelTaskAddOption == "" {
		config.Ui.Editor.LabelTaskAddOption = editorLabelTaskAddOption
	}

	if config.Ui.Editor.LabelAddAttachment == "" {
		config.Ui.Editor.LabelAddAttachment = editorLabelAddAttachment
	}

	if config.Ui.Editor.LabelAttachmentName == "" {
		config.Ui.Editor.LabelAttachmentName = editorLabelAttachmentName
	}

	if config.Ui.Editor.LabelAttachmentType == "" {
		config.Ui.Editor.LabelAttachmentType = editorLabelAttachmentType
	}

	if config.Ui.Editor.LabelAttachmentTypeFile == "" {
		config.Ui.Editor.LabelAttachmentTypeFile = editorLabelAttachmentTypeFile
	}

	if config.Ui.Editor.LabelAttachmentTypeImage == "" {
		config.Ui.Editor.LabelAttachmentTypeImage = editorLabelAttachmentTypeImage
	}

	if config.Ui.Editor.LabelAttachmentTypeVideo == "" {
		config.Ui.Editor.LabelAttachmentTypeVideo = editorLabelAttachmentTypeVideo
	}

	if config.Ui.Editor.LabelAttachmentTypeAudio == "" {
		config.Ui.Editor.LabelAttachmentTypeAudio = editorLabelAttachmentTypeAudio
	}

	if config.Ui.Editor.LabelAttachmentSrc == "" {
		config.Ui.Editor.LabelAttachmentSrc = editorLabelAttachmentSrc
	}

	if config.Ui.Editor.LabelUploadTestInput == "" {
		config.Ui.Editor.LabelUploadTestInput = editorLabelUploadTestInput
	}

	if config.Ui.Editor.LabelUploadTestButton == "" {
		config.Ui.Editor.LabelUploadTestButton = editorLabelUploadTestButton
	}

	if config.Ui.Editor.LabelNewTest == "" {
		config.Ui.Editor.LabelNewTest = editorLabelNewTest
	}

	if config.Ui.Editor.LabelDownloadTest == "" {
		config.Ui.Editor.LabelDownloadTest = editorLabelDownloadTest
	}

	if config.Ui.Error.ErrorHeaderLabel == "" {
		config.Ui.Error.ErrorHeaderLabel = errorHeaderLabel
	}

	if config.Ui.Error.ErrorDetailsLabel == "" {
		config.Ui.Error.ErrorDetailsLabel = errorDetailsLabel
	}

	if config.Ui.Test.StudentNameLabel == "" {
		config.Ui.Test.StudentNameLabel = studentNameLabel
	}

	if config.Ui.Test.OpenAnswerLabel == "" {
		config.Ui.Test.OpenAnswerLabel = openAnswerLabel
	}

	if config.Ui.Test.SubmitButtonLabel == "" {
		config.Ui.Test.SubmitButtonLabel = submitButtonLabel
	}

	return config
}

func (c Config) Print() {
	fmt.Println("genaral")
	c.PrintField("general")
	fmt.Println()
	fmt.Println("server")
	c.PrintField("server")
	fmt.Println()
	fmt.Println("ui")
	c.PrintField("ui")
}

func (c Config) PrintField(field string) {
	switch field {
	case "general":
		tbl := table.New("Key", "Value")
		tbl.AddRow("tests_directory", c.General.TestsDirectory)
		tbl.AddRow("results_directory", c.General.ResultsDirectory)
		tbl.Print()
	case "general.tests_directory":
		fmt.Println(c.General.TestsDirectory)
	case "general.results_directory":
		fmt.Println(c.General.ResultsDirectory)
	case "server":
		tbl := table.New("Key", "Value")
		tbl.AddRow("port", c.Server.Port)
		tbl.Print()
	case "server.port":
		fmt.Println(c.Server.Port)
	case "ui":
		tbl := table.New("Key", "Value")
		tbl.AddRow("error.error_header_label", c.Ui.Error.ErrorHeaderLabel)
		tbl.AddRow("error.error_details_label", c.Ui.Error.ErrorDetailsLabel)
		tbl.AddRow("test.open_answer_label", c.Ui.Test.OpenAnswerLabel)
		tbl.AddRow("test.student_name_label", c.Ui.Test.StudentNameLabel)
		tbl.AddRow("test.submit_button_label", c.Ui.Test.SubmitButtonLabel)
		tbl.Print()
	case "ui.error":
		tbl := table.New("Key", "Value")
		tbl.AddRow("error_header_label", c.Ui.Error.ErrorHeaderLabel)
		tbl.AddRow("error_details_label", c.Ui.Error.ErrorDetailsLabel)
		tbl.Print()
	case "ui.error.error_header_label":
		fmt.Println(c.Ui.Error.ErrorHeaderLabel)
	case "ui.error.error_details_label":
		fmt.Println(c.Ui.Error.ErrorDetailsLabel)
	case "ui.test":
		tbl := table.New("Key", "Value")
		tbl.AddRow("open_answer_label", c.Ui.Test.OpenAnswerLabel)
		tbl.AddRow("student_name_label", c.Ui.Test.StudentNameLabel)
		tbl.AddRow("submit_button_label", c.Ui.Test.SubmitButtonLabel)
		tbl.Print()
	case "ui.test.open_answer_label":
		fmt.Println(c.Ui.Test.OpenAnswerLabel)
	case "ui.test.student_name_label":
		fmt.Println(c.Ui.Test.StudentNameLabel)
	case "ui.test.submit_button_label":
		fmt.Println(c.Ui.Test.SubmitButtonLabel)
	default:
		log.Fatal("unknown field:", field)
	}
}

func (c *Config) SetField(field, value string) error {
	switch field {
	case "general.tests_directory":
		c.General.TestsDirectory = value
	case "general.results_directory":
		c.General.ResultsDirectory = value
	case "server.port":
		c.Server.Port = value
	case "ui.error.error_header_label":
		c.Ui.Error.ErrorHeaderLabel = value
	case "ui.error.error_details_label":
		c.Ui.Error.ErrorDetailsLabel = value
	case "ui.test.open_answer_label":
		c.Ui.Test.OpenAnswerLabel = value
	case "ui.test.student_name_label":
		c.Ui.Test.StudentNameLabel = value
	case "ui.test.submit_button_label":
		c.Ui.Test.SubmitButtonLabel = value
	case "general", "server", "ui", "ui.error", "ui.test":
		return errors.New("can only set primitive values")
	default:
		return errors.New("unknown field: " + field)
	}

	return nil
}
