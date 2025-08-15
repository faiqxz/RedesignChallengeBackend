package models

type Certification struct {
	ID                   uint   `json:"id" gorm:"primaryKey"`
	CertificationName    string `json:"certificationName"`
	IssuingOrganization  string `json:"issuingOrganization"`
	Description          string `json:"description"`
	CertificationLink    string `json:"certificationLink"`
}
