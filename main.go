package main

import (
	"log"
	"net/http"

	"github.com/albertlockett/learn-baseball-graphql-api/resolver"
	"github.com/friendsofgo/graphiql"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

const s = `
type Player {
	name: String
	team: String
	position: String
}

type Team {
	city: String
	division: String
	league: String
	name: String
}

type Query {
	players: [Player]
	teams: [Team]
}
`

type query struct{}

func main() {
	root, err := resolver.NewRoot()
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, root)
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(nil)
	}
	http.Handle("/", graphiqlHandler)

	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
