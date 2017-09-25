![Datacentred](https://assets-cdn.datacentred.io/assets/DC_Mono_B-903aac5ca4f5c6887193d880dbd1196deb8a978027eef5cb32de78b66d085935.png)

Go client library for automating DataCentred account management.

[www.datacentred.co.uk](https://www.datacentred.co.uk)

# Installation

```go
import(
    "github.com/datacentred/datacentred-go"
)
```

# Usage

This API allows you to automate operations against your DataCentred account.

Operations include:

* Creating and managing users;
* Creating and managing roles for users;
* Managing OpenStack Project creation, quota adjustments, and user assignments;
* Viewing detailed usage/billing information for your account.

## Authentication

The API uses two pieces of information to authenticate access.

A unique access key specific to your DataCentred account, and a secret key which is generated once.

To get started:

1. Find your API access key and secret key at [my.datacentred.io](https://my.datacentred.io)

![API Credentials](https://user-images.githubusercontent.com/98526/30334767-79f4617c-97d8-11e7-962c-ec3115d13896.png)

2. Set your credentials by exporting your access key and secret key as environment variables:

```
export DATACENTRED_ACCESS="my_access"
export DATACENTRED_SECRET="my_secret"
```

Or setting your keys manually using the following methods:

```go
TODO: add example code
```

## Usage Examples

### List all available users

```go
package main

import (
    "fmt"
    "github.com/datacentred/datacentred-go"
)

func main() {
    users := datacentred.ListUsers()
    fmt.Printf("Number of users: %d\n", len(users))
    for i := range users {
        fmt.Printf("ID: %s, %s %s", users[i].Id, users[i].FirstName, users[i].LastName)
    }

}
```

### Find a user by id

```go
func main() {
    user := datacentred.FindUser("c165f2794d5941e78493275654572fd6")
    fmt.Println(user.FirstName, user.LastName, user.Email)
}
```

### Create a project

```go
func createproject(projectname string) string {
    params := map[string]interface{}{
        "name": projectname,
    }

    project := datacentred.CreateProject(params)
    return project.Id
}
```

### Update a project

### Create a role

### Add a user to a role

### Remove a user from a project

### Get usage data for a given year and month

## Documentation

TODO: add

## API Reference Manual

Please check out the [DataCentred API Documentation](https://my.datacentred.io/api/docs/v1) for a comprehensive description of the API itself.
