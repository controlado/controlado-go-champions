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
		if champion.Name == "Nunu & Willump" {
			fmt.Println(index, champion.NameURL)

			for index, skin := range champion.Skins {
				fmt.Println(index, skin.Name)
			}
		}
	}
}
```

    Output:
    88 Nunu
    0 Demolisher Nunu & Willump
    1 Grungy Nunu & Willump
    2 Nunu & Beelump
    3 Nunu & Willump Bot
    4 Papercraft Nunu & Willump
    5 Sasquatch Nunu & Willump
    6 Space Groove Nunu & Willump
    7 TPA Nunu & Willump
    8 Workshop Nunu & Willump
    9 Zombie Nunu & Willump
