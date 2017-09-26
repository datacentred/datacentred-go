package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullProjectLifeCycle(t *testing.T) {
	r := initRecorder("fixtures/project_lifecycle1")
	defer r.Stop()

	projects := ListProjects()
	projectsCount := len(projects)

	firstProject := projects[0]

	assert.Equal(t, firstProject.Id, FindProject(firstProject.Id).Id, "they should be equal")

	params := map[string]interface{}{
		"project": map[string]interface{}{
			"name": "SanDimasHigh",
		},
	}

	newProject := CreateProject(params)

	r2 := initRecorder("fixtures/project_lifecycle2")
	defer r2.Stop()

	assert.Equal(t, newProject.Id, FindProject(newProject.Id).Id, "they should be equal")

	assert.Equal(t, "SanDimasHigh", newProject.Name, "they should be equal")

	r3 := initRecorder("fixtures/project_lifecycle3")
	defer r3.Stop()

	assert.Equal(t, projectsCount+1, len(ListProjects()), "they should be equal")

	newProject.Name = "BattleOfTheBands"
	newProject.Save()

	assert.Equal(t, "BattleOfTheBands", FindProject(newProject.Id).Name, "they should be equal")

	usersCount := len(newProject.Users())
	assert.Equal(t, 0, usersCount, "they should be equal")

	firstUser := ListUsers()[0]

	r4 := initRecorder("fixtures/project_lifecycle4")
	defer r4.Stop()

	result := newProject.AddUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	assert.Equal(t, usersCount+1, len(newProject.Users()), "they should be equal")

	result = newProject.RemoveUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	r5 := initRecorder("fixtures/project_lifecycle5")
	defer r5.Stop()

	assert.Equal(t, usersCount, len(newProject.Users()), "they should be equal")

	result = newProject.Destroy()
	assert.Equal(t, true, result, "they should be equal")
}
