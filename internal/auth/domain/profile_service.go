package domain

type ProfileService struct {
	profileRepo ProfileRepository
}

func NewProfileService(profileRepo ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
	}
}

func (ps *ProfileService) Register(
	username, email, password string,
) error {
	return nil
}

func (ps *ProfileService) Login(
	username, password string,
) error {
	return nil
}
