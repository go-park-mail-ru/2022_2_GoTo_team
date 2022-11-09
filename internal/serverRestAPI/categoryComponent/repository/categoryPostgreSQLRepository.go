package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/categoryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
	"database/sql"
)

type categoryPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewCategoryPostgreSQLRepository(database *sql.DB, logger *logger.Logger) categoryComponentInterfaces.CategoryRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewCategoryPostgreSQLRepository function.")

	categoryRepository := &categoryPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("categoryPostgreSQLRepository has created.")

	return categoryRepository
}

func (cpsr *categoryPostgreSQLRepository) GetCategoryInfo(ctx context.Context, categoryName string) (*models.Category, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetCategoryInfo function.")

	row := cpsr.database.QueryRow(`
SELECT category_name, description, subscribers_count
FROM categories
WHERE category_name = $1;
`, categoryName)

	category := &models.Category{}
	if err := row.Scan(&category.CategoryName, &category.Description, &category.SubscribersCount); err != nil {
		if err == sql.ErrNoRows {
			cpsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.CategoryRepositoryCategoryDontExistsError
		}
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.CategoryRepositoryError
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got category: %#v", category)

	return category, nil
}