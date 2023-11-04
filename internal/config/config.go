package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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

func getConfigDir() string {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "hakutest"
	}

	return path.Join(configDir, "hakutest")

}

func Init() Config {
	configDir := getConfigDir()
	configPath := path.Join(configDir, "config.yaml")

	testsDirectory := "user_test"
	resultsDirectory := "user_results"

	cacheDir, err := os.UserCacheDir()

	if err == nil {
		testsDirectory = path.Join(cacheDir, "hakutest", "tests")
		resultsDirectory = path.Join(cacheDir, "hakutest", "results")
	}

	config := Config{
		General: GeneralConfig{
			TestsDirectory:   testsDirectory,
			ResultsDirectory: resultsDirectory,
		},
		Server: ServerConfig{
			Port: "8080",
		},
		Ui: UiConfig{
			Editor: UiEditorConfig{
				Header:                   "Test Editor",
				LabelTitle:               "Title:",
				LabelDescription:         "Description:",
				LabelSubject:             "Subject:",
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
				ErrorHeaderLabel:  "An error occurred!",
				ErrorDetailsLabel: "Details",
			},
			Test: UiTestConfig{
				StudentNameLabel:  "Your name:",
				OpenAnswerLabel:   "Answer:",
				SubmitButtonLabel: "Submit",
			},
		},
	}

	v := viper.New()

	v.AddConfigPath(configDir)
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = os.MkdirAll(path.Dir(configDir), 0770)

			if err == nil || os.IsExist(err) {
				data, err := yaml.Marshal(config)

				if err != nil {
					data = []byte{}
				}

				os.WriteFile(configPath, data, 0666)
			}
		}

		return config
	}

	v.Unmarshal(&config)
	return config
}
