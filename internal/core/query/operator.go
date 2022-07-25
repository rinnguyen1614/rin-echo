package query

import (
	gormx "github.com/rinnguyen1614/rin-echo/internal/core/gorm"
)

type Operator string

const (
	EqualOperator              = Operator("=")
	GreaterThanOperator        = Operator(">")
	LessThanOperator           = Operator("<")
	GreaterThanOrEqualOperator = Operator(">=")
	LessThanOrEqualOperator    = Operator("<=")
	NotEqualOperator           = Operator("!=")
	LikeOperator               = Operator("like")
	InOperator                 = Operator("in")
	NotOperator                = Operator("not")
	OrOperator                 = Operator("or")
	AndOperator                = Operator("and")
)

var (
	ConditionOperators = []Operator{
		EqualOperator,
		GreaterThanOperator,
		LessThanOperator,
		GreaterThanOrEqualOperator,
		LessThanOrEqualOperator,
		NotEqualOperator,
		LikeOperator,
		InOperator,
	}

	LogicalOperators = []Operator{
		OrOperator,
		AndOperator,
		NotOperator,
	}

	MapConditionOperators = make(map[string]Operator)
	MapLogicalOperators   = make(map[string]Operator)
)

func init() {
	for _, v := range ConditionOperators {
		MapConditionOperators[string(v)] = v
	}

	for _, v := range LogicalOperators {
		MapLogicalOperators[string(v)] = v
	}
}

func (op Operator) DbOperator() (gormx.Operator, error) {
	switch op {
	case EqualOperator:
		return gormx.EqualOperator, nil
	case GreaterThanOperator:
		return gormx.GreaterThanOperator, nil
	case LessThanOperator:
		return gormx.LessThanOperator, nil
	case GreaterThanOrEqualOperator:
		return gormx.GreaterThanOrEqualOperator, nil
	case LessThanOrEqualOperator:
		return gormx.LessThanOrEqualOperator, nil
	case NotEqualOperator:
		return gormx.NotEqualOperator, nil
	case LikeOperator:
		return gormx.LikeOperator, nil
	case InOperator:
		return gormx.InOperator, nil
	case NotOperator:
		return gormx.NotOperator, nil
	case AndOperator:
		return gormx.AndOperator, nil
	case OrOperator:
		return gormx.OrOperator, nil
	default:
		return gormx.Operator(""), ERR_FILTER_CAST_OPERATOR
	}
}
