package profile

import (
	"time"

	"github.com/charmingruby/clize/pkg/cryptography"
	"github.com/charmingruby/clize/pkg/uuid"
)

type Profile struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Verified     bool      `json:"verified"`
	CreatedAt    time.Time `json:"created_at"`
	LastActivity time.Time `json:"last_activity"`
}

func NewProfile(
	username,
	email,
	password string,
) (*Profile, error) {
	p := &Profile{
		ID:           uuid.GenerateUUID(),
		Username:     username,
		Email:        email,
		Verified:     false,
		CreatedAt:    time.Now(),
		LastActivity: time.Now(),
	}

	if err := p.ChangePassword(password); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Profile) ChangePassword(password string) error {
	p.Password = password
	return encPass(p)
}

func encPass(p *Profile) error {

	hashedPassword, err := cryptography.GenerateHash(p.Password)
	if err != nil {
		return err
	}

	p.Password = hashedPassword

	return nil
}
