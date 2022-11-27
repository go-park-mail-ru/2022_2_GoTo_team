package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/profileComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/profileComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/pkg/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileController struct {
	profileUsecase profileComponentInterfaces.ProfileUsecaseInterface
	sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewProfileController(profileUsecase profileComponentInterfaces.ProfileUsecaseInterface, sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) *ProfileController {
	logger.LogrusLogger.Debug("Enter to the NewProfileController function.")

	profileController := &ProfileController{
		profileUsecase: profileUsecase,
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("ProfileController has created.")

	return profileController
}

func (pc *ProfileController) isAuthorized(c echo.Context) bool {
	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the isAuthorized function.")

	authorized := false
	if cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME); err == nil && cookie != nil {
		if authorized, err = pc.sessionUsecase.SessionExists(c.Request().Context(), &models.Session{SessionId: cookie.Value}); err != nil {
			return false
		}
	}

	return authorized
}

func (pc *ProfileController) GetProfileHandler(c echo.Context) error {
	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the GetProfileHandler function.")

	if !pc.isAuthorized(c) {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Info("Unauthorized!")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Info(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	profile, err := pc.profileUsecase.GetProfileBySession(c.Request().Context(), &models.Session{SessionId: cookie.Value})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionNotFoundError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusUnauthorized)
		case *usecaseToDeliveryErrors.UserForSessionNotFoundError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusUnauthorized)
		default:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	profileOutput := modelsRestApi.Profile{
		Email:         profile.Email,
		Login:         profile.Login,
		Password:      profile.Password,
		Username:      profile.Username,
		AvatarImgPath: profile.AvatarImgPath,
	}
	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed profileInfo = ", profileOutput)

	return c.JSON(http.StatusOK, profileOutput)
}

func (pc *ProfileController) UpdateProfileHandler(c echo.Context) error {
	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the UpdateProfileHandler function.")
	defer c.Request().Body.Close()

	if !pc.isAuthorized(c) {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Info("Unauthorized!")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Info(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	parsedInputProfile := new(modelsRestApi.Profile)
	if err := c.Bind(parsedInputProfile); err != nil {
		pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	pc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed parsedInputProfile: %#v", parsedInputProfile)

	err = pc.profileUsecase.UpdateProfileBySession(c.Request().Context(), &models.Profile{Email: parsedInputProfile.Email, Login: parsedInputProfile.Login, Username: parsedInputProfile.Username, AvatarImgPath: parsedInputProfile.AvatarImgPath}, &models.Session{SessionId: cookie.Value})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotValidError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusBadRequest, "email is not valid")
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusBadRequest, "login is not valid")
		case *usecaseToDeliveryErrors.PasswordIsNotValidError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusBadRequest, "password is not valid")
		case *usecaseToDeliveryErrors.EmailExistsError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusConflict, "email exists")
		case *usecaseToDeliveryErrors.LoginExistsError:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.JSON(http.StatusConflict, "login exists")
		default:
			pc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)
}
