package auth

import (
	"context"
	"strconv"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
	"google.golang.org/grpc/metadata"
)

func (c *cli) GetByID(ctx context.Context, id string, userID int64) (*model.Movie, error) {
	header := metadata.New(map[string]string{"user_id": strconv.Itoa(int(userID))})
	outgoingContext := metadata.NewOutgoingContext(context.Background(), header)

	response, err := c.movieClient.GetByID(outgoingContext, converter.ToGetMovieByIDFromService(id))
	if err != nil {
		return nil, err
	}
	return converter.ToMovieFromMovieDesc(response.GetMovie()), nil
}
