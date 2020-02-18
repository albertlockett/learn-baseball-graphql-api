package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/albertlockett/learn-baseball-graphql-api/dao"
	"github.com/olivere/elastic/v7"
)

// PlayerResolverArgs argumentfor players resolver
type PlayerResolverArgs struct {
	Teams *[]*string
}

// PlayerResolver resolve fields on the player
type PlayerResolver struct {
	PlayerName     string `json:"name"`
	PlayerTeam     string `json:"team"`
	PlayerPosition string `json:"position"`
}

// AllPlayers returns all the players in the league
func AllPlayers(ctx context.Context, args PlayerResolverArgs) *[]*PlayerResolver {
	ctx = context.Background()
	client := dao.GetESClient()

	for i := 0; i < len(*args.Teams); i++ {
		team := *(*args.Teams)[i]
		fmt.Println(team)
	}

	query := elastic.NewBoolQuery()

	if *args.Teams != nil {
		teamNames := make([]interface{}, len(*args.Teams))
		for i := 0; i < len(*args.Teams); i++ {
			team := *(*args.Teams)[i]
			teamNames[i] = team
		}

		teamTerms := elastic.NewTermsQuery("team", teamNames[:]...)
		query.Filter(teamTerms)
	}

	fmt.Println(query)

	result, err := client.
		Search().
		Index("players").
		From(0).
		Size(20).
		Query(query).
		Do(ctx)
	if err != nil {
		log.Println("an error happened fetching teams")
		panic(err)
	}

	hits := result.Hits.Hits
	numHits := len(hits)
	var players = make([]*PlayerResolver, numHits)

	for i := 0; i < len(hits); i++ {
		var playerResolver PlayerResolver
		hit := hits[i]
		err := json.Unmarshal(hit.Source, &playerResolver)
		if err != nil {
			panic(err)
		}
		players[i] = &playerResolver
	}

	return &players
}

// Name returns the name of the player
func (r PlayerResolver) Name(ctx context.Context) *string {
	return &r.PlayerName
}

// Team returns the name of the team
func (r PlayerResolver) Team(ctx context.Context) *string {
	return &r.PlayerTeam
}

// Position returns the position of the player
func (r PlayerResolver) Position(ctx context.Context) *string {
	return &r.PlayerPosition
}
