package test

type InputCreateTest struct {
	Name string `json:"name" validate:"required,lowercase"`
}
