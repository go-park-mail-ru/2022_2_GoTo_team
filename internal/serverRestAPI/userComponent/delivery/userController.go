package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils/httpCookieUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUsecase    userComponentInterfaces.UserUsecaseInterface
	sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewUserController(userUsecase userComponentInterfaces.UserUsecaseInterface, sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) *UserController {
	logger.LogrusLogger.Debug("Enter to the NewUserController function.")

	userController := &UserController{
		userUsecase:    userUsecase,
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("UserController has created.")

	return userController
}

func (uc *UserController) SignupUserHandler(c echo.Context) error {
	ctx := c.Request().Context()
	uc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SignupUserHandler function.")
	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.User)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	uc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed input user json data: %#v", parsedInput)

	if err := uc.userUsecase.AddNewUser(ctx, parsedInput.NewUserData.Email, parsedInput.NewUserData.Login, parsedInput.NewUserData.Username, parsedInput.NewUserData.Password); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotValidError:
			return c.JSON(http.StatusBadRequest, "email is not valid")
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			return c.JSON(http.StatusBadRequest, "login is not valid")
		case *usecaseToDeliveryErrors.PasswordIsNotValidError:
			return c.JSON(http.StatusBadRequest, "password is not valid")
		case *usecaseToDeliveryErrors.EmailExistsError:
			return c.JSON(http.StatusConflict, "email exists")
		case *usecaseToDeliveryErrors.LoginExistsError:
			return c.JSON(http.StatusConflict, "login exists")
		default:
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	session, err := uc.sessionUsecase.CreateSessionForUser(ctx, parsedInput.NewUserData.Email, parsedInput.NewUserData.Password)
	if err != nil {
		uc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	c.SetCookie(httpCookieUtils.MakeHttpCookie(session.SessionId))

	uc.logger.LogrusLoggerWithContext(ctx).Infof("User with the email %#v registered successful!", parsedInput.NewUserData.Email)

	return c.NoContent(http.StatusOK)
}

func (uc *UserController) UserInfoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	uc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserInfoHandler function.")

	login := c.QueryParam("login")
	uc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed login: %#v", login)
	if login == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := uc.userUsecase.GetUserInfo(ctx, login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.LoginDoesntExistError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusBadRequest) // TODO
		default:
			uc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	isSubscribed, err := uc.userUsecase.IsUserSubscribedOnUser(ctx, login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		default:
			uc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	userInfo := modelsRestApi.UserInfo{
		Username:         user.Username,
		RegistrationDate: user.RegistrationDate,
		SubscribersCount: user.SubscribersCount,
		Subscribed:       isSubscribed,
	}

	uc.logger.LogrusLoggerWithContext(ctx).Debug("Formed userInfo: ", userInfo)

	jsonBytes, err := userInfo.MarshalJSON()
	if err != nil {
		uc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (uc *UserController) GetUserAvatar(c echo.Context) error {
	ctx := c.Request().Context()
	uc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserAvatar function.")

	login := c.QueryParam("login")
	uc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed login: %#v", login)
	if login == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := uc.userUsecase.GetUserAvatar(ctx, login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.LoginDoesntExistError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusBadRequest) // TODO
		default:
			uc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	userAvatar := modelsRestApi.UserAvatar{
		AvatarImgPath: user.AvatarImgPath,
	}

	uc.logger.LogrusLoggerWithContext(ctx).Debug("Formed userAvatar: ", userAvatar)

	jsonBytes, err := userAvatar.MarshalJSON()
	if err != nil {
		uc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (uc *UserController) SubscribeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	uc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SubscribeHandler function.")
	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.Subscribe)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	uc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed input json data: %#v", parsedInput)

	if err := uc.userUsecase.SubscribeOnUser(ctx, parsedInput.Login); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusInternalServerError)
		default:
			uc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	uc.logger.LogrusLoggerWithContext(ctx).Info("User subscribed successfully!")

	return c.NoContent(http.StatusOK)
}

func (uc *UserController) UnsubscribeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	uc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UnsubscribeHandler function.")
	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.Subscribe)
	if err := c.Bind(parsedInput); err != nil {
		uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	uc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed input json data: %#v", parsedInput)

	if err := uc.userUsecase.UnsubscribeFromUser(ctx, parsedInput.Login); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			uc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusInternalServerError)
		default:
			uc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	uc.logger.LogrusLoggerWithContext(ctx).Info("User unsubscribed successfully!")

	return c.NoContent(http.StatusOK)
}
