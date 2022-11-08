package usecase

import (
	repositoryToUsecaseErrors2 "2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
	"errors"
)

type sessionUsecase struct {
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	userRepository    userComponentInterfaces.UserRepositoryInterface
	logger            *logger.Logger
}

func NewSessionUsecase(sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, userRepository userComponentInterfaces.UserRepositoryInterface, logger *logger.Logger) sessionComponentInterfaces.SessionUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewSessionUsecase function.")

	sessionUsecase := &sessionUsecase{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("SessionUsecase has created.")

	return sessionUsecase
}

func (su *sessionUsecase) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SessionExists function.")

	wrappingErrorMessage := "error while checking session exists"

	exists, err := su.sessionRepository.SessionExists(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return false, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return exists, nil
}

func (su *sessionUsecase) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateSessionForUser function.")

	wrappingErrorMessage := "error while creating session for user"

	exists, err := su.userRepository.CheckUserEmailAndPassword(ctx, email, password)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}
	if !exists {
		su.logger.LogrusLoggerWithContext(ctx).Warn("Incorrect email or password.")
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.IncorrectEmailOrPasswordError{Err: errors.New("incorrect email or password")})
	}

	session, err := su.sessionRepository.CreateSessionForUser(ctx, email)
	if err != nil {
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return session, nil
}

func (su *sessionUsecase) RemoveSession(ctx context.Context, session *models.Session) error {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")

	wrappingErrorMessage := "error while removing session"

	if err := su.sessionRepository.RemoveSession(ctx, session); err != nil {
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return nil
}

func (su *sessionUsecase) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfoByEmail function.")

	wrappingErrorMessage := "error while getting username by session"

	email, err := su.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		switch err {
		case repositoryToUsecaseErrors2.SessionRepositoryEmailDontExistsError:
			su.logger.LogrusLoggerWithContext(ctx).Debug("Trying to remove the garbage session: %#v", session)
			_ = su.RemoveSession(ctx, session) // We should try to remove "garbage" session
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDontFoundError{Err: err})
		default:
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	user, err := su.userRepository.GetUserInfoForSessionComponentByEmail(ctx, email)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		switch err {
		case repositoryToUsecaseErrors.UserRepositoryEmailDontExistsError:
			su.logger.LogrusLoggerWithContext(ctx).Debug("Trying to remove the garbage session: %#v", session)
			_ = su.RemoveSession(ctx, session) // We should try to remove "garbage" session
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.UserForSessionDontFoundError{Err: err})
		default:
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return user, nil
}