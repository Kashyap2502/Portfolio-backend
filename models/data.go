package models

type AboutResp struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Profile     string `json:"profile"`
}

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Techstack   string `json:"techstack"`
	Url         string `json:"url"`
}

type ProjectsResp struct {
	Projects []Project `json:"projects"`
}

type Education struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Year        int    `json:"year"`
}

type EducationResp struct {
	Education []Education `json:"education"`
}

type Experience struct {
	Company     string `json:"company"`
	Role        string `json:"role"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}

type ExperienceResp struct {
	Experience []Experience `json:"experience"`
}
