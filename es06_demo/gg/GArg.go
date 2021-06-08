package gg

import "github.com/graphql-go/graphql"

func IntArg(name string) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{name: &graphql.ArgumentConfig{Type: graphql.Int}}
}

//单参数
func StringArg(name string) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{name: &graphql.ArgumentConfig{Type: graphql.Int}}
}

//多参数
func StringArgs(names ...string) graphql.FieldConfigArgument {
	if len(names) == 0 {
		return nil
	}
	config := make(graphql.FieldConfigArgument)
	for _, name := range names {
		config[name] = &graphql.ArgumentConfig{Type: graphql.String}
	}
	return config
}
