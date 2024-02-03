package application

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppJSONSerialization(t *testing.T) {
	app, _ := NewApplication(
		"clize",
		"productivity cli",
	)

	// Serialization
	data, err := json.Marshal(app)
	assert.NoError(t, err)

	// Deserialization
	var newApp *Application
	err = json.Unmarshal(data, &newApp)
	assert.NoError(t, err)

	assert.Equal(t, app.Name, newApp.Name)
	assert.Equal(t, app.Context, newApp.Context)
	assert.Equal(t, app.Assignments, newApp.Assignments)
}
