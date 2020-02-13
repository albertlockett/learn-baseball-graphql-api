package resolver

import "context"

// PlayerResolver resolve fields on the player
type PlayerResolver struct {
	name     string
	team     string
	position string
}

var players = []*PlayerResolver{
	&PlayerResolver{name: "Gerrit Cole", team: "Yankees", position: "P"},
	&PlayerResolver{name: "Yasmani Grandal", team: "White Sox", position: "C"},
	&PlayerResolver{name: "Madison Bumgarner", team: "Diamondbacks", position: "P"},
	&PlayerResolver{name: "Stephen Strasburg", team: "Washington", position: "P"},
	&PlayerResolver{name: "Zack Wheeler", team: "Phillies", position: "P"},
	&PlayerResolver{name: "Hyn-Jin Ryu", team: "Blue Jays", position: "P"},
}

// AllPlayers returns all the players in the league
func AllPlayers(ctx context.Context) *[]*PlayerResolver {
	return &players
}

// Name returns the name of the player
func (r PlayerResolver) Name(ctx context.Context) *string {
	return &r.name
}

// Team returns the name of the team
func (r PlayerResolver) Team(ctx context.Context) *string {
	return &r.team
}

// Position returns the position of the player
func (r PlayerResolver) Position(ctx context.Context) *string {
	return &r.position
}
