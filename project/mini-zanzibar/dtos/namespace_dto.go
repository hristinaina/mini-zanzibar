package dtos

type Namespace struct {
	Namespace string              `json:"namespace"`
	Relations map[string][]string `json:"relations"`
}

type Namespaces struct {
	Namespaces []Namespace `json:"namespaces"`
}
