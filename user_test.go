package datacentred

import (
  "fmt"
  "testing"
  "net/http"
  "github.com/dnaeon/go-vcr/recorder"
)

func init() {
  r, err := recorder.New("fixtures/user_lifecycle")
  if err != nil {
    fmt.Errorf("Recorder failed: %s", err)
  }
  defer r.Stop()

  http_config.Client = http.Client{
    Transport: r,
  } 
}

func TestFullUserLifeCycle(t *testing.T) {
  users := ListUsers()
  fmt.Println(users[0])
}

  // projects := datacentred.ListProjects()
  // fmt.Println(projects)

  // usage := datacentred.ShowUsage(2017, 6)
  // fmt.Println(usage.Projects[3].Usage.Vpns[0].Usage[0].Value)

  // roles := datacentred.ListRoles()
  // fmt.Println(roles[2])

  // user := datacentred.FindUser("2bd21ee25cde40fdb9454954e4fbb4b5")
  // fmt.Println(user)

  // role := datacentred.FindRole("654f423e-646a-4742-849d-d8c9ab9b4f39")
  // fmt.Println(role)

 //  project := datacentred.FindProject("37033518a4514f12adeb8346ac3f188c")
 //  fmt.Println(project)

  // input := map[string]interface{}{
  //   "name": "sadsdsa",
  // }
  // fmt.Println(input)

  // datacentred.CreateProject(input)

  // input := map[string]interface{}{
  //   "name": "ffdsfsd",
  // }
  // fmt.Println(input)

  // project := datacentred.FindProject("dd7bfc05def9404784dab5f660f42d11")
  // project.Name = "fdsjkfdsjfaf"
  // project2 := datacentred.UpdateProject("dd7bfc05def9404784dab5f660f42d11", project)
  // fmt.Println(project2.Name)

  // fmt.Println(datacentred.ListProjects())

  // datacentred.DestroyProject("a4ee704fba9d4997be72ff1bae6a6ac1")

  // project := datacentred.FindProject("3e516421b80a4b70aa816aef1dfd79fd")
  // fmt.Println(project.Users())

  // fmt.Println(datacentred.ListRoles()[2].Users())
  // fmt.Println(datacentred.ListUsers()[1])
  // datacentred.ListRoles()[2].AddUser("69a34c127dcb439fa9366762234687ac")
  // fmt.Println(datacentred.ListRoles()[2].Users())
// }
