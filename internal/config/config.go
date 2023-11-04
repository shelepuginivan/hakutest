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

type GeneralConfig struct {
	TestsDirectory   string `yaml:"tests_directory"`
	ResultsDirectory string `yaml:"results_directory"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type UiTestConfig struct {
	StudentNameLabel  string `yaml:"student_name_label"`
	OpenAnswerLabel   string `yaml:"open_answer_label"`
	SubmitButtonLabel string `yaml:"submit_button_label"`
}

type UiErrorConfig struct {
	ErrorHeaderLabel  string `yaml:"error_header_label"`
	ErrorDetailsLabel string `yaml:"error_details_label"`
}

type UiConfig struct {
	Test  UiTestConfig  `yaml:"test"`
	Error UiErrorConfig `yaml:"error"`
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
	var (
		port              = "8080"
		testsDirectory    = "user_test"
		resultsDirectory  = "user_results"
		studentNameLabel  = "Your name:"
		openAnswerLabel   = "Answer:"
		submitButtonLabel = "Submit"
		errorHeaderLabel  = "An error occurred!"
		errorDetailsLabel = "Details"
	)

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
			Test: UiTestConfig{
				StudentNameLabel:  studentNameLabel,
				OpenAnswerLabel:   openAnswerLabel,
				SubmitButtonLabel: submitButtonLabel,
			},
			Error: UiErrorConfig{
				ErrorHeaderLabel:  errorHeaderLabel,
				ErrorDetailsLabel: errorDetailsLabel,
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

	if config.Ui.Test.StudentNameLabel == "" {
		config.Ui.Test.StudentNameLabel = studentNameLabel
	}

	if config.Ui.Test.OpenAnswerLabel == "" {
		config.Ui.Test.OpenAnswerLabel = openAnswerLabel
	}

	if config.Ui.Test.SubmitButtonLabel == "" {
		config.Ui.Test.SubmitButtonLabel = submitButtonLabel
	}

	if config.Ui.Error.ErrorHeaderLabel == "" {
		config.Ui.Error.ErrorHeaderLabel = errorHeaderLabel
	}

	if config.Ui.Error.ErrorDetailsLabel == "" {
		config.Ui.Error.ErrorDetailsLabel = errorDetailsLabel
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
