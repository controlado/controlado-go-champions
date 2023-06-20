package champions

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

	for index, champion := range champions {
		if champion.Name == "K'Sante" {
			fmt.Println(index, champion.ID)
		}
	}

	// Output: 65 897
}
