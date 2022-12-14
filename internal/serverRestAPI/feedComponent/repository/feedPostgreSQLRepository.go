package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/feedComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
	"fmt"
)

type feedPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewFeedPostgreSQLRepository(database *sql.DB, logger *logger.Logger) feedComponentInterfaces.FeedRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewFeedPostgreSQLRepository function.")

	feedRepository := &feedPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("FeedPostgreSQLRepository has created.")

	return feedRepository
}

func (fpsr *feedPostgreSQLRepository) getArticlesString(articles []*models.Article) string {
	articlesString := ""
	for _, v := range articles {
		articlesString += fmt.Sprintf("%#v\n", v)
	}

	return articlesString
}

func (fpsr *feedPostgreSQLRepository) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetTagsForArticle function.")

	tags := make([]string, 0, 10)
	rows, err := fpsr.database.Query(`
SELECT T.tag_name
FROM tags T
         JOIN tags_articles TA ON T.tag_id = TA.tag_id
         JOIN articles A ON TA.article_id = A.article_id
WHERE A.article_id = $1;
`, articleId)
	if err != nil {
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.FeedRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		tag := ""
		if err := rows.Scan(&tag); err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

// GetFeed TODO OFFSET LIMIT
func (fpsr *feedPostgreSQLRepository) GetFeed(ctx context.Context, email string) ([]*models.Article, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeed function.")

	articles := make([]*models.Article, 0, 10)

	rows, err := fpsr.database.Query(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
       COALESCE(A.cover_img_path, ''),
       COALESCE(COALESCE(UC.username, ''), ''),
       COALESCE(UC.login, ''),
       COALESCE(UP.username, ''),
       UP.login,
       COALESCE(C.category_name, ''),
       (CASE WHEN AL.is_like = true THEN 1 ELSE (CASE WHEN AL.is_like = false THEN -1 ELSE 0 END) END) liked
FROM articles A
         LEFT JOIN users UC ON A.co_author_id = UC.user_id
         JOIN users UP ON A.publisher_id = UP.user_id
         LEFT JOIN categories C ON A.category_id = C.category_id
         LEFT JOIN articles_likes AL ON AL.user_id = (SELECT user_id FROM users WHERE email = $1) AND AL.article_id = A.article_id
ORDER BY A.article_id DESC;
`, email)
	if err != nil {
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.FeedRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(&article.ArticleId, &article.Title, &article.Description, &article.Rating, &article.CommentsCount, &article.CoverImgPath, &article.CoAuthor.Username, &article.CoAuthor.Login, &article.Publisher.Username, &article.Publisher.Login, &article.CategoryName, &article.Liked); err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}
		article.Tags, err = fpsr.GetTagsForArticle(ctx, article.ArticleId)
		if err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}

		articles = append(articles, article)
	}

	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got articles: \n" + fpsr.getArticlesString(articles))

	return articles, nil
}

// GetFeedForUserByLogin TODO OFFSET LIMIT
func (fpsr *feedPostgreSQLRepository) GetFeedForUserByLogin(ctx context.Context, login string, email string) ([]*models.Article, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeedForUserByLogin function.")

	articles := make([]*models.Article, 0, 10)

	rows, err := fpsr.database.Query(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
       COALESCE(A.cover_img_path, ''),
       COALESCE(COALESCE(UC.username, ''), ''),
       COALESCE(UC.login, ''),
       COALESCE(UP.username, ''),
       UP.login,
       COALESCE(C.category_name, ''),
       (CASE WHEN AL.is_like = true THEN 1 ELSE (CASE WHEN AL.is_like = false THEN -1 ELSE 0 END) END) liked
FROM articles A
         LEFT JOIN users UC ON A.co_author_id = UC.user_id
         JOIN users UP ON A.publisher_id = UP.user_id
         LEFT JOIN categories C ON A.category_id = C.category_id
         LEFT JOIN articles_likes AL ON AL.user_id = (SELECT user_id FROM users WHERE email = $2) AND AL.article_id = A.article_id
WHERE UP.login = $1
ORDER BY A.article_id DESC;
`, login, email)
	if err != nil {
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.FeedRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(&article.ArticleId, &article.Title, &article.Description, &article.Rating, &article.CommentsCount, &article.CoverImgPath, &article.CoAuthor.Username, &article.CoAuthor.Login, &article.Publisher.Username, &article.Publisher.Login, &article.CategoryName, &article.Liked); err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}
		article.Tags, err = fpsr.GetTagsForArticle(ctx, article.ArticleId)
		if err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}

		articles = append(articles, article)
	}

	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got articles: \n" + fpsr.getArticlesString(articles))

	return articles, nil
}

