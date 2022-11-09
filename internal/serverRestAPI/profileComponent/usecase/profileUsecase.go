package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/profileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/profileComponentErrors/usecaseToDeliveryErrors"
	repositoryToUsecaseErrors_sessionComponent "2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
)

type profileUsecase struct {
	profileRepository profileComponentInterfaces.ProfileRepositoryInterface
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	logger            *logger.Logger
}

func NewProfileUsecase(profileRepository profileComponentInterfaces.ProfileRepositoryInterface, sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, logger *logger.Logger) profileComponentInterfaces.ProfileUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewCategoryUsecase function.")

	profileUsecase := &profileUsecase{
		profileRepository: profileRepository,
		sessionRepository: sessionRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("profileUsecase has created.")

	return profileUsecase
}

func (pu *profileUsecase) GetProfileBySession(ctx context.Context, session *models.Session) (*models.Profile, error) {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileBySession function.")

	wrappingErrorMessage := "error while getting profile by session"

	email, err := pu.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors_sessionComponent.SessionRepositoryEmailDontExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDontFoundError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	profile, err := pu.profileRepository.GetProfileByEmail(ctx, email)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.ProfileRepositoryEmailDontExistsError:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.UserForSessionDontFoundError{Err: err})
		default:
			pu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return profile, nil
}
