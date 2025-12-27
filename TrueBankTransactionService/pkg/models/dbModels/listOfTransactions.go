package dbModels

type ListTransaction struct {
	Id                               string `json:"id" gorm:"primaryKey"`
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	Company                          string `json:"company"`
	Documents                        string `json:"documents"`
	LinkToIndividualEntrepreneurship string `json:"linkToIndividualEntrepreneurship"`
}
