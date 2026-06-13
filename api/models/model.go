package models

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Profile holds personal introduction and contact links
type Profile struct {
	Name      string            `json:"name"`
	Title     string            `json:"title"`
	AvatarURL string            `json:"avatarUrl"`
	Bio       string            `json:"bio"`
	Location  string            `json:"location"`
	Contacts  map[string]string `json:"contacts"`
}

// SkillCategory represents a group of skills
type SkillCategory struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

// Experience represents a single job role
type Experience struct {
	Company     string   `json:"company"`
	Role        string   `json:"role"`
	Location    string   `json:"location"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Current     bool     `json:"current"`
	Description []string `json:"description"`
	Tags        []string `json:"tags"`
}

// Project represents a portfolio item
type Project struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	GithubURL   string   `json:"githubUrl"`
	LiveURL     string   `json:"liveUrl"`
	Featured    bool     `json:"featured"`
}

// Education represents a degree or academic qualification
type Education struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

// CV represents the entire structured data of the resume
type CV struct {
	Profile    Profile         `json:"profile"`
	Skills     []SkillCategory `json:"skills"`
	Experience []Experience    `json:"experience"`
	Projects   []Project       `json:"projects"`
	Education  []Education     `json:"education"`
}

//go:embed data/cv.json
var cvJSON []byte

// CVData is the parsed CV struct accessible by handlers and templates
var CVData CV

func init() {
	if err := json.Unmarshal(cvJSON, &CVData); err != nil {
		panic(fmt.Sprintf("failed to parse embedded cv.json: %v", err))
	}
}
