package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port              string `yaml:"port"`
	TestsDirectory    string `yaml:"tests_directory"`
	ResultsDirectory  string `yaml:"results_directory"`
	StudentNameLabel  string `yaml:"student_name_label"`
	OpenAnswerLabel   string `yaml:"open_answer_label"`
	SubmitButtonLabel string `yaml:"submit_button_label"`
	ErrorHeaderLabel  string `yaml:"error_header_label"`
	ErrorDetailsLabel string `yaml:"error_details_label"`
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
		Port:              port,
		TestsDirectory:    testsDirectory,
		ResultsDirectory:  resultsDirectory,
		StudentNameLabel:  studentNameLabel,
		OpenAnswerLabel:   openAnswerLabel,
		SubmitButtonLabel: submitButtonLabel,
		ErrorHeaderLabel:  errorHeaderLabel,
		ErrorDetailsLabel: errorDetailsLabel,
	}

	configFile, err := os.ReadFile(configPath)

	if err != nil {
		return defaultConfig
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return defaultConfig
	}

	if config.TestsDirectory == "" {
		config.TestsDirectory = testsDirectory
	}

	if config.ResultsDirectory == "" {
		config.ResultsDirectory = testsDirectory
	}

	if config.Port == "" {
		config.Port = port
	}

	if config.StudentNameLabel == "" {
		config.StudentNameLabel = studentNameLabel
	}

	if config.OpenAnswerLabel == "" {
		config.OpenAnswerLabel = openAnswerLabel
	}

	if config.SubmitButtonLabel == "" {
		config.SubmitButtonLabel = submitButtonLabel
	}

	if config.ErrorHeaderLabel == "" {
		config.ErrorHeaderLabel = errorHeaderLabel
	}

	if config.ErrorDetailsLabel == "" {
		config.ErrorDetailsLabel = errorDetailsLabel
	}

	return config
}
