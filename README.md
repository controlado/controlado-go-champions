<div align=center> 

# `🌱` League Champions <br>

![coverage](https://img.shields.io/badge/coverage-96.2%25-brightgreen)
[![wakatime](https://wakatime.com/badge/github/controlado/league-champions-go.svg)](https://wakatime.com/@programador/projects/pjllknspfy) <br>
[![report](https://goreportcard.com/badge/github.com/controlado/league-champions-go)](https://goreportcard.com/report/github.com/controlado/league-champions-go)
[![documentation](https://img.shields.io/badge/documentation-lighgreen)](https://pkg.go.dev/github.com/controlado/league-champions-go)
[![latest](https://img.shields.io/github/v/tag/controlado/league-champions-go?label=latest)](https://github.com/controlado/league-champions-go/commits)
[![go](https://img.shields.io/github/go-mod/go-version/controlado/league-champions-go?label=go&color=blue)](https://go.dev/dl/)

This is my first project in Go! <br>
I hope it can be useful for you :)

</div>
<br>

## Use in your project

    go get github.com/controlado/league-champions-go

```go
package main

import (
	"fmt"

	"github.com/controlado/league-champions-go"
)

func main() {
	lol, err := league.New("default")
	if err != nil {
		panic(err)
	}

	for index, champion := range lol.Champions {
		if champion.ID == 4 || champion.ID == 20 {
			fmt.Println(index, champion.Name)
		}
	}
}
```

    Output:
    int Twisted Fate
    int Nunu & Willump
