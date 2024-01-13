package results

import "time"

type TestResults struct {
	Student     string    `yaml:"student"`
	SubmittedAt time.Time `yaml:"submittedAt"`
	Results     Results   `yaml:"results"`
	Test        TestInfo  `yaml:"test"`
}

type TaskResult struct {
	Answer  string `yaml:"answer"`
	Correct bool   `yaml:"correct"`
}

type Results struct {
	Points     int                   `yaml:"points"`
	Total      int                   `yaml:"total"`
	Percentage int                   `yaml:"percentage"`
	Tasks      map[string]TaskResult `yaml:"tasks"`
}

type TestInfo struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Sha256 string `yaml:"sha256"`
}
