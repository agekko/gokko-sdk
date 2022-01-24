# Gokko SDK
Basic SDK for integrating services with GCP features.

## Google

### Project

#### Get Project ID
```go
import (
    "fmt"

    gokko "agekko/google-sdk"
)

func main() {
    projectId := gokko.Project_GetProjectId()
    fmt.Println("Project ID:", projectId)
}
```