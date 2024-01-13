package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Attachment struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Src  string `json:"src"`
}

type Task struct {
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	Attachment Attachment `json:"attachment"`
	Options    []string   `json:"options"`
	Answer     string     `json:"answer"`
}

type Test struct {
	Title       string    `json:"title"`
	Target      string    `json:"target"`
	Description string    `json:"description"`
	Subject     string    `json:"subject"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresIn   time.Time `json:"expiresIn"`
	Tasks       []Task    `json:"tasks"`
}

func (t Test) IsExpired() bool {
	return !t.ExpiresIn.IsZero() && t.ExpiresIn.Before(time.Now())
}

func (t Test) Sha256Sum() string {
	hasher := sha256.New()
	data, err := json.Marshal(t)

	if err != nil {
		return ""
	}

	hasher.Write(data)

	return hex.EncodeToString(hasher.Sum(nil))
}
