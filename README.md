![Datacentred](https://assets-cdn.datacentred.io/assets/DC_Mono_B-903aac5ca4f5c6887193d880dbd1196deb8a978027eef5cb32de78b66d085935.png)

[![CircleCI](https://circleci.com/gh/datacentred/datacentred-go.svg?style=svg)](https://circleci.com/gh/datacentred/datacentred-go) [![Go Report Card](https://goreportcard.com/badge/github.com/datacentred/datacentred-go)](https://goreportcard.com/report/github.com/datacentred/datacentred-go) [![GoDoc](https://godoc.org/github.com/datacentred/datacentred-go?status.svg)](https://godoc.org/github.com/datacentred/datacentred-go)

Go client library for automating DataCentred account management.

[www.datacentred.co.uk](https://www.datacentred.co.uk)

# Installation

Run:

```
go get github.com/datacentred/datacentred-go
```

In your code:

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

2. Set your credentials by exporting your access key and secret key as environment variables in your shell:

```
export DATACENTRED_ACCESS="my_access"
export DATACENTRED_SECRET="my_secret"
```

Or set your keys manually using the following code:

```go
datacentred.Config.AccessKey = "my_access"
datacentred.Config.SecretKey = "my_secret"
```

## Usage Examples

### List all available users

```go
users := datacentred.Users()
fmt.Println(users)
// => [{2bd21ee25cde40fdb9454954e4fbb4b5 bill.s.preston@esquire.com Bill Preston 2015-02-13 11:07:00 +0000 UTC 2017-09-26 09:11:38 +0000 UTC } {69a34c127dcb439fa9366762234687ac ted.theodore@logan.com Ted Logan 2014-08-22 14:32:31 +0000 UTC 2017-09-21 14:55:43 +0000 UTC }]
```

### Find a user by id

```go
user := datacentred.FindUser("2bd21ee25cde40fdb9454954e4fbb4b5")
fmt.Println(user)
// => {2bd21ee25cde40fdb9454954e4fbb4b5 bill.s.preston@esquire.com Bill Preston 2015-02-13 11:07:00 +0000 UTC 2017-09-26 09:11:38 +0000 UTC } 
```

### Update a project

```go
project := datacentred.FindProject("37033518a4514f12adeb8346ac3f188c")
project.QuotaSet.Compute.Instances = 50
project.Save
fmt.Println(project)
// => [{37033518a4514f12adeb8346ac3f188c seancentred {{40 50 60000} {40 10 5} {0 10 50 10 10 100 10}} 2015-04-09 08:14:19 +0000 UTC 2016-12-08 11:44:05 +0000 UTC}
```

### Create a role

```go
params := map[string]interface{}{
  "role": map[string]interface{}{
    "name": "Wyld Stallyns",
  },
}

role := CreateRole(params)
fmt.Println(role)
// => {5713b281-b9f7-41d7-bc8c-9eb92920d1d3 Wyld Stallyns false [] 2017-09-26 09:42:56 +0000 UTC 2017-09-26 09:42:56 +0000 UTC}
```

### Add a user to a role

```go
user := datacentred.Users()[0]
fmt.Println(role.AddUser(user.Id))
// => true
```

### Remove a user from a project

```go
user := datacentred.Users()[0]
fmt.Println(project.RemoveUser(user.Id))
// => true
```

### Get usage data for a given year and month


```go
usage := datacentred.ShowUsage()
fmt.Println(usage.Projects[0].Usage.Instances[0].Usage.Value)
// => 744
fmt.Println(usage.Projects[0].Usage.Instances[0].Usage.Unit)
// => "hours"
```

## Documentation

TODO: add

## API Reference Manual

Please check out the [DataCentred API Documentation](https://my.datacentred.io/api/docs/v1) for a comprehensive description of the API itself.
