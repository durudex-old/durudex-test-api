package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/durudex/durudex-test-api/internal/domain"
	"github.com/durudex/durudex-test-api/internal/transport/graphql/generated"
	"github.com/segmentio/ksuid"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error) {
	return r.service.Auth.SignUp(ctx, input)
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	return r.service.Auth.SignIn(ctx, input)
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context, input domain.SessionCredInput) (bool, error) {
	return r.service.Auth.SignOut(ctx, input)
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input domain.SessionCredInput) (string, error) {
	return r.service.Auth.RefreshToken(ctx, input)
}

// DeleteSession is the resolver for the deleteSession field.
func (r *mutationResolver) DeleteSession(ctx context.Context, input domain.DeleteSessionInput) (bool, error) {
	return r.service.Auth.DeleteSession(ctx, input)
}

// Session is the resolver for the session field.
func (r *queryResolver) Session(ctx context.Context, id ksuid.KSUID) (*domain.Session, error) {
	return r.service.Auth.Session(ctx, id)
}

// Sessions is the resolver for the sessions field.
func (r *queryResolver) Sessions(ctx context.Context, first *int, last *int, before *string, after *string) (*domain.SessionConnection, error) {
	return r.service.Auth.Sessions(ctx, domain.SortOptions{
		First:  first,
		Last:   last,
		Before: before,
		After:  after,
	})
}

// Edges is the resolver for the edges field.
func (r *sessionConnectionResolver) Edges(ctx context.Context, obj *domain.SessionConnection) ([]*domain.SessionEdge, error) {
	edges := make([]*domain.SessionEdge, len(obj.Nodes))

	for i, node := range obj.Nodes {
		edges[i] = &domain.SessionEdge{
			Cursor: base64.StdEncoding.EncodeToString(node.Id.Bytes()),
			Node:   node,
		}
	}

	return edges, nil
}

// PageInfo is the resolver for the pageInfo field.
func (r *sessionConnectionResolver) PageInfo(ctx context.Context, obj *domain.SessionConnection) (*domain.PageInfo, error) {
	n := len(obj.Nodes)

	if n == 0 {
		return &domain.PageInfo{}, nil
	}

	start := base64.StdEncoding.EncodeToString(obj.Nodes[0].Id.Bytes())
	end := base64.StdEncoding.EncodeToString(obj.Nodes[n-1].Id.Bytes())

	return &domain.PageInfo{StartCursor: &start, EndCursor: &end}, nil
}

// SessionConnection returns generated.SessionConnectionResolver implementation.
func (r *Resolver) SessionConnection() generated.SessionConnectionResolver {
	return &sessionConnectionResolver{r}
}

type sessionConnectionResolver struct{ *Resolver }
