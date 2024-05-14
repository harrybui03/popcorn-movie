package movie

import (
	"PopcornMovie/ent"
	"PopcornMovie/ent/movie"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	MovieQuery() *ent.MovieQuery
	MovieCreate() *ent.MovieCreate
	MovieUpdate() *ent.MovieUpdate
	MovieDelete() *ent.MovieDelete
	GetAllMovie(ctx context.Context, query *ent.MovieQuery) ([]*ent.Movie, error)
	CountMovies(ctx context.Context, query *ent.MovieQuery) (*int, error)
	GetMovieByID(ctx context.Context, id uuid.UUID) (*ent.Movie, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) MovieUpdate() *ent.MovieUpdate {
	return i.client.Movie.Update()
}

func (i impl) MovieDelete() *ent.MovieDelete {
	return i.client.Movie.Delete()
}

func (i impl) MovieCreate() *ent.MovieCreate {
	return i.client.Movie.Create()
}

func (i impl) GetMovieByID(ctx context.Context, id uuid.UUID) (*ent.Movie, error) {
	movieRecord, err := i.client.Movie.Query().Where(movie.ID(id)).First(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return movieRecord, nil
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
