package go_league

type EventName string

const (
	EventSoulKill EventName = "SoulKill"

	EventDragonKill   EventName = "DragonKill"
	EventBaronKill    EventName = "BaronKill"
	EventHeraldKill   EventName = "HeraldKill"
	EventInhibKill    EventName = "InhibKilled"
	EventTurretKill   EventName = "TurretKilled"
	EventChampionKill EventName = "ChampionKill"
	EventMultiKill    EventName = "Multikill"
	EventAce          EventName = "Ace"
	EventGameStart    EventName = "GameStart"
	EventMinionSpawn  EventName = "MinionsSpawning"
	EventGameEnd      EventName = "GameEnd"
)

type DragonType string

const (
	DragonTypeEarth DragonType = "Earth"
	DragonTypeAir   DragonType = "Air"
	DragonTypeFire  DragonType = "Fire"
	DragonTypeWater DragonType = "Water"
	DragonTypeElder DragonType = "Elder"
)

func (d DragonType) ToDrake() string {
	switch d {
	case DragonTypeAir:
		return "Cloud"
	case DragonTypeEarth:
		return "Mountain"
	case DragonTypeFire:
		return "Infernal"
	case DragonTypeWater:
		return "Ocean"
	default:
		return "UNKNOWN DRAKE TYPE"
	}
}

type Result string

const (
	ResultWin  Result = "Win"
	ResultLoss Result = "Lose"
)

type EventID int

type Events struct {
	IngameEvents []IngameEvent `json:"Events"`
}

type IngameEvent struct {
	EventID    EventID        `json:"EventID"`
	EventName  EventName      `json:"EventName"`
	EventTime  float64        `json:"EventTime"`
	Stolen     Stolen         `json:"Stolen,omitempty"`
	KillStreak int            `json:"KillStreak,omitempty"`
	VictimName SummonerName   `json:"VictimName,omitempty"`
	DragonType DragonType     `json:"DragonType,omitempty"`
	Assisters  []SummonerName `json:"Assisters,omitempty"`
	Acer       SummonerName   `json:"Acer,omitempty"`
	AcingTeam  TeamName       `json:"AcingTeam,omitempty"`
	Result     Result         `json:"Result,omitempty"`
	KillerName SummonerName   `json:"KillerName,omitempty"`

	KillerTeam TeamName `json:"KillerTeam,omitempty"`
}

// IsStealable returns true if an event can have the "Stolen" flag.
func (i IngameEvent) IsStealable() bool {
	switch i.EventName {
	case EventDragonKill:
		return true
	case EventHeraldKill:
		return true
	case EventBaronKill:
		return true
	default:
		return false
	}
}

func (i IngameEvent) IsStolen() bool {
	return i.Stolen == StolenTrue
}
