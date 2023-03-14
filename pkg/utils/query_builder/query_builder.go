package query_builder

import "gorm.io/gorm"

type QueryBuilder interface {
	AddWhereCondition(query interface{}, args ...interface{})
	AddWhereConditions(queries ...whereCondition)
	GetWhereConditions() *whereConditions
	WhereConditionsScope(db *gorm.DB) *gorm.DB
}

type whereCondition struct {
	Query interface{}
	Args  []interface{}
}

type whereConditions = []whereCondition

type queryBuilder struct {
	whereConditions []whereCondition
}

func NewQueryBuilder() QueryBuilder {
	return &queryBuilder{}
}

func NewQueryBuilderWhereCondition() QueryBuilder {
	return &queryBuilder{
		whereConditions: []whereCondition{},
	}
}

func NewWhereCondition(query interface{}, args ...interface{}) *whereCondition {
	return &whereCondition{query, args}
}

func (c *queryBuilder) AddWhereCondition(query interface{}, args ...interface{}) {
	c.whereConditions = append(c.whereConditions, whereCondition{query, args})
}

func (c *queryBuilder) AddWhereConditions(queries ...whereCondition) {
	c.whereConditions = append(c.whereConditions, queries...)
}

func (c *queryBuilder) GetWhereConditions() *whereConditions {
	return &c.whereConditions
}

func (c *queryBuilder) WhereConditionsScope(db *gorm.DB) *gorm.DB {
	for _, cond := range *c.GetWhereConditions() {
		db.Where(cond.Query, cond.Args...)
	}
	return db
}
