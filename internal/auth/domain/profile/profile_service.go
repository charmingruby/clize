package profile

import (
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/charmingruby/clize/pkg/hash"
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
		return &errors.UniqueValueViolationError{
			Entity:  "profile",
			Field:   "username",
			Message: errors.NewUniqueValueViolationErrorMessage("profile", "username"),
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

	err = hash.VerifyIfHashMatches(
		profile.Password,
		password,
	)
	if err != nil {
		return &errors.GenericValidationError{
			Message: "credentials do not match",
		}
	}

	return nil
}
