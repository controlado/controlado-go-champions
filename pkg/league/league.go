// Package league é um pacote que te
// manipular os dados de todos os campeões,
// skins e chromas do League of Legends.
package league

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/controlado/league-champions-go/internal/requests"
)

// GetChampions retorna uma lista de campeões com suas skins e chromas.
//
// É recomendado que New seja utilizado para acessar métodos de manipulação.
//
// O usuário pode escolher o idioma da resposta, escolhendo entre as regiões
// conhecidas, como "default" (inglês) ou "pt_br" (português brasileiro).
func GetChampions(region string) (champions []Champion, err error) {
	units, err := requests.GetUnits(region)
	if err != nil {
		return nil, err
	}

	for _, unit := range units {
		if !unit.IsBase {
			continue
		}
		championNameUrl := strings.Split(unit.LoadScreenPath, "/")[5]
		championIdString := strings.Split(unit.SplashPath, "/")[5]
		championId, _ := strconv.Atoi(championIdString)
		champion := Champion{
			ID:          championId,
			Name:        unit.Name,
			NameURL:     championNameUrl,
			Description: unit.Description,
			Rarity:      unit.Rarity,
			IsLegacy:    unit.IsLegacy,
			Chromas:     unit.Chromas,
		}
		champions = append(champions, champion)
	}

	for _, unit := range units {
		if unit.IsBase {
			continue
		}
		championIdString := strings.Split(unit.SplashPath, "/")[5]
		championId, _ := strconv.Atoi(championIdString)
		skin := Skin{
			ID:          unit.ID,
			ChampionId:  championId,
			Name:        unit.Name,
			Description: unit.Description,
			Rarity:      unit.Rarity,
			IsLegacy:    unit.IsLegacy,
			Chromas:     unit.Chromas,
		}
		for index := range champions {
			champion := &champions[index]
			if champion.ID == championId {
				champion.Skins = append(champion.Skins, skin)
			}
		}
	}

	return champions, nil
}

// New instancia a estrutura central do pacote.
func New(region string) (lol League, err error) {
	lol.Champions, err = GetChampions(region)
	if err != nil {
		return lol, err
	}
	lol.sortChampions()
	lol.sortSkins()
	return lol, nil
}

// GetChampionsNames retorna uma lista apenas com os nomes dos campeões.
//
// Útil para usuários que querem passar os nomes em um checker.
func (lol League) GetChampionsNames() (result []string) {
	for _, champion := range lol.Champions {
		result = append(result, champion.NameURL)
	}
	return result
}

// Export faz um parse dos dados gerados para JSON.
func (lol League) Export(indent int) ([]byte, error) {
	indentString := strings.Repeat(" ", indent) // Indentação do arquivo.
	return json.MarshalIndent(lol.Champions, "", indentString)
}

// Save faz parse com Export e salva os dados em um arquivo.
func (lol League) Save(filename string, indent int) error {
	leagueData, err := lol.Export(indent)
	if err != nil {
		return err
	}
	// os.WriteFile tem o mesmo retorno de Save.
	return os.WriteFile(filename, leagueData, 0644)
}

// Ordena os campeões com base no ID.
func (lol League) sortChampions() {
	sort.SliceStable(lol.Champions, func(i, j int) bool {
		return lol.Champions[i].ID < lol.Champions[j].ID
	})
}

// Ordena os skins com base no ID.
func (lol League) sortSkins() {
	for _, champion := range lol.Champions {
		sort.SliceStable(champion.Skins, func(i, j int) bool {
			return champion.Skins[i].ID < champion.Skins[j].ID
		})
	}
}
