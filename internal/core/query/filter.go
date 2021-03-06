package query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	gormx "github.com/rinnguyen1614/rin-echo/internal/core/gorm"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Filter struct {
	tokens []Token
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

	return Filter{tokens: tokens}, nil
}

func (f Filter) Tokens() []Token {
	return f.tokens
}

func (f Filter) Fields() []string {
	var fields []string
	for _, token := range f.tokens {
		if token.Kind == FIELD {
			fields = append(fields, token.Value.(string))
		}
	}
	return fields
}

func (f Filter) Validate(mapFields map[string]reflect.StructField) (err error) {
	notFounds, err := FindFieldNotExists(f.Fields(), mapFields)
	if err != nil {
		return err
	}
	if len(notFounds) != 0 {
		return fmt.Errorf("failed to found filter's fields: %s", strings.Join(notFounds, ", "))
	}
	return nil
}

func (f Filter) BuildQuery(db *gorm.DB, primarySchema *schema.Schema, preloadSchemas map[string]*schema.Schema, primaryFields map[string]reflect.StructField) (*gorm.DB, bool, map[string]string, error) {

	var (
		tx           = db.WithContext(db.Statement.Context)
		primaryTable = primarySchema.Table
		argsCount    = make(map[string]int, 0)
		nToken       = len(f.tokens)
		// key is a full association, value is a dotted string table.
		// Ex: key = UserRoles.Role & value= user_roles.roles
		tableJoinedByFullAssociation = make(map[string]string)
		// key is a full association, value is a dotted string table.
		// Ex: key = user_roles.role & value = UserRoles.Role
		fullAssociationsByRequestName = make(map[string]string)
		// key is a name of table, value is a table in DB.
		fieldNamesByTableDB = make(map[string]string)
		fieldNameByNames    = make(map[string]string)
		valueArgs           = make(map[string]interface{})
		sqlBuilder          strings.Builder
		iter                int
	)

	if nToken == 0 {
		return tx, false, nil, nil
	}

	for iter < nToken {
		var (
			step  = 0
			token = f.tokens[iter]
			dbOp  gormx.Operator
			exp   Expression
			err   error
		)

		if token.Kind == CLAUSE || token.Kind == CLAUSE_CLOSE {
			sqlBuilder.WriteString(token.Value.(string))
		} else {
			if token.Kind == LOGICAL_OPERATOR {
				dbOp, err = Operator(token.Value.(string)).DbOperator()
				if err != nil {
					return nil, false, nil, err
				}

				if Operator(token.Value.(string)) == NotOperator {
					sqlBuilder.WriteString(string(dbOp) + " ")
				} else {
					sqlBuilder.WriteString(" " + string(dbOp) + " ")
				}
			} else {
				exp, step, err = parseExpression(f.tokens[iter:])
				if err != nil {
					return nil, false, nil, err
				}

				dbOp, err = exp.Operator.DbOperator()
				if err != nil {
					return nil, false, nil, err
				}

				var (
					idot       = strings.LastIndexByte(exp.FieldName, byte('.'))
					columnName string
					namedArgs  string
				)

				if idot == -1 {
					columnName = primaryTable + "." + exp.FieldName
				} else {
					var (
						prevFields       = primaryFields
						requestTableName = exp.FieldName[:idot]
						spliteds         = strings.Split(requestTableName, ".")
						association      = fullAssociationsByRequestName[requestTableName]
						tableJoined      = tableJoinedByFullAssociation[association]
						tableName        = tableJoined
					)

					if len(association) == 0 {
						//var lastFieldName string
						for i, splited := range spliteds {
							field, ok := prevFields[splited]
							if !ok {
								return nil, false, nil, fmt.Errorf("failed to found '%s' field", splited)
							}
							schema, ok := preloadSchemas[field.Name]
							if !ok {
								return nil, false, nil, fmt.Errorf("failed to found '%s' schema", field.Name)
							}
							currentModel := reflect.New(field.Type).Interface()
							prevFields, _, err = utils.GetFieldsByJsonTag(currentModel)
							if err != nil {
								return nil, false, nil, err
							}

							tableJoined += schema.Table
							association += field.Name
							fieldNamesByTableDB[tableJoined] = field.Name
							fieldNameByNames[field.Name] = field.Name
							if i != len(spliteds)-1 {
								association += "."
								tableJoined += "."
							} else {
								tableName = schema.Table
							}
						}
						fullAssociationsByRequestName[requestTableName] = association
						tableJoinedByFullAssociation[association] = tableJoined
					} else if idot := strings.LastIndexByte(tableJoined, '.'); idot != -1 {
						tableName = tableJoined[idot:]
					}

					columnName = tableName + exp.FieldName[idot:]
				}
				namedArgs = fmt.Sprintf("%s_%v", exp.FieldName, argsCount[exp.FieldName])
				argsCount[exp.FieldName]++
				sqlBuilder.WriteString(fmt.Sprintf("%v %v @%s", columnName, dbOp, namedArgs))
				valueArgs[namedArgs] = exp.Value
			}
		}

		iter += step + 1
	}

	clauseFrom, err := getClauseFrom(primarySchema, primaryTable, fieldNamesByTableDB)
	if err != nil {
		return nil, false, nil, err
	}
	sql := sqlBuilder.String()
	return tx.Clauses(clauseFrom).Where(sql, valueArgs), true, fieldNameByNames, nil
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
