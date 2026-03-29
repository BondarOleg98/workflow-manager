package filtration

type Sorter struct {
	Order Order  `json:"order" validate:"oneof=ASC DESC"`
	Field string `json:"field"`
}
