# Documentation for getting lambda clients

This package is created for getting different type of aws service clients for synectiks appkube-awsx project

## List of function that are created in this package

| S.No. | Functions            | Accepted arguments              | Return type (from aws) |
|-------|----------------------|---------------------------------|------------------------|
| 1.    | GetClient()          | (auth Auth, clientType string)  | interface{}            |

These function accepts 'auth' struct that is following

```go
    type Auth struct {
        Region              string
        CrossAccountRoleArn string
        AccessKey           string
        SecretKey           string
        ExternalId          string
    }
```

### List of client types
1. lambda
2. costExplorer
3. cloudWatchLog

### Project structure
```
client
    - client.go
    - session.go

main.go (This is just for testing)
```
