package filtration

type Filtration struct {
	Filter Filter `json:"filter" validate:"required"`
	Sorter Sorter `json:"sorter" validate:"required"`
}
