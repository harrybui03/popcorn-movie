package movie

import (
	"PopcornMovie/ent"
	"context"
	"github.com/pkg/errors"
)

type Repository interface {
	MovieQuery() *ent.MovieQuery
	GetAllMovie(ctx context.Context, query *ent.MovieQuery) ([]*ent.Movie, error)
	CountMovies(ctx context.Context, query *ent.MovieQuery) (*int, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) MovieQuery() *ent.MovieQuery {
	return i.client.Movie.Query()
}

func (i impl) GetAllMovie(ctx context.Context, query *ent.MovieQuery) ([]*ent.Movie, error) {
	movies, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return movies, nil
}

func (i impl) CountMovies(ctx context.Context, query *ent.MovieQuery) (*int, error) {
	cnt, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &cnt, nil
}

// New is function init Room Repo
func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
