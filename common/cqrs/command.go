package cqrs

type CreateCommand struct {
	ID uint `json:"-"`
}

type CreateManyCommand struct {
	IDs []uint `json:"-"`
}
