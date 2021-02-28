package go_league

type GameMode string

const (
	GameModeUnknown GameMode = ""
	GameModeTFT     GameMode = "TFT"
	GameModeClassic GameMode = "CLASSIC"
)

type GameData struct {
	GameMode   GameMode `json:"gameMode"`
	GameTime   float64  `json:"gameTime"`
	MapName    string   `json:"mapName"`
	MapNumber  int      `json:"mapNumber"`
	MapTerrain string   `json:"mapTerrain"`
}
