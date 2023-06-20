// Package league é um pacote que te
// manipular os dados de todos os campeões,
// skins e chromas do League of Legends.
package league

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Unit é uma estrutura que representa uma skin
// ou um campeão do League of Legends.
type Unit struct {
	IsBase         bool // Se a Unit é um campeão (ou skin).
	IsLegacy       bool
	ID             int
	Name           string
	Description    string
	Rarity         string
	SplashPath     string
	LoadScreenPath string
	Chromas        []Chroma
}

// Champion é uma estrutura que representa os
// campeões do League of Legends.
//
// Nem todos os campeões possuem Chromas.
type Champion struct {
	ID          int
	Name        string   // Nome do campeão.
	NameURL     string   // Nome de URL do campeão.
	Description string   // Descrição do campeão.
	Rarity      string   // Por exemplo: kMythic.
	IsLegacy    bool     // Se é um campeão legado.
	Chromas     []Chroma // Chromas do campeão.
	Skins       []Skin   // Skins do campeão.
}

// Skin é uma estrutura que representa as
// skins do League of Legends.
//
// Todos os campeões possuem uma skin.
// Nem todas as skins possuem Chromas.
type Skin struct {
	ID          int
	ChampionId  int      // ID do campeão.
	Name        string   // Nome da skin.
	Description string   // Descrição da skin.
	Rarity      string   // Por exemplo: kMythic.
	IsLegacy    bool     // Se é uma skin legado.
	Chromas     []Chroma // Chromas da skin, opcional.
}

// Chroma é uma skin personalizada que
// pode existir em campeões e skins.
type Chroma struct {
	ID     int      // ID do Chroma.
	Name   string   // Nome do Chroma.
	Colors []string // Cores do Chroma.
}

const baseURL = "https://raw.communitydragon.org" // Dados crus da Riot.

func getUnits(region string) (units map[string]Unit, err error) {
	// Existem outros endpoints que podem fazer o mesmíssimo trabalho desta.
	endpoint := fmt.Sprintf("/latest/plugins/rcp-be-lol-game-data/global/%s/v1/skins.json", region)

	response, err := http.Get(baseURL + endpoint)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("requisitar os campeões: %v", err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&units); err != nil {
		return nil, fmt.Errorf("estruturação do JSON: %v", err)
	}
	return units, nil
}

// GetChampions retorna uma lista de campeões com suas skins e chromas.
//
// O usuário pode escolher o idioma da resposta, escolhendo entre as regiões
// conhecidas, como "default" (inglês) ou "pt_br" (português brasileiro).
func GetChampions(region string) (champions []Champion, err error) {
	units, err := getUnits(region)
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
			Skins:       nil,
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
