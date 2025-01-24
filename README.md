# healthchecksio-api

Golang package for [Healthchecks.io](https://healthchecks.io/) API

API reference - [pkg.go.dev](https://pkg.go.dev/github.com/gfx687/healthcheckio-api)

### Example usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/gfx687/healthchecksio-api"
)

var (
	healthcheckUrl string = "https://hc-ping.com/<CHECK_ID>"
	processID      string = "some_uuid"
)

func main() {
	err := healthchecksio.Healthcheck(healthcheckUrl, healthchecksio.HealthcheckStart, processID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = do_something()
	if err != nil {
		fmt.Println(err)

		err = healthchecksio.Healthcheck(healthcheckUrl, healthchecksio.HealthcheckFail, processID)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	err = healthchecksio.Healthcheck(healthcheckUrl, healthchecksio.HealthcheckSuccess, processID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
