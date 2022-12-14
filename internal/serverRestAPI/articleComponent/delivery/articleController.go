package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/createArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/getArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/likeData"
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/removeArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/updateArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ArticleController struct {
	articleUsecase articleComponentInterfaces.ArticleUsecaseInterface
	logger         *logger.Logger
}

func NewArticleController(articleUsecase articleComponentInterfaces.ArticleUsecaseInterface, logger *logger.Logger) *ArticleController {
	logger.LogrusLogger.Debug("Enter to the NewArticleController function.")

	articleController := &ArticleController{
		articleUsecase: articleUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("ArticleController has created.")

	return articleController
}

func (ac *ArticleController) ArticleHandler(c echo.Context) error {
	ctx := c.Request().Context()
	ac.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the ArticleHandler function.")

	idStr := c.QueryParam("id")
	ac.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed id: %#v", idStr)
	if idStr == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if id < 1 {
		ac.logger.LogrusLoggerWithContext(ctx).Warn(errors.New(fmt.Sprintf("id = %d < 1", id)))
		return c.NoContent(http.StatusBadRequest)
	}

	article, err := ac.articleUsecase.GetArticleById(ctx, id)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.ArticleDoesntExistError:
			ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			ac.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	articleOutput := getArticle.Article{
		Id:           article.ArticleId,
		Title:        article.Title,
		Description:  article.Description,
		Tags:         article.Tags,
		Category:     article.CategoryName,
		Rating:       article.Rating,
		Comments:     article.CommentsCount,
		Content:      article.Content,
		CoverImgPath: article.CoverImgPath,
		Publisher: getArticle.Publisher{
			Username: article.Publisher.Username,
			Login:    article.Publisher.Login,
		},
		CoAuthor: getArticle.CoAuthor{
			Username: article.CoAuthor.Username,
			Login:    article.CoAuthor.Login,
		},
		Liked: article.Liked,
	}
	ac.logger.LogrusLoggerWithContext(ctx).Debug("Formed articleOutput: ", articleOutput)

	jsonBytes, err := articleOutput.MarshalJSON()
	if err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (ac *ArticleController) CreateArticleHandler(c echo.Context) error {
	ctx := c.Request().Context()
	ac.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateArticleHandler function.")
	defer c.Request().Body.Close()

	parsedInputArticle := new(createArticle.Article)
	if err := c.Bind(parsedInputArticle); err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	ac.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputArticle: %#v", parsedInputArticle)

	err := ac.articleUsecase.AddArticleBySession(ctx, &models.Article{Title: parsedInputArticle.Title, Description: parsedInputArticle.Description, Tags: parsedInputArticle.Tags, CategoryName: parsedInputArticle.Category, CoverImgPath: parsedInputArticle.CoverImgPath, Content: parsedInputArticle.Content, CoAuthor: models.CoAuthor{Login: parsedInputArticle.CoAuthorLogin}})
	if err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (ac *ArticleController) UpdateArticleHandler(c echo.Context) error {
	ctx := c.Request().Context()
	ac.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateArticleHandler function.")
	defer c.Request().Body.Close()

	parsedInputArticle := new(updateArticle.Article)
	if err := c.Bind(parsedInputArticle); err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	ac.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputArticle: %#v", parsedInputArticle)

	err := ac.articleUsecase.UpdateArticle(ctx, &models.Article{ArticleId: parsedInputArticle.Id, Title: parsedInputArticle.Title, Description: parsedInputArticle.Description, Tags: parsedInputArticle.Tags, CategoryName: parsedInputArticle.Category, Content: parsedInputArticle.Content})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotAuthorError:
			ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusForbidden)
		default:
			ac.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)
}

func (ac *ArticleController) RemoveArticleHandler(c echo.Context) error {
	ctx := c.Request().Context()
	ac.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveArticleHandler function.")
	defer c.Request().Body.Close()

	parsedInputArticleId := new(removeArticle.ArticleId)
	if err := c.Bind(parsedInputArticleId); err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	ac.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputArticleId: %#v", parsedInputArticleId)

	err := ac.articleUsecase.RemoveArticleById(ctx, parsedInputArticleId.Id)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.ArticleDoesntExistError:
			ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			ac.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)
}

func (ac *ArticleController) LikeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	ac.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the LikeHandler function.")
	defer c.Request().Body.Close()

	parsedInputLikeData := new(likeData.LikeData)
	if err := c.Bind(parsedInputLikeData); err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	sign := parsedInputLikeData.Sign
	if sign != -1 && sign != 0 && sign != 1 {
		ac.logger.LogrusLoggerWithContext(ctx).Warnf("Incorrect sign value = %#v, should be -1 or 0 or 1", sign)
		return c.NoContent(http.StatusBadRequest)
	}

	ac.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputLikeData: %#v", parsedInputLikeData)

	updatedRating, err := ac.articleUsecase.ProcessLike(ctx, &models.LikeData{Id: parsedInputLikeData.Id, Sign: parsedInputLikeData.Sign})
	if err != nil {
		ac.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	likeResponse := likeData.LikeResponse{
		Rating: updatedRating,
	}

	return c.JSON(http.StatusOK, likeResponse)
}
