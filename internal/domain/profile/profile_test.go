package profile

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileJSONSerialization(t *testing.T) {
	profile, err := NewProfile(
		"username",
		"email@email.com",
		"password",
	)

	assert.NoError(t, err)

	// Serialization
	data, err := json.Marshal(profile)
	assert.NoError(t, err)

	// Deserialization
	var newProfile Profile
	err = json.Unmarshal(data, &newProfile)
	assert.NoError(t, err)

	assert.Equal(t, profile.ID, newProfile.ID)
	assert.Equal(t, profile.Username, newProfile.Username)
	assert.Equal(t, profile.Password, newProfile.Password)
	assert.Equal(t, profile.Email, newProfile.Email)
}
