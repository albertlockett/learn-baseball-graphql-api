package resolver

import (
	"context"
	"encoding/json"
	"log"

	"github.com/albertlockett/learn-baseball-graphql-api/dao"
)

// PlayerResolver resolve fields on the player
type PlayerResolver struct {
	PlayerName     string `json:"name"`
	PlayerTeam     string `json:"team"`
	PlayerPosition string `json:"position"`
}

// var players = []*PlayerResolver{
// 	&PlayerResolver{name: "Gerrit Cole", team: "Yankees", position: "P"},
// 	&PlayerResolver{name: "Yasmani Grandal", team: "White Sox", position: "C"},
// 	&PlayerResolver{name: "Madison Bumgarner", team: "Diamondbacks", position: "P"},
// 	&PlayerResolver{name: "Stephen Strasburg", team: "Nationals", position: "P"},
// 	&PlayerResolver{name: "Zack Wheeler", team: "Phillies", position: "P"},
// 	&PlayerResolver{name: "Hyn-Jin Ryu", team: "Blue Jays", position: "P"},
// }

// AllPlayers returns all the players in the league
func AllPlayers(ctx context.Context) *[]*PlayerResolver {
	ctx = context.Background()
	client := dao.GetESClient()

	result, err := client.
		Search().
		Index("players").
		From(0).
		Size(20).
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
