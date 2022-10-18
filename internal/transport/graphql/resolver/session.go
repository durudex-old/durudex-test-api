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

// DeleteSession is the resolver for the deleteSession field.
func (r *mutationResolver) DeleteSession(ctx context.Context, id ksuid.KSUID) (bool, error) {
	return r.service.Session.Delete(ctx, id)
}

// Session is the resolver for the session field.
func (r *queryResolver) Session(ctx context.Context, id ksuid.KSUID) (*domain.Session, error) {
	return r.service.Session.Get(ctx, id)
}

// Sessions is the resolver for the sessions field.
func (r *queryResolver) Sessions(ctx context.Context, first *int, last *int, before *string, after *string) (*domain.SessionConnection, error) {
	return r.service.Session.GetList(ctx, domain.SortOptions{
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
