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
