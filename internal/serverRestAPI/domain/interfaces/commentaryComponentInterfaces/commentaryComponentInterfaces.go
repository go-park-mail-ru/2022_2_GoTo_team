package commentaryComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type CommentaryUsecaseInterface interface {
	AddCommentaryBySession(ctx context.Context, commentary *models.Commentary) error
	GetAllCommentariesForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error)
}

type CommentaryRepositoryInterface interface {
	AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error)
	GetAllCommentsForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error)
}