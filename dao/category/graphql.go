package category

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// GraphqlType for category
var GraphqlType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"name":       &graphql.Field{Type: graphql.String},
		"createTime": &graphql.Field{Type: graphql.String},
		"updateTime": &graphql.Field{Type: graphql.String},
	},
})

var rootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"category": &graphql.Field{
			Type:        graphql.NewList(GraphqlType),
			Description: "Find single category by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, fmt.Errorf("Failed to parse id %+v %T", p.Args["id"], p.Args["id"])
				}
				return Get(uint64(id))
			},
		},
	},
}

var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

// Schema graphql schema for content query
var Schema, _ = graphql.NewSchema(schemaConfig)
