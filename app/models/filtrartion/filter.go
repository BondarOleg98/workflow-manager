package filtrartion

type Filter struct {
	Operator Operator `json:"operator"`
	Field    string   `json:"field"`
	Value    string   `json:"value"`
}
