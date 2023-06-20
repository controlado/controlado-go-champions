package league

import (
	"fmt"
	"log"
	"testing"
)

// Representa a contagem de unidades.
type UnitCounter struct {
	Champions int
	Skins     int
}

// Esse teste verifica se a quantidade de
// campeões na requisição crua é a mesma
// da requisição final: GetChampions.
func TestGetChampions(t *testing.T) {
	champions, err := GetChampions("default")
	if err != nil {
		t.Error("Não foi possível requisitar os campeões")
	}

	units, err := getUnits("default")
	if err != nil {
		t.Error("Não foi possível requisitar as unidades")
	}

	var expectedLength = len(champions)
	var unitsLength = countUnits(units)

	if expectedLength != unitsLength.Champions {
		t.Error("Expected:", unitsLength, "Got:", expectedLength)
	}
}

// Esse teste verifica a resposta para o usuário
// caso a região argumentada seja inválida.
func TestGetChampionsInvalidRegion(t *testing.T) {
	_, err := GetChampions("something_invalid")
	if err == nil {
		t.Error(err)
	}
}

// Exemplifica o uso de GetChampions.
func ExampleGetChampions() {
	champions, err := GetChampions("default")
	if err != nil {
		log.Println("Não foi possível requisitar os campeões")
		panic(err)
	}

	for _, champion := range champions {
		if champion.ID == "4" {
			fmt.Println(champion.Name)
		}
	}

	// Output: Twisted Fate
}

// Conta quantos campeões e skins existem nas unidades.
func countUnits(units map[string]Unit) (counter UnitCounter) {
	for _, unit := range units {
		if unit.IsBase {
			counter.Champions++
		} else {
			counter.Skins++
		}
	}

	return counter
}
