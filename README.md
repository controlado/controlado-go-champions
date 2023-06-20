<div align=center> 

![Coverage](https://img.shields.io/badge/Coverage-96.4%25-brightgreen)

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