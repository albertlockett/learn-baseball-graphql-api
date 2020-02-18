package resolver

import "context"

// QueryResolver the root query resolver
type QueryResolver struct{}

// NewRoot create a new root of the thing
func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}

// Teams resolve all the teams in the league
func (r QueryResolver) Teams(ctx context.Context) *[]*TeamResolver {
	return AllTeams(ctx)
}

// Players returns all the players
func (r QueryResolver) Players(ctx context.Context, args PlayerResolverArgs) *[]*PlayerResolver {
	return AllPlayers(ctx, args)
}
