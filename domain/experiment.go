package domain

type ExperimentDto struct {
	TemplateID string `json:"templateID" binding:"required"`
	Name       string `json:"name" binding:"required"`
}
