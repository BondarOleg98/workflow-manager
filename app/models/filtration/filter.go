package filtration

type Operator string

type Filter struct {
	Operator Operator `json:"operator" validate:"oneof=IN EQ LIKE"`
	Field    string   `json:"field"`
	Value    string   `json:"value"`
}
