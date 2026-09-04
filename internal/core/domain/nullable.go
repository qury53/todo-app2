package domain

type Nullable[T any] struct {
	Value *T   `json:"value" example:"some_value"`
	Set   bool `json:"set" example:"true"`
}
