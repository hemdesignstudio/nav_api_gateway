package types

import (
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
)

func graphQlObject(name string, fields graphql.Fields) *graphql.Object {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})
	return obj
}

func GenerateGraphQlType(name string, object interface{}, extraFields graphql.Fields) *graphql.Object {
	fields := graphql.Fields{}
	t := reflect.TypeOf(object)

	for i := 0; i < t.NumField(); i++ {
		elemName := t.Field(i).Tag.Get("json")
		elemType := t.Field(i).Type.String()
		if elemType == "string" {
			fields[elemName] = &graphql.Field{Type: graphql.String}
		} else if strings.Contains(elemType, "float") {
			fields[elemName] = &graphql.Field{Type: graphql.Float}
		} else if elemType == "int" {
			fields[elemName] = &graphql.Field{Type: graphql.Int}
		} else if elemType == "bool" {
			fields[elemName] = &graphql.Field{Type: graphql.Boolean}
		}
	}
	if extraFields != nil {
		for key, value := range extraFields {
			fields[key] = value
		}
	}
	return graphQlObject(name, fields)
}