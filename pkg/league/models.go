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
