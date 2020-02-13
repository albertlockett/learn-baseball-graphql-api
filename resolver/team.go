package resolver

import "context"

// TeamResolver resolver for a baseball team
type TeamResolver struct {
	city     string
	division string
	league   string
	name     string
}

// AllTeams return list of all teamss
func AllTeams(ctx context.Context) *[]*TeamResolver {
	var teams = []*TeamResolver{
		// AL East
		&TeamResolver{league: "AL", division: "east", city: "Baltimore", name: "Orioles"},
		&TeamResolver{league: "AL", division: "east", city: "Boston", name: "Red Sox"},
		&TeamResolver{league: "AL", division: "east", city: "New York", name: "Yankees"},
		&TeamResolver{league: "AL", division: "east", city: "Tampa Bay", name: "Rays"},
		&TeamResolver{league: "AL", division: "east", city: "Toronto", name: "Blue Jays"},

		// AL Central
		&TeamResolver{league: "AL", division: "central", city: "Chicago", name: "White Sox"},
		&TeamResolver{league: "AL", division: "central", city: "Cleveland", name: "Indians"},
		&TeamResolver{league: "AL", division: "central", city: "Detroit", name: "Tigers"},
		&TeamResolver{league: "AL", division: "central", city: "Kansas City", name: "Royals"},
		&TeamResolver{league: "AL", division: "central", city: "Minnesota", name: "Twins"},

		// AL West
		&TeamResolver{league: "AL", division: "west", city: "Houston", name: "Astros"},
		&TeamResolver{league: "AL", division: "west", city: "Los Angeles", name: "Angels"},
		&TeamResolver{league: "AL", division: "west", city: "Oakland", name: "Athletics"},
		&TeamResolver{league: "AL", division: "west", city: "Seattle", name: "Mariners"},
		&TeamResolver{league: "AL", division: "west", city: "Texas", name: "Rangers"},

		// NL East
		&TeamResolver{league: "NL", division: "east", city: "Atlanta", name: "Braves"},
		&TeamResolver{league: "NL", division: "east", city: "Miami", name: "Marlins"},
		&TeamResolver{league: "NL", division: "east", city: "New York", name: "Mets"},
		&TeamResolver{league: "NL", division: "east", city: "Philadelphia", name: "Phillies"},
		&TeamResolver{league: "NL", division: "east", city: "Washington", name: "Nationals"},

		// NL Central
		&TeamResolver{league: "NL", division: "central", city: "Chicago", name: "Cubs"},
		&TeamResolver{league: "NL", division: "central", city: "Cincinati", name: "Reds"},
		&TeamResolver{league: "NL", division: "central", city: "Milwaukee", name: "Brewers"},
		&TeamResolver{league: "NL", division: "central", city: "Pittsburgh", name: "Pirates"},
		&TeamResolver{league: "NL", division: "central", city: "St. Lous", name: "Cardinals"},

		// NL West
		&TeamResolver{league: "NL", division: "west", city: "Arizona", name: "Diamondbacks"},
		&TeamResolver{league: "NL", division: "west", city: "Colorado", name: "Rockies"},
		&TeamResolver{league: "NL", division: "west", city: "Los Angeles", name: "Dodgers"},
		&TeamResolver{league: "NL", division: "west", city: "San Diego", name: "Padres"},
		&TeamResolver{league: "NL", division: "west", city: "San Francisco", name: "Giants"},
	}
	return &teams
}

// City return the city for a team
func (r TeamResolver) City(ctx context.Context) *string {
	return &r.city
}

// Division return the division for a team
func (r TeamResolver) Division(ctx context.Context) *string {
	return &r.division
}

// League return the division for a team
func (r TeamResolver) League(ctx context.Context) *string {
	return &r.league
}

// Name return the name of a team
func (r TeamResolver) Name(ctx context.Context) *string {
	return &r.name
}
