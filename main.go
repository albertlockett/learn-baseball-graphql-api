package main

import (
	"log"
	"net/http"

	"github.com/albertlockett/learn-baseball-graphql-api/resolver"
	"github.com/friendsofgo/graphiql"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	mux.Handle("/graphql", &relay.Handler{Schema: schema})
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(nil)
	}
	mux.Handle("/", graphiqlHandler)
	handler := cors.Default().Handler(mux)

	log.Println("Server ready at 80")
	log.Fatal(http.ListenAndServe(":80", handler))
}
