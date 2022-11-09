package usecaseToDeliveryErrors

type RepositoryError struct {
	Err error
}

func (re *RepositoryError) Error() string {
	return re.Err.Error()
}

type ArticleDontExistsError struct {
	Err error
}

func (adee *ArticleDontExistsError) Error() string {
	return adee.Err.Error()
}
