package api

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/loaders"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

type directorResolver struct {
	*Resolver
}

func (r *Resolver) Director() DirectorResolver {
	return &directorResolver{r}
}

func (r *directorResolver) Talent(ctx context.Context, obj *model.Director) (*model.Talent, error) {
	talent, err := ctx.Value(middleware.TalentLoaderKey).(*loaders.TalentLoader).Load(obj.TalentID)
	return talent, err
}

type cinematographerResolver struct {
	*Resolver
}

func (r *Resolver) Cinematographer() CinematographerResolver {
	return &cinematographerResolver{r}
}

func (r *cinematographerResolver) Talent(ctx context.Context, obj *model.Cinematographer) (*model.Talent, error) {
	talent, err := ctx.Value(middleware.TalentLoaderKey).(*loaders.TalentLoader).Load(obj.TalentID)
	return talent, err
}

type composerResolver struct {
	*Resolver
}

func (r *Resolver) Composer() ComposerResolver {
	return &composerResolver{r}
}

func (r *composerResolver) Talent(ctx context.Context, obj *model.Composer) (*model.Talent, error) {
	talent, err := ctx.Value(middleware.TalentLoaderKey).(*loaders.TalentLoader).Load(obj.TalentID)
	return talent, err
}

type producerResolver struct {
	*Resolver
}

func (r *Resolver) Producer() ProducerResolver {
	return &producerResolver{r}
}

func (r *producerResolver) Talent(ctx context.Context, obj *model.Producer) (*model.Talent, error) {
	talent, err := ctx.Value(middleware.TalentLoaderKey).(*loaders.TalentLoader).Load(obj.TalentID)
	return talent, err
}

type roleResolver struct {
	*Resolver
}

func (r *Resolver) Role() RoleResolver {
	return &roleResolver{r}
}

func (r *roleResolver) Talent(ctx context.Context, obj *model.Role) (*model.Talent, error) {
	talent, err := ctx.Value(middleware.TalentLoaderKey).(*loaders.TalentLoader).Load(obj.TalentID)
	return talent, err
}

type screenwriterResolver struct {
	*Resolver
}

func (r *Resolver) Screenwriter() ScreenwriterResolver {
	return &screenwriterResolver{r}
}

func (r *screenwriterResolver) Talent(ctx context.Context, obj *model.Screenwriter) (*model.Talent, error) {
	talent, err := ctx.Value(middleware.TalentLoaderKey).(*loaders.TalentLoader).Load(obj.TalentID)
	return talent, err
}
