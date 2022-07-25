package query

type (
	Token struct {
		Kind  TokenKind
		Value interface{}
	}

	TokenKind int
)

const (
	UNKNOWN TokenKind = iota
	// value type
	BOOLEAN
	UINT
	INT
	FLOAT
	STRING
	TIME

	//
	FIELD
	CONDITION_OPERATOR
	LOGICAL_OPERATOR
	SEPARATOR
	CLAUSE
	CLAUSE_CLOSE
)
