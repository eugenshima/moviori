package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eugenshima/moviori/internal/model"
	proto "github.com/eugenshima/moviori_movies/proto"
)

type MovieRepository struct {
	client proto.MovioriMoviesClient
}

func NewMovieRepository(client proto.MovioriMoviesClient) *MovieRepository {
	return &MovieRepository{client: client}
}

func (r *MovieRepository) FindByName(ctx context.Context, id string) (*model.FinalMovie, error) {
	response, err := r.client.FindByName(ctx, &proto.FindMovieRequest{
		ID: id,
	})
	movie := &model.Movie{}
	err = json.Unmarshal(response.MovieInfo, &movie)
	if err != nil {
		return nil, fmt.Errorf(" Unmarshal: %w", err)
	}
	finalMovie := &model.FinalMovie{
		ID:              movie.ID,
		Name:            movie.Name,
		AlternativeName: movie.AlternativeName,
		Year:            movie.Year,
		MovieLength:     movie.MovieLength,
	}
	return finalMovie, nil
}
