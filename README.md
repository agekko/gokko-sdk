[![Maintainability](https://api.codeclimate.com/v1/badges/233ac79886aff9f1580a/maintainability)](https://codeclimate.com/repos/61ede84a526fc70163005543/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/233ac79886aff9f1580a/test_coverage)](https://codeclimate.com/repos/61ede84a526fc70163005543/test_coverage)

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