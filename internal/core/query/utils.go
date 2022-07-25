package query

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Get clause.From
// primarySchema is the schema of table field in the clause.From. It uses to get a relationship by FieldName
//
// table is a primary table field in clause.From
//
// fieldsByTableJoin with the key is a set tableDB with '.'. The dot '.' represent the relationship between two tables.
// and the value is a FieldName that uses to get relationship of schema.
func getClauseFrom(primarySchema *schema.Schema, table string, fieldsByTableJoin map[string]string) (clause.From, error) {
	var (
		tableSchemas = make(map[string]*schema.Schema)
		clauseFrom   = clause.From{
			Tables: []clause.Table{{Name: table}},
		}
		tableJoinsSorted = make([]string, 0)
	)

	for k := range fieldsByTableJoin {
		tableJoinsSorted = append(tableJoinsSorted, k)
	}

	sort.Strings(tableJoinsSorted)

	for _, tableJoin := range tableJoinsSorted {
		var (
			fieldRelation = fieldsByTableJoin[tableJoin]
			tbSplit       = strings.Split(tableJoin, ".")
		)

		if len(tbSplit) > 1 {
			var (
				prevSchema = primarySchema
				prevTable  string
			)

			for _, tb := range tbSplit {
				if _, ok := tableSchemas[tb]; !ok {
					if rel, ok := prevSchema.Relationships.Relations[fieldRelation]; ok {
						clauseFrom.Joins = append(clauseFrom.Joins, getClauseJoins(rel, prevTable, tb)...)
					} else {
						return clause.From{}, fmt.Errorf("clause: failed to found schema for '%s'", tb)
					}
				}

				prevSchema = tableSchemas[tb]
				prevTable = tb
			}
		} else if rel, ok := primarySchema.Relationships.Relations[fieldRelation]; ok {
			tableSchemas[tableJoin] = rel.FieldSchema
			clauseFrom.Joins = append(clauseFrom.Joins, getClauseJoins(rel, table, tableJoin)...)

		} else {
			return clause.From{}, fmt.Errorf("clause: failed to found schema for '%s'", tableJoin)
		}
	}

	return clauseFrom, nil
}

func getClauseJoins(relation *schema.Relationship, table, tableJoin string) []clause.Join {
	var clauseJoins []clause.Join
	if relation.JoinTable == nil {
		clauseJoins = append(clauseJoins, clause.Join{
			Type:  clause.LeftJoin,
			Table: clause.Table{Name: tableJoin},
			ON: clause.Where{
				Exprs: getEqExps(relation, table, tableJoin),
			},
		})
	} else {
		clauseJoins = append(clauseJoins,
			clause.Join{
				Type:  clause.LeftJoin,
				Table: clause.Table{Name: relation.JoinTable.Table},
				ON: clause.Where{
					Exprs: getEqExps(relation, table, relation.JoinTable.Table),
				},
			},
			clause.Join{
				Type:  clause.LeftJoin,
				Table: clause.Table{Name: tableJoin},
				ON: clause.Where{
					Exprs: getEqExps(relation, tableJoin, relation.JoinTable.Table),
				},
			})
	}
	return clauseJoins
}

func getEqExps(relation *schema.Relationship, table, tableJoin string) []clause.Expression {
	var (
		exprs = make([]clause.Expression, 0)
	)

	if relation.JoinTable == nil {
		for _, ref := range relation.References {
			if ref.OwnPrimaryKey {
				exprs = append(exprs, clause.Eq{
					Column: clause.Column{Table: tableJoin, Name: ref.ForeignKey.DBName},
					Value:  clause.Column{Table: table, Name: ref.PrimaryKey.DBName},
				})
			} else {
				exprs = append(exprs, clause.Eq{
					Column: clause.Column{Table: tableJoin, Name: ref.PrimaryKey.DBName},
					Value:  clause.Column{Table: table, Name: ref.ForeignKey.DBName},
				})
			}
		}
	} else {
		for _, ref := range relation.References {
			if ref.OwnPrimaryKey {
				exprs = append(exprs, clause.Eq{
					Column: clause.Column{Table: tableJoin, Name: ref.ForeignKey.DBName},
					Value:  clause.Column{Table: table, Name: ref.PrimaryKey.DBName},
				})
			}
		}
	}

	return exprs
}

func getJoins(db *gorm.DB, tableJoins []string, queryBuilder iuow.QueryBuilder, preloadBuilders map[string]iuow.QueryBuilder) (*gorm.DB, error) {
	var (
		tableJoineds = make(map[string]bool)
		tx           = db
	)

	for _, tableJoin := range tableJoins {
		tbSplit := strings.Split(tableJoin, ".")

		if len(tbSplit) > 1 {
			var (
				prevBuilder = queryBuilder
			)

			for _, tb := range tbSplit {
				curBuilder, ok := preloadBuilders[tb]
				if !ok {
					return nil, fmt.Errorf("clause: failed to found the preloadBuilder for '%s'", tableJoin)
				}

				if joined := tableJoineds[tb]; !joined {
					if err := tx.SetupJoinTable(prevBuilder.Model(), tb, curBuilder.Model()); err != nil {
						return nil, err
					}

					tableJoineds[tb] = true
				}

				prevBuilder = curBuilder
			}
		} else if builder, ok := preloadBuilders[tableJoin]; ok {
			if joined := tableJoineds[tableJoin]; !joined {
				if err := tx.SetupJoinTable(queryBuilder.Model(), tableJoin, builder.Model()); err != nil {
					return nil, err
				}

				tableJoineds[tableJoin] = true
			}
		} else {
			return nil, fmt.Errorf("clause: failed to found the preloadBuilder for '%s'", tableJoin)
		}
	}

	return tx, nil
}

//
func FindFieldNotExists(fields []string, mapFields map[string]reflect.StructField) (fieldNotFounds []string, err error) {

	sort.Strings(fields)

	for i := 0; i < len(fields); i++ {
		var (
			field      = fields[i]
			prevFields = mapFields
		)
		for {
			var (
				name = field
				iDot = strings.IndexByte(field, '.')
			)
			if iDot == -1 {
				if _, ok := prevFields[name]; !ok {
					fieldNotFounds = append(fieldNotFounds, fields[i])
				}
				break
			} else {
				name = field[:iDot]
				if structField, ok := prevFields[name]; !ok {
					fieldNotFounds = append(fieldNotFounds, fields[i])
					// next field that has same prefix
					for ; i < len(fields)-1; i++ {
						part := fields[i+1]
						iPartDot := strings.IndexByte(part, '.')
						if iPartDot != -1 && part[:iPartDot] == name {
							fieldNotFounds = append(fieldNotFounds, part)
						} else {
							// because all fields in fields that are sorted, so fields after the index will not be equal
							break
						}
					}
					break
				} else {
					prevFields, _, err = utils.GetFullFieldsByJsonTag(reflect.New(structField.Type).Interface())
					if err != nil {
						return nil, err
					}
					field = field[iDot+1:]
				}
			}
		}
	}

	return fieldNotFounds, nil
}
