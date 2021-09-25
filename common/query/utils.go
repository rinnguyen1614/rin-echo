package query

import (
	"fmt"
	"sort"
	"strings"

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
						tableSchemas[tb] = rel.FieldSchema
						clauseFrom.Joins = append(clauseFrom.Joins, clause.Join{
							Type:  clause.LeftJoin,
							Table: clause.Table{Name: tb},
							ON: clause.Where{
								Exprs: getEqExps(rel, prevTable, tb),
							},
						})
					} else {
						return clause.From{}, fmt.Errorf("failed to found schema for '%s'", tb)
					}
				}

				prevSchema = tableSchemas[tb]
				prevTable = tb
			}
		} else if rel, ok := primarySchema.Relationships.Relations[fieldRelation]; ok {
			tableSchemas[tableJoin] = rel.FieldSchema
			clauseFrom.Joins = append(clauseFrom.Joins, clause.Join{
				Type:  clause.LeftJoin,
				Table: clause.Table{Name: tableJoin},
				ON: clause.Where{
					Exprs: getEqExps(rel, table, tableJoin),
				},
			})
		} else {
			return clause.From{}, fmt.Errorf("failed to found schema for '%s'", tableJoin)
		}
	}

	return clauseFrom, nil
}

func getEqExps(relation *schema.Relationship, table, tableJoin string) []clause.Expression {
	var (
		exprs = make([]clause.Expression, 0)
	)

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

	return exprs
}

func getJoins(db *gorm.DB, tableJoins []string, queryBuilder QueryBuilder, preloadBuilders map[string]QueryBuilder) (*gorm.DB, error) {
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
					return nil, fmt.Errorf("failed to found the preloadBuilder for '%s'", tableJoin)
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
			return nil, fmt.Errorf("failed to found the preloadBuilder for '%s'", tableJoin)
		}
	}

	return tx, nil
}
