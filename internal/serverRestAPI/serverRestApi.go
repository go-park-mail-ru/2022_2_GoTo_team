package serverRestAPI

import (
	feedComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery"
	feedComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/repository"
	feedComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/usecase"
	sessionComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/delivery"
	sessionComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/repository"
	sessionComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/usecase"
	userComponentDelivery "2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery"
	userComponentRepository "2022_2_GoTo_team/internal/serverRestAPI/userComponent/repository"
	userComponentUsecase "2022_2_GoTo_team/internal/serverRestAPI/userComponent/usecase"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"

	"2022_2_GoTo_team/internal/serverRestAPI/utils/configReader"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

const (
	LAYER_DELIVERY   = "delivery"
	LAYER_USECASE    = "usecase"
	LAYER_REPOSITORY = "repository"
)

func Run(configFilePath string) {
	config, err := configReader.NewConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config settings: ")
	log.Println(config)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     config.AllowOriginsAddressesCORS,
			AllowMethods:     []string{http.MethodPost, http.MethodGet},
			AllowCredentials: true,
		},
	))

	e.Use(middleware.Recover())

	if err := configureServer(e, config); err != nil {
		e.Logger.Fatal(errorsUtils.WrapError("error while configuring server", err))
	}

	e.Logger.Fatal(e.Start(config.ServerAddress))
}

func configureServer(e *echo.Echo, config *configReader.Config) error {

	sessionDeliveryLogger, err := logger.NewLogger("sessionComponent", LAYER_DELIVERY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionUsecaseLogger, err := logger.NewLogger("sessionComponent", LAYER_USECASE, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	sessionRepositoryLogger, err := logger.NewLogger("sessionComponent", LAYER_REPOSITORY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userDeliveryLogger, err := logger.NewLogger("userComponent", LAYER_DELIVERY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userUsecaseLogger, err := logger.NewLogger("userComponent", LAYER_USECASE, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	userRepositoryLogger, err := logger.NewLogger("userComponent", LAYER_REPOSITORY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	feedComponentDeliveryLogger, err := logger.NewLogger("feedComponent", LAYER_DELIVERY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	feedComponentUsecaseLogger, err := logger.NewLogger("feedComponent", LAYER_USECASE, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}
	feedComponentRepositoryLogger, err := logger.NewLogger("feedComponent", LAYER_REPOSITORY, config.LogLevel, config.LogFilePath)
	if err != nil {
		return err
	}

	sessionRepository, err := sessionComponentRepository.NewSessionCustomRepository(sessionRepositoryLogger)
	if err != nil {
		return err
	}
	userRepository := userComponentRepository.NewUserCustomRepository(userRepositoryLogger)
	feedRepository := feedComponentRepository.NewFeedCustomRepository(feedComponentRepositoryLogger)

	sessionUsecase := sessionComponentUsecase.NewSessionUsecase(sessionRepository, userRepository, sessionUsecaseLogger)
	sessionController := sessionComponentDelivery.NewSessionController(sessionUsecase, sessionDeliveryLogger)

	userUsecase := userComponentUsecase.NewUserUsecase(userRepository, userUsecaseLogger)
	userController := userComponentDelivery.NewUserController(userUsecase, sessionUsecase, userDeliveryLogger)

	feedUsecase := feedComponentUsecase.NewFeedUsecase(feedRepository, feedComponentUsecaseLogger)
	feedController := feedComponentDelivery.NewFeedController(feedUsecase, feedComponentDeliveryLogger)

	e.POST("/api/v1/session/create", sessionController.CreateSessionHandler)
	e.POST("/api/v1/session/remove", sessionController.RemoveSessionHandler)
	e.GET("/api/v1/session/info", sessionController.SessionInfoHandler)

	//e.POST("/api/v1/article/create", Api.CreateArticleHandler)
	//e.POST("/api/v1/article/update", Api.UpdateArticleHandler)

	e.POST("/api/v1/user/signup", userController.SignupUserHandler)
	//e.GET("/api/v1/user/info", Api.UserInfoHandler)
	//e.GET("/api/v1/user/feed", Api.UserFeedHandler)

	//e.GET("/api/v1/category/info", Api.CategoryInfoHandler)
	//e.GET("/api/v1/category/feed", Api.CategoryFeedHandler)

	e.GET("/api/v1/feed", feedController.FeedHandler)

	return nil
}
