package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-blog/graph/generated"
	"go-blog/graph/model"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	article := model.Article{
		Title:   input.Title,
		Content: input.Content,
		Status:  input.Status,
	}
	_, err := r.DB.Model(&article).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new Article: %v", err)
	}
	return &article, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	author := model.Author{
		Name: input.Name,
	}
	_, err := r.DB.Model(&author).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new author: %v", err)
	}
	return &author, nil
}

func (r *queryResolver) Article(ctx context.Context) ([]*model.Article, error) {
	var articles []*model.Article

	err := r.DB.Model(&articles).Select()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *queryResolver) Author(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author

	err := r.DB.Model(&authors).Select()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
