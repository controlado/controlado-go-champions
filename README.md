<div align=center> 

# `🌱` League Champions <br>

![coverage](https://img.shields.io/badge/coverage-95.5%25-brightgreen)
[![wakatime](https://wakatime.com/badge/github/controlado/league-champions-go.svg)](https://wakatime.com/badge/github/controlado/league-champions-go) <br>
[![languages](https://img.shields.io/badge/Documentation-gray)](https://pkg.go.dev/github.com/controlado/league-champions-go)
![latest](https://img.shields.io/github/v/tag/controlado/league-champions-go?label=Latest)
![go](https://img.shields.io/github/go-mod/go-version/controlado/league-champions-go?color=blue)

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
	lol, err := league.New("default")
	if err != nil {
		panic(err)
	}
    
	for index, champion := range lol.Champions{
		if champion.ID == 4 {
			fmt.Println(index, champion.Name)
		}
	}
}
```

    Output: int Twisted Fate
