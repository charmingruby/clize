package domain

import "github.com/charmingruby/clize/internal/auth/domain/profile"

type Service struct {
	ProfileService *profile.ProfileService
}
