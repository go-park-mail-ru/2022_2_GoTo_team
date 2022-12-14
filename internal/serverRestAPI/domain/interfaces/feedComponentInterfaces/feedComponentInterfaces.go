package feedComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

//go:generate mockgen -destination=./mock/feedRepositoryMock.go -package=mock 2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces FeedRepositoryInterface

type FeedUsecaseInterface interface {
	GetFeed(ctx context.Context) ([]*models.Article, error)
	GetFeedForUserByLogin(ctx context.Context, login string) ([]*models.Article, error)
	GetFeedForCategory(ctx context.Context, category string) ([]*models.Article, error)
	GetNewArticlesFromIdForSubscriber(ctx context.Context, articleId int) ([]int, error)
}

type FeedRepositoryInterface interface {
	GetFeed(ctx context.Context, email string) ([]*models.Article, error)
	GetFeedForUserByLogin(ctx context.Context, login string, email string) ([]*models.Article, error)
	GetFeedForCategory(ctx context.Context, category string, email string) ([]*models.Article, error)
	UserExistsByLogin(ctx context.Context, login string) (bool, error)
	CategoryExists(ctx context.Context, category string) (bool, error)
	GetNewArticlesFromIdForSubscriber(ctx context.Context, articleId int, email string) ([]int, error)
}
