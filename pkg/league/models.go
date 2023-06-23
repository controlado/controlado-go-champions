package league

// League é a estrutura central do pacote.
//
// Possui todos os campeões, skins e cromas,
// além de fornecer métodos que auxiliam na
// manipulação dos dados gerados.
type League struct {
	Champions []Champion
}

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
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`        // Nome do campeão.
	NameURL     string   `json:"nameUrl,omitempty"`     // Nome de URL do campeão.
	Description string   `json:"description,omitempty"` // Descrição do campeão.
	Rarity      string   `json:"rarity,omitempty"`      // Por exemplo: kMythic.
	IsLegacy    bool     `json:"isLegacy,omitempty"`    // Se é um campeão legado.
	Chromas     []Chroma `json:"chromas,omitempty"`     // Chromas do campeão.
	Skins       []Skin   `json:"skins,omitempty"`       // Skins do campeão.
}

// Skin é uma estrutura que representa as
// skins do League of Legends.
//
// Todos os campeões possuem uma skin.
// Nem todas as skins possuem Chromas.
type Skin struct {
	ID          int      `json:"id,omitempty"`
	ChampionId  int      `json:"championId,omitempty"`  // ID do campeão.
	Name        string   `json:"name,omitempty"`        // Nome da skin.
	Description string   `json:"description,omitempty"` // Descrição da skin.
	Rarity      string   `json:"rarity,omitempty"`      // Por exemplo: kMythic.
	IsLegacy    bool     `json:"isLegacy,omitempty"`    // Se é uma skin legado.
	Chromas     []Chroma `json:"chromas,omitempty"`     // Chromas da skin, opcional.
}

// Chroma é uma skin personalizada que
// pode existir em campeões e skins.
type Chroma struct {
	ID     int      `json:"id,omitempty"`     // ID do Chroma.
	Name   string   `json:"name,omitempty"`   // Nome do Chroma.
	Colors []string `json:"colors,omitempty"` // Cores do Chroma.
}
