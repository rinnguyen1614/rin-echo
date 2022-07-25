package gorm

import "fmt"

type Operator string

const (
	EqualOperator              = Operator("=")
	GreaterThanOperator        = Operator(">")
	LessThanOperator           = Operator("<")
	GreaterThanOrEqualOperator = Operator(">=")
	LessThanOrEqualOperator    = Operator("<=")
	NotEqualOperator           = Operator("<>")
	BetweenOperator            = Operator("BETWEEN")
	LikeOperator               = Operator("LIKE")
	InOperator                 = Operator("IN")
	OrOperator                 = Operator("OR")
	AndOperator                = Operator("AND")
	NotOperator                = Operator("NOT")
)

func (op Operator) Condition(field string) string {
	if op.isJoinOperator() || op == NotOperator {
		return ""
	}

	if op == BetweenOperator {
		return fmt.Sprintf("%v %v ? AND ?", field, op)
	}

	return fmt.Sprintf("%v %v ?", field, op)
}

func (op Operator) Join(query1, query2 string) string {
	if !op.isJoinOperator() {
		return ""
	}

	return fmt.Sprintf("%v %v %v", query1, op, query2)
}

func (op Operator) Not(query string) string {
	if op != NotOperator {
		return ""
	}

	return fmt.Sprintf("%v %v", op, query)
}

func (op Operator) isJoinOperator() bool {
	return op == OrOperator || op == AndOperator
}
