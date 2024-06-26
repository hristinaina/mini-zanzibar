package dtos

type Relation struct {
	Object   string `json:"object"`
	Relation string `json:"relation"`
	User     string `json:"user"`
}
