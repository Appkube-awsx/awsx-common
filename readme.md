# Documentation for getting lambda clients

This package is created for getting different type of aws service clients for synectiks appkube-awsx project

## List of function that are created in this package

| S.No. | Functions                | Accepted arguments | Return type (from aws)         |
|-------|--------------------------|--------------------|--------------------------------|
| 1.    | GetLambdaClient()        | (auth Auth)        | *lambda.Lambda                 |
| 2.    | GetCostClient()          | (auth Auth)        | *costexplorer.CostExplorer     |
| 3.    | GetCloudWatchLogClient() | (auth Auth)        | *cloudwatchlogs.CloudWatchLogs |

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

### Project structure
```
client
    - client.go
    - session.go

main.go (This is just for testing)
```
