package datacentred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullProjectLifeCycle(t *testing.T) {
	r := initRecorder("fixtures/project_lifecycle1")
	defer r.Stop()

	projects, _ := Projects()
	projectsCount := len(projects)

	firstProject := projects[0]

  p, _ := FindProject(firstProject.Id)
	assert.Equal(t, firstProject.Id, p.Id, "they should be equal")

	params := map[string]string{
		"name": "SanDimasHigh",
	}

	newProject, _ := CreateProject(params)

	r2 := initRecorder("fixtures/project_lifecycle2")
	defer r2.Stop()

  p, _ = FindProject(newProject.Id)
	assert.Equal(t, newProject.Id, p.Id, "they should be equal")

	assert.Equal(t, "SanDimasHigh", newProject.Name, "they should be equal")

	r3 := initRecorder("fixtures/project_lifecycle3")
	defer r3.Stop()

	projects, _ = Projects()

	assert.Equal(t, projectsCount+1, len(projects), "they should be equal")

	newProject.Name = "BattleOfTheBands"
	newProject.Save()

  p, _ = FindProject(newProject.Id)
	assert.Equal(t, "BattleOfTheBands", p.Name, "they should be equal")

	users, _ := newProject.Users()

	usersCount := len(users)
	assert.Equal(t, 0, usersCount, "they should be equal")

  users, _ = Users()
	firstUser := users[0]

	r4 := initRecorder("fixtures/project_lifecycle4")
	defer r4.Stop()

	result, _ := newProject.AddUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	users, _ = newProject.Users()

	assert.Equal(t, usersCount+1, len(users), "they should be equal")

	result, _ = newProject.RemoveUser(firstUser.Id)
	assert.Equal(t, true, result, "they should be equal")

	r5 := initRecorder("fixtures/project_lifecycle5")
	defer r5.Stop()

	users, _ = newProject.Users()

	assert.Equal(t, usersCount, len(users), "they should be equal")

	result, _ = newProject.Destroy()
	assert.Equal(t, true, result, "they should be equal")
}

func TestProjectErrors(t *testing.T) {
	r := initRecorder("fixtures/project_errors")
	defer r.Stop()	

  _, err := FindProject("bogus")
  assert.Equal(t, "Not found", err.Error(), "they should be equal")
}
