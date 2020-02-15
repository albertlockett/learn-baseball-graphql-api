package resolver

import (
	"context"
	"encoding/json"
	"log"

	"github.com/albertlockett/learn-baseball-graphql-api/dao"
)

// TeamResolver resolver for a baseball team
type TeamResolver struct {
	TeamCity     string `json:"city"`
	TeamDivision string `json:"division"`
	TeamLeague   string `json:"league"`
	TeamName     string `json:"name"`
	TeamCode     string `json:"code"`
}

// AllTeams return list of all teamss
func AllTeams(ctx context.Context) *[]*TeamResolver {
	ctx = context.Background()
	client := dao.GetESClient()

	result, err := client.
		Search().
		Index("teams").
		From(0).
		Size(50). // TODO increase limit when the league has > 50 teams
		Do(ctx)
	if err != nil {
		log.Println("an error happened fetching teams")
		panic(err)
	}

	hits := result.Hits.Hits
	numHits := len(hits)
	var teams = make([]*TeamResolver, numHits)

	for i := 0; i < len(hits); i++ {
		var teamResolver TeamResolver
		hit := hits[i]
		err := json.Unmarshal(hit.Source, &teamResolver)
		if err != nil {
			panic(err)
		}
		teams[i] = &teamResolver
	}
	return &teams
}

// City return the city for a team
func (r TeamResolver) City(ctx context.Context) *string {
	return &r.TeamCity
}

// Division return the division for a team
func (r TeamResolver) Division(ctx context.Context) *string {
	return &r.TeamDivision
}

// League return the division for a team
func (r TeamResolver) League(ctx context.Context) *string {
	return &r.TeamLeague
}

// Name return the name of a team
func (r TeamResolver) Name(ctx context.Context) *string {
	return &r.TeamName
}

// Code return the name of a team
func (r TeamResolver) Code(ctx context.Context) *string {
	return &r.TeamCode
}
