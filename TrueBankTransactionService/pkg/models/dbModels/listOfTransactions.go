package dbModels

type ListTransaction struct {
	ID                               uint   `gorm:"primaryKey;autoIncrement"`
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	Company                          string `json:"company"`
	Documents                        string `json:"documents"`
	LinkToIndividualEntrepreneurship string `json:"linkToIndividualEntrepreneurship"`
}
