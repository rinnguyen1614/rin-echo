package query

import (
	"errors"
	"fmt"
	"reflect"
	"rin-echo/common/gorm"
	"rin-echo/common/utils"
	"strings"
)

const (
	// it represents the table without dot '.'
	tableTemp = "__table__"
	// it represents the table with dot '.'
	tableDotTemp = "__tableDot__"
)

type Filter struct {
	tokens     []Token
	SQL        strings.Builder
	Vars       map[string]interface{}
	FieldNames map[string]string
	// include '.' to easy get table reference.
	Tables []string
	// key: table without '.', value: table include '.'
	TableByNames map[string]string
}

func newFilter() Filter {
	return Filter{}
}

func ParseFilter(str string) (Filter, error) {
	if strings.Trim(str, " ") == "" {
		return Filter{}, nil
	}

	tokens, err := ParseToken(str)
	if err != nil {
		return Filter{}, err
	}

	var (
		iter       int
		nToken     int = len(tokens)
		fieldCount     = make(map[string]int, 0)
		fil            = Filter{
			tokens:       tokens,
			SQL:          strings.Builder{},
			Vars:         make(map[string]interface{}),
			FieldNames:   make(map[string]string),
			Tables:       make([]string, 0),
			TableByNames: make(map[string]string),
		}
	)

	for iter < nToken {
		var (
			step  = 0
			token = tokens[iter]
			dbOp  gorm.Operator
			exp   Expression
		)

		if token.Kind == CLAUSE || token.Kind == CLAUSE_CLOSE {
			fil.SQL.WriteString(token.Value.(string))
		} else {
			if token.Kind == LOGICAL_OPERATOR {
				dbOp, err = Operator(token.Value.(string)).DbOperator()
				if err != nil {
					return Filter{}, err
				}

				if Operator(token.Value.(string)) == NotOperator {
					fil.SQL.WriteString(string(dbOp) + " ")
				} else {
					fil.SQL.WriteString(" " + string(dbOp) + " ")
				}
			} else {
				exp, step, err = parseExpression(tokens[iter:])
				if err != nil {
					return Filter{}, err
				}

				dbOp, err = exp.Operator.DbOperator()
				if err != nil {
					return Filter{}, err
				}

				var (
					idot          = strings.LastIndexByte(exp.FieldName, byte('.'))
					fieldName     string
					keyFieldParam string
				)

				if idot == -1 {
					fieldName = tableTemp + "." + exp.FieldName
				} else {
					// Ex: FieldName: A.B.C
					// table:  => A.B is a table of field C.
					tableName := exp.FieldName[:idot]
					// field: A.B.C => slipt => B.C <=> E is a field of table B.
					tableSplit := strings.Split(tableName, ".")
					fieldName = getKeyTableDotTemp(tableSplit[len(tableSplit)-1]) + exp.FieldName[idot:]
					if _, ok := fil.TableByNames[tableName]; !ok {
						fil.Tables = append(fil.Tables, tableName)
						fil.TableByNames[tableName] = tableName
					}
				}
				keyFieldParam = fmt.Sprintf("%s_%v", exp.FieldName, fieldCount[exp.FieldName])
				fieldCount[exp.FieldName]++
				fil.Vars[keyFieldParam] = exp.Value
				fil.FieldNames[exp.FieldName] = fieldName
				fil.SQL.WriteString(fmt.Sprintf("%v %v @%s", fieldName, dbOp, keyFieldParam))
			}
		}

		iter += step + 1
	}

	return fil, nil
}

func (f Filter) Tokens() []Token {
	return f.tokens
}

// Field's name of model by table. Key: table, value: field's name
//
// Ex: table = user_roles.role => fieldName = UserRole.Role
func (f Filter) GetFieldNamesByTable(model interface{}) (map[string]string, error) {
	fieldOfTables := make(map[string]string)
	fields, err := utils.GetFieldsByJsonTag(model)
	if err != nil {
		return nil, err
	}

	for _, table := range f.Tables {
		if ts := strings.Split(table, "."); len(ts) > 1 {
			var (
				prevFields = fields
				fieldName  string
			)

			for i, t := range ts {
				field, ok := prevFields[t]
				if !ok {
					return nil, fmt.Errorf("failed to found '%s' field", t)
				}

				prevFields, err = utils.GetFieldsByJsonTag(reflect.New(field.Type).Interface())
				if err != nil {
					return nil, err
				}

				fieldName += field.Name
				if i != len(ts)-1 {
					fieldName += "."
				}
			}

			fieldOfTables[table] = fieldName
		} else {
			field, ok := fields[table]
			if !ok {
				return nil, fmt.Errorf("failed to found '%s' field", table)
			}

			fieldOfTables[table] = field.Name
		}
	}

	return fieldOfTables, nil
}

type Expression struct {
	FieldName string
	Operator  Operator
	Value     []interface{}
}

func (exp Expression) IsNil() bool {
	return len(exp.FieldName) == 0
}

func parseExpression(tokens []Token) (Expression, int, error) {
	exp := Expression{Value: make([]interface{}, 0)}
	var openBrakets int
	for i, token := range tokens {
		switch token.Kind {
		case CLAUSE:
			openBrakets++
			if !exp.IsNil() {
				continue
			}
			return exp, i - 1, nil
		case CLAUSE_CLOSE:
			openBrakets--
			if openBrakets == 0 {
				return exp, i, nil
			}
			return exp, i - 1, nil
		case LOGICAL_OPERATOR:
			return exp, i - 1, nil
		case FIELD:
			exp.FieldName = token.Value.(string)
		case CONDITION_OPERATOR:
			if exp.IsNil() {
				return Expression{}, i, errors.New(fmt.Sprintf("Missing field name. Syntax error at or near '%v'.", token.Value))
			}
			exp.Operator = Operator(token.Value.(string))
		case SEPARATOR:
			if exp.IsNil() {
				return Expression{}, i, errors.New(fmt.Sprintf("Missing field name. Syntax error at or near '%v'.", token.Value))
			}
		default:
			if isValue(token) {
				if exp.IsNil() {
					return Expression{}, i, errors.New(fmt.Sprintf("Missing field name or operator for expression. Syntax error at or near '%v'.", token.Value))
				}
				exp.Value = append(exp.Value, token.Value)
			} else {
				return Expression{}, i, errors.New(fmt.Sprintf("failed to found token kind '%v'", token.Kind))
			}
		}
	}

	return exp, len(tokens) - 1, nil
}

func isValue(token Token) bool {
	return token.Kind == BOOLEAN ||
		token.Kind == UINT ||
		token.Kind == INT ||
		token.Kind == FLOAT ||
		token.Kind == STRING ||
		token.Kind == TIME
}

func getKeyTableDotTemp(table string) string {
	return tableDotTemp + "." + table
}
