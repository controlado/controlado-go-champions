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
	resultFilename := "result.json"

	// Instanciando League.
	lol, err := New("default")
	if err != nil {
		log.Panicln(err)
	}

	// Traduzindo os dados para JSON.
	championsJSON, err := lol.Export(4)
	if err != nil {
		log.Panicln(err)
	}

	// Salvando o JSON para um arquivo local.
	if os.WriteFile(resultFilename, championsJSON, 0644) != nil {
		log.Fatal("Não foi possível salvar os dados no arquivo")
	}

	// Verificando se o arquivo existe.
	fileInfo, err := os.Stat(resultFilename)
	if err != nil {
		log.Panicln(err)
	}

	resultFilename = fileInfo.Name()
	fmt.Println(resultFilename)

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
