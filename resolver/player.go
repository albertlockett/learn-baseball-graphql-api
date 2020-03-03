package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/albertlockett/learn-baseball-graphql-api/dao"
	"github.com/olivere/elastic/v7"
)

// PlayerResolverArgs argumentfor players resolver
type PlayerResolverArgs struct {
	MaxFantasyRank *int32
	Teams          *[]*string
}

// PlayerResolver resolve fields on the player
type PlayerResolver struct {
	PlayerName         string `json:"name"`
	PlayerTeam         string `json:"team"`
	PlayerPosition     string `json:"position"`
	PlayerFantasyRank  string `json:"fantasyRank"`
	PlayerBats         string `json:"bats"`
	PlayerThrows       string `json:"throws"`
	PlayerDebut        string `json:"debut"`
	PlayerBorn         string `json:"born"`
	PlayerBirthCity    string `json:"birthCity"`
	PlayerBirthState   string `json:"birthState"`
	PlayerBirthCountry string `json:"birthCountry"`
	PlayerPlayerID     string `json:"playerId"`
}

// AllPlayers returns all the players in the league
func AllPlayers(ctx context.Context, args PlayerResolverArgs) *[]*PlayerResolver {
	ctx = context.Background()
	client := dao.GetESClient()

	query := elastic.NewBoolQuery()

	if args.Teams != nil {
		teamNames := make([]interface{}, len(*args.Teams))
		for i := 0; i < len(*args.Teams); i++ {
			team := *(*args.Teams)[i]
			teamNames[i] = team
		}

		teamTerms := elastic.NewTermsQuery("team", teamNames[:]...)
		query.Filter(teamTerms)
	}

	if args.MaxFantasyRank != nil {
		fantasyRankRangeQuery := elastic.NewRangeQuery("fantasyRank").Lte(*args.MaxFantasyRank)
		query.Filter(fantasyRankRangeQuery)
	}

	fmt.Println(query)

	result, err := client.
		Search().
		Index("players").
		From(0).
		Size(26 * 30).
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

// FantasyRank returns the fantasyRank of the player
func (r PlayerResolver) FantasyRank(ctx context.Context) *int32 {
	if r.PlayerFantasyRank != "" {
		i64, err := strconv.ParseInt(r.PlayerFantasyRank, 10, 32)
		if err != nil {
			panic(err)
		}
		i := int32(i64)
		return &i
	}

	i := int32(-1)
	return &i
}

// Bats returns the bats of the player
func (r PlayerResolver) Bats(ctx context.Context) *string {
	return &r.PlayerBats
}

// Throws returns the throws of the player
func (r PlayerResolver) Throws(ctx context.Context) *string {
	return &r.PlayerThrows
}

// Debut returns the debut of the player
func (r PlayerResolver) Debut(ctx context.Context) *string {
	return &r.PlayerDebut
}

// Born returns the born of the player
func (r PlayerResolver) Born(ctx context.Context) *string {
	return &r.PlayerBorn
}

// BirthCity returns the birthCity of the player
func (r PlayerResolver) BirthCity(ctx context.Context) *string {
	return &r.PlayerBirthCity
}

// BirthState returns the birthState of the player
func (r PlayerResolver) BirthState(ctx context.Context) *string {
	return &r.PlayerBirthState
}

// BirthCountry returns the birthCountry of the player
func (r PlayerResolver) BirthCountry(ctx context.Context) *string {
	return &r.PlayerBirthCountry
}

// PlayerID returns the playerID of the player
func (r PlayerResolver) PlayerID(ctx context.Context) *string {
	return &r.PlayerPlayerID
}
