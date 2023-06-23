package league

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// Testa as funcionalidades principais
// que regem a instância de League.
func TestNew(t *testing.T) {
	_, err := New("default")
	if err != nil {
		t.Error("expected nil got", err)
	}

	_, err = New("invalid")
	if err == nil {
		t.Error("expected error got nil")
	}
}

// Exemplifica o uso de New.
func ExampleNew() {
	lol, err := New("default")
	if err != nil {
		log.Panicln(err)
	}

	for index, champion := range lol.Champions {
		if champion.Name == "Nunu & Willump" {
			fmt.Println(index, champion.NameURL)

			for index, skin := range champion.Skins {
				fmt.Println(index, skin.Name)
			}
		}
	}

	// Output:
	// 19 Nunu
	// 0 Sasquatch Nunu & Willump
	// 1 Workshop Nunu & Willump
	// 2 Grungy Nunu & Willump
	// 3 Nunu & Willump Bot
	// 4 Demolisher Nunu & Willump
	// 5 TPA Nunu & Willump
	// 6 Zombie Nunu & Willump
	// 7 Papercraft Nunu & Willump
	// 8 Space Groove Nunu & Willump
	// 9 Nunu & Beelump
}

// Exemplifica o uso de GetChampions.
func ExampleGetChampions() {
	champions, err := GetChampions("default")
	if err != nil {
		log.Panicln(err)
	}

	for _, champion := range champions {
		if champion.Name == "Nunu & Willump" {
			fmt.Println(champion.NameURL)
		}
	}

	// Output: Nunu
}

// Exemplifica o uso de League.Export.
func ExampleLeague_Export() {
	// Instanciando League.
	lol, err := New("default")
	if err != nil {
		log.Panicln(err)
	}

	// Traduzindo os dados para JSON.
	leagueData, err := lol.Export(4)
	if err != nil {
		log.Panicln(err)
	}

	// Exibindo o tipo da variável leagueData
	fmt.Printf("LeagueData: %T", leagueData)

	// Output: LeagueData: []uint8
}

// Como salvar os dados de League.
func ExampleLeague_Save() {
	lol, err := New("default")
	if err != nil {
		log.Panicln(err)
	}

	err = lol.Save("result.json", 4)
	if err != nil {
		log.Panicln(err)
	}

	fileInfo, err := os.Stat("result.json")
	if err != nil {
		log.Panicln(err)
	}

	exportedFilename := fileInfo.Name()
	fmt.Println(exportedFilename)

	// Output: result.json
}

// Exemplifica o uso de League.GetChampionsNames.
func ExampleLeague_GetChampionsNames() {
	lol, err := New("default")
	if err != nil {
		log.Panicln(err)
	}

	championsNames := lol.GetChampionsNames()
	for index, championName := range championsNames {
		if index == 5 {
			fmt.Println(index, championName)
		}
	}

	// Output: 5 Urgot
}
