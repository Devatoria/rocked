# rocked

[![GoDoc](https://godoc.org/github.com/Devatoria/rocked?status.svg)](https://godoc.org/github.com/Devatoria/rocked)

rocked has been created because of the need to get a lot of information about Docker containers very often. Some of these operations could lead to timeout if done through the Docker daemon because of overload.

rocked do the same things as you could through the Docker daemon, but by using the filesystem data. You will not solicit the daemon anymore.

## How to use it

Just import the library in your project, create a new rocked and start to use it!
```go
package main

import (
	"fmt"

	"github.com/Devatoria/rocked"
)

func main() {
	r, err := rocked.NewRocked("/var/lib/docker")
	if err != nil {
		panic(err)
	}

	containers, err := r.ListContainers()
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		data, err := r.InspectContainer(container)
		if err != nil {
			panic(err)
		}

		fmt.Println(container)
		fmt.Println(data.Config.Hostname)
	}
}
```

Please read the godoc reference to have the full list of features.
