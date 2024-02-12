package profile

type ProfileRepository interface {
	Create(p *Profile) error
	FindById(id string) (*Profile, error)
	FindByEmail(email string) (*Profile, error)
	FindByUsername(username string) (*Profile, error)
}