// GetFeedForCategory TODO OFFSET LIMIT
func (fpsr *feedPostgreSQLRepository) GetFeedForCategory(ctx context.Context, category string, email string) ([]*models.Article, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeedForCategory function.")

	articles := make([]*models.Article, 0, 10)

	rows, err := fpsr.database.Query(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
       COALESCE(A.cover_img_path, ''),
       COALESCE(COALESCE(UC.username, ''), ''),
       COALESCE(UC.login, ''),
       COALESCE(UP.username, ''),
       UP.login,
       COALESCE(C.category_name, ''),
       (CASE WHEN AL.is_like = true THEN 1 ELSE (CASE WHEN AL.is_like = false THEN -1 ELSE 0 END) END) liked
FROM articles A
         LEFT JOIN users UC ON A.co_author_id = UC.user_id
         JOIN users UP ON A.publisher_id = UP.user_id
         LEFT JOIN categories C ON A.category_id = C.category_id
         LEFT JOIN articles_likes AL ON AL.user_id = (SELECT user_id FROM users WHERE email = $2) AND AL.article_id = A.article_id
WHERE C.category_name = $1
ORDER BY A.article_id DESC;
`, category, email)
	if err != nil {
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.FeedRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(&article.ArticleId, &article.Title, &article.Description, &article.Rating, &article.CommentsCount, &article.CoverImgPath, &article.CoAuthor.Username, &article.CoAuthor.Login, &article.Publisher.Username, &article.Publisher.Login, &article.CategoryName, &article.Liked); err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}
		article.Tags, err = fpsr.GetTagsForArticle(ctx, article.ArticleId)
		if err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}

		articles = append(articles, article)
	}

	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got articles: \n" + fpsr.getArticlesString(articles))

	return articles, nil
}

func (fpsr *feedPostgreSQLRepository) UserExistsByLogin(ctx context.Context, login string) (bool, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserExistsByLogin function.")

	row := fpsr.database.QueryRow(`
SELECT U.login
FROM users U WHERE U.login = $1;
`, login)

	loginTmp := ""
	if err := row.Scan(&loginTmp); err != nil {
		if err == sql.ErrNoRows {
			fpsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.FeedRepositoryError
	}

	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got login: ", loginTmp)

	return true, nil
}

func (fpsr *feedPostgreSQLRepository) CategoryExists(ctx context.Context, category string) (bool, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CategoryExists function.")

	row := fpsr.database.QueryRow(`
SELECT C.category_name
FROM categories C WHERE C.category_name = $1;
`, category)

	categoryTmp := ""
	if err := row.Scan(&categoryTmp); err != nil {
		if err == sql.ErrNoRows {
			fpsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.FeedRepositoryError
	}

	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got category: ", categoryTmp)

	return true, nil
}

func (fpsr *feedPostgreSQLRepository) GetNewArticlesFromIdForSubscriber(ctx context.Context, articleId int, email string) ([]int, error) {
	fpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetNewArticlesFromIdForSubscriber function.")

	newArticlesIds := make([]int, 0, 4)

	rows, err := fpsr.database.Query(`
SELECT A.article_id
FROM articles A
    	 JOIN users U ON U.email = $2
    	 LEFT JOIN users_categories_subscriptions UCS ON UCS.category_id = A.category_id AND UCS.user_id = U.user_id
    	 LEFT JOIN subscriptions S ON S.user_id = U.user_id
WHERE (A.article_id > $1) AND (A.publisher_id = S.subscribed_to_id OR A.category_id = UCS.category_id)
ORDER BY A.article_id DESC;
`, articleId, email)
	if err != nil {
		fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.FeedRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		var articleId *int
		if err := rows.Scan(&articleId); err != nil {
			fpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.FeedRepositoryError
		}

		newArticlesIds = append(newArticlesIds, *articleId)
	}

	fpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got newArticles ids: %#v\n", newArticlesIds)

	return newArticlesIds, nil
}
