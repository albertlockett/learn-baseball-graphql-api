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
enum TeamCode {
	ARI
	ATL
	BAL
	BOS
	CHC
	CIN
	CLE
	COL
	CWS
	DET
	HOU
	KC
	LAA
	LAD
	MIA
	MIL
	MIN
	NYM
	NYY
	OAK
	PHI
	PIT
	SD
	SEA
	SF
	STL
	TB
	TEX
	TOR
	WSH
}


type Player {
	name: String
	team: String
	position: String
	fantasyRank: Int
	bats: String
	throws: String
	debut: String
	born: String
	birthCity: String
	birthState: String
	birthCountry: String
	playerID: String
}

type Team {
	city: String
	division: String
	league: String
	name: String
	code: String
}

type Query {
	players(
		teams: [TeamCode]
		maxFantasyRank: Int
	): [Player]
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
