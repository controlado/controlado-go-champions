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
	if _, err := New("default"); err != nil {
		t.Errorf("Not working: %v", err)
	}

	if _, err := New("invalid"); err == nil {
		t.Error("Expected: error Got: nil")
	}
}

// Exemplifica o uso de New.
func ExampleNew() {
	lol, err := New("default")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Instância: %T\n", lol)
	fmt.Printf("Champions: %T", lol.Champions)

	// Output:
	// Instância: league.League
	// Champions: []league.Champion
}

// Exemplifica o uso de GetChampions.
func ExampleGetChampions() {
	champions, err := GetChampions("default")
	if err != nil {
		log.Panicln("Não foi possível requisitar os campeões")
	}

	for _, champion := range champions {
		if champion.ID == 4 {
			fmt.Println(champion.Name)
		}
	}

	// Output: Twisted Fate
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
	namesLength := len(championsNames)
	fmt.Println(namesLength)

	// Output: 163
}
