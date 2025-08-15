package models

type ResearchTeam struct {
	ID                  uint   `json:"id" gorm:"primaryKey"`
	TeamName            string `json:"teamName"`
	TeamLead            string `json:"teamLead"`
	Members             string `json:"members"` // Storing as comma-separated string for simplicity
	ResearchDescription string `json:"researchDescription"`
	Image               string `json:"image"`
}
