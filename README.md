# LeCloud Go SDK
Golang SDK for LeCloud

## Installation:
```bash
go get -u github.com/ricer0ll/lecloud-go
```

## Usage Example
```go
package main

import (
	"fmt"

	"github.com/ricer0ll/lecloud-go/lecloud"
	"resty.dev/v3"
)

func main() {
	credentials := lecloud.NewCredentials(
        "bob@gmail.com", 
        "passGoesHere", 
        "accountIdGoesHere")
    
	restyClient := resty.New()

	client := lecloud.NewClient(restyClient, credentials, apiKeyGoesHere)
	fmt.Println(client.GetSecret("secretUuidGoesHere"))
}
```
