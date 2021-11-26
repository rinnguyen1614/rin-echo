package cqrs

type Command struct {
	ID uint `json:"-"`
}

type CreateCommand struct {
	ID uint `json:"-"`
}

type CreateManyCommand struct {
	IDs []uint `json:"-"`
}

type UpdateCommand struct {
	ID uint `json:"-"`
}
