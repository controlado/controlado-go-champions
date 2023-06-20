package league

import (
	"fmt"
	"log"
)

func ExampleGetChampions() {
	champions, err := GetChampions("default")
	if err != nil {
		log.Println("Não foi possível requisitar os campeões")
		panic(err)
	}

	for _, champion := range champions {
		if champion.Name == "K'Sante" {
			fmt.Println(champion.Rarity)
		}
	}

	// Output: kNoRarity
}
