package assignment

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignmentJSONSerialization(t *testing.T) {
	assignment, err := NewAssignment(
		"get endpoint broken",
		"fix the GET application endpoint",
		1,
	)

	assert.NoError(t, err)

	// Serialization
	data, err := json.Marshal(assignment)
	assert.NoError(t, err)

	// Deserialization
	var newAssignment Assignment
	err = json.Unmarshal(data, &newAssignment)
	assert.NoError(t, err)

	assert.Equal(t, assignment.Title, newAssignment.Title)
	assert.Equal(t, assignment.Description, newAssignment.Description)
	assert.Equal(t, assignment.Status, newAssignment.Status)
}
