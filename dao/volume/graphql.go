package volume

import "github.com/graphql-go/graphql"

// GraphqlType for volume
var GraphqlType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"name":       &graphql.Field{Type: graphql.String},
		"createTime": &graphql.Field{Type: graphql.String},
		"updateTime": &graphql.Field{Type: graphql.String},
		"status":     &graphql.Field{Type: graphql.Int},
	},
})
