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

```go 
    const (
        LAMBDA_CLIENT  = "lambda"
        COST_EXPLORER  = "costExplorer"
        CLOUDWATCH_LOG = "cloudWatchLog"
    )
```

## Example of using function

`example for lambda`
```go
    auth := client.Auth{
        "us-east-1",
        os.Getenv("AWS_CROSS_ARN"),
        os.Getenv("AWS_ACCKEY"),
        os.Getenv("AWS_SECKEY"),
        os.Getenv("AWS_EXTERNALID"),
    }

    lambdaClient := client.GetClient(auth, client.LAMBDA_CLIENT).(*lambda.Lambda)
```

### Project structure
```
client
    - client.go
    - session.go

main.go (This is just for testing)
```
