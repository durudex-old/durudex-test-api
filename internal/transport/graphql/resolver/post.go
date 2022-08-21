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

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input domain.CreatePostInput) (ksuid.KSUID, error) {
	return r.service.Post.CreatePost(ctx, input)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id ksuid.KSUID) (bool, error) {
	return r.service.Post.DeletePost(ctx, id)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	return r.service.Post.UpdatePost(ctx, input)
}

// Edges is the resolver for the edges field.
func (r *postConnectionResolver) Edges(ctx context.Context, obj *domain.PostConnection) ([]*domain.PostEdge, error) {
	edges := make([]*domain.PostEdge, len(obj.Nodes))

	for i, node := range obj.Nodes {
		edges[i] = &domain.PostEdge{
			Cursor: base64.StdEncoding.EncodeToString(node.Id.Bytes()),
			Node:   node,
		}
	}

	return edges, nil
}

// PageInfo is the resolver for the pageInfo field.
func (r *postConnectionResolver) PageInfo(ctx context.Context, obj *domain.PostConnection) (*domain.PageInfo, error) {
	n := len(obj.Nodes)

	if n == 0 {
		return &domain.PageInfo{}, nil
	}

	start := base64.StdEncoding.EncodeToString(obj.Nodes[0].Id.Bytes())
	end := base64.StdEncoding.EncodeToString(obj.Nodes[n-1].Id.Bytes())

	return &domain.PageInfo{StartCursor: &start, EndCursor: &end}, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id ksuid.KSUID) (*domain.Post, error) {
	return r.service.Post.Post(ctx, id)
}

// PostConnection returns generated.PostConnectionResolver implementation.
func (r *Resolver) PostConnection() generated.PostConnectionResolver {
	return &postConnectionResolver{r}
}

type postConnectionResolver struct{ *Resolver }
