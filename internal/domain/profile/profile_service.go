package profile

import (
	"github.com/charmingruby/clize/internal/validation"
	"github.com/charmingruby/clize/pkg/cryptography"
)

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
	p, err := NewProfile(
		username,
		email,
		password,
	)
	if err != nil {
		return err
	}

	if _, err := ps.profileRepo.FindByUsername(username); err == nil {
		return &validation.UniqueValueViolationError{
			Entity:  "profile",
			Field:   "username",
			Message: validation.NewUniqueValueViolationErrorMessage("profile", "username"),
		}
	}

	return ps.profileRepo.Create(p)
}

func (ps *ProfileService) Login(
	username, password string,
) error {
	profile, err := ps.profileRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	err = cryptography.VerifyIfHashMatches(
		profile.Password,
		password,
	)
	if err != nil {
		return &validation.GenericValidationError{
			Message: "credentials do not match",
		}
	}

	return nil
}
