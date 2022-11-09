package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
	"database/sql"
)

type articlePostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewArticlePostgreSQLRepository(database *sql.DB, logger *logger.Logger) articleComponentInterfaces.ArticleRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewArticlePostgreSQLRepository function.")

	articleRepository := &articlePostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("articlePostgreSQLRepository has created.")

	return articleRepository
}

func (apsr *articlePostgreSQLRepository) GetArticleById(ctx context.Context, id int) (*models.Article, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticleById function.")

	row := apsr.database.QueryRow(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
       A.content,
       COALESCE(A.cover_img_path, ''),
       COALESCE(COALESCE(UC.username, ''), ''),
       COALESCE(UC.login, ''),
       COALESCE(UP.username, ''),
       UP.login,
       COALESCE(C.category_name, '')
FROM articles A
         LEFT JOIN users UC ON A.co_author_id = UC.user_id
         JOIN users UP ON A.publisher_id = UP.user_id
         LEFT JOIN categories C ON A.category_id = C.category_id
WHERE A.article_id = $1;
`, id)

	article := &models.Article{}
	if err := row.Scan(
		&article.ArticleId,
		&article.Title,
		&article.Description,
		&article.Rating,
		&article.CommentsCount,
		&article.Content,
		&article.CoverImgPath,
		&article.CoAuthor.Username,
		&article.CoAuthor.Login,
		&article.Publisher.Username,
		&article.Publisher.Login,
		&article.CategoryName); err != nil {
		if err == sql.ErrNoRows {
			apsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.ArticleRepositoryArticleDontExistsError
		}
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.ArticleRepositoryError
	}

	tags, err := apsr.GetTagsForArticle(ctx, article.ArticleId)
	if err != nil {
		return nil, err
	}
	article.Tags = tags

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got article: %#v", article)

	return article, nil
}

func (apsr *articlePostgreSQLRepository) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetTagsForArticle function.")

	tags := make([]string, 0, 10)
	rows, err := apsr.database.Query(`
SELECT T.tag_name
FROM tags T
         JOIN tags_articles TA ON T.tag_id = TA.tag_id
         JOIN articles A ON TA.article_id = A.article_id
WHERE A.article_id = $1;
`, articleId)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.ArticleRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		tag := ""
		if err := rows.Scan(&tag); err != nil {
			apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.ArticleRepositoryError
		}

		tags = append(tags, tag)
	}

	return tags, nil
}