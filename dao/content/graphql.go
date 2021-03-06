package content

import (
	"fmt"

	"github.com/FrontMage/HelloGithubNavBackend/dao/category"
	"github.com/FrontMage/HelloGithubNavBackend/dao/volume"
	"github.com/graphql-go/graphql"
)

var graphqlType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Content",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"projectURL":  &graphql.Field{Type: graphql.String},
		"title":       &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"imagePath":   &graphql.Field{Type: graphql.String},
		"categoryID":  &graphql.Field{Type: graphql.Int},
		"volumeID":    &graphql.Field{Type: graphql.Int},
		"createTime":  &graphql.Field{Type: graphql.String},
		"updateTime":  &graphql.Field{Type: graphql.String},
		"status":      &graphql.Field{Type: graphql.Int},
		"category":    &graphql.Field{Type: category.GraphqlType},
		"volume":      &graphql.Field{Type: volume.GraphqlType},
	},
})

var rootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"contents": &graphql.Field{
			Type:        graphql.NewList(graphqlType),
			Description: "Return content list",
			Args: graphql.FieldConfigArgument{
				"ids": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ids, ok := p.Args["ids"].([]interface{})
				if !ok {
					return nil, fmt.Errorf("Failed to parse ids=%+v %T", p.Args["ids"], p.Args["ids"])
				}
				parsedIds := make([]uint64, len(ids))
				for idx, id := range ids {
					parsedID, ok := id.(int)
					if !ok {
						fmt.Printf("Failed to parse id=%+v %T to int\n", id, id)
					}
					parsedIds[idx] = uint64(parsedID)
				}
				return BatchGet(parsedIds)
			},
		},
	},
}

var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

// Schema graphql schema for content query
var Schema, _ = graphql.NewSchema(schemaConfig)
