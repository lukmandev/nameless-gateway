package auth

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) GetRecommendations(ctx context.Context, limit int32, userID int64) ([]*model.Movie, error) {
	header := metadata.New(map[string]string{"user_id": strconv.Itoa(int(userID))})
	outgoingContext := metadata.NewOutgoingContext(context.Background(), header)

	response, err := c.movieClient.GetRecommendations(outgoingContext, converter.ToGetRecommendationsInputFromService(limit))
	if err != nil {
		return nil, err
	}
	return converter.ToMovieFromMovieDescList(response.GetMovies()), nil
}
