<div align=center> 

# `🔎` League Champions <br>

[![Wakatime](https://wakatime.com/badge/github/controlado/league-champions-go.svg)](https://wakatime.com/badge/github/controlado/league-champions-go)
![Coverage](https://img.shields.io/badge/Coverage-96.4%25-blue) <br>
[![Languages](https://img.shields.io/badge/Documentation-gray)](https://pkg.go.dev/github.com/controlado/league-champions-go)
![Go version](https://img.shields.io/github/go-mod/go-version/controlado/league-champions-go?color=blue)
[![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?style=flat&logo=discord&logoColor=white&color=blue)](https://discordapp.com/users/854886148455399436)

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
	"log"

	"github.com/controlado/league-champions-go"
)

func main() {
	champions, err := league.GetChampions("default")
	if err != nil {
		log.Println("Não foi possível requisitar os campeões")
		panic(err)
	}

	for _, champion := range champions {
		if champion.Name == "K'Sante" {
			fmt.Println(champion.Rarity)
		}
	}
}
```

    Output: kNoRarity
