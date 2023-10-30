package parser

import "time"

type Test struct {
	Title       string    `json:"title"`
	Target      string    `json:"target"`
	Subject     string    `json:"subject"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresIn   time.Time `json:"expiresIn"`
	Tasks       []struct {
		Type       string `json:"type"`
		Text       string `json:"text"`
		Attachment struct {
			Type string `json:"type"`
			Src  string `json:"src"`
		} `json:"attachment"`
		Options []string `json:"options"`
		Answer  string   `json:"answer"`
	} `json:"tasks"`
}
