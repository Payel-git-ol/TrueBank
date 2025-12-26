package requestModels

type RegTransaction struct {
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	Company                          string `json:"company"`
	Documents                        string `json:"documents"`
	LinkToIndividualEntrepreneurship string `json:"linkToIndividualEntrepreneurship"`
}
