package go_league

import (
	"errors"
	"fmt"
	"math"
)

type GameStats struct {
	ActivePlayer ActivePlayer `json:"activePlayer"`
	AllPlayers   []GamePlayer `json:"allPlayers"`
	Events       Events       `json:"events"`
	GameData     GameData     `json:"gameData"`
}

func (g *GameStats) SecondsElapsed() int {
	return int(math.Round(g.GameData.GameTime))
}

func (g *GameStats) MinutesElapsed() int {
	return int(math.Round(g.GameData.GameTime) / 60)
}

type Ability struct {
	AbilityLevel   int    `json:"abilityLevel"`
	DisplayName    string `json:"displayName"`
	ID             string `json:"id"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}

type Abilities struct {
	E       Ability `json:"E"`
	Q       Ability `json:"Q"`
	R       Ability `json:"R"`
	W       Ability `json:"W"`
	Passive Passive `json:"Passive"`
}

type Passive struct {
	DisplayName    string `json:"displayName"`
	ID             string `json:"id"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}

type ChampionStats struct {
	AbilityPower                 float64 `json:"abilityPower"`
	Armor                        float64 `json:"armor"`
	ArmorPenetrationFlat         float64 `json:"armorPenetrationFlat"`
	ArmorPenetrationPercent      float64 `json:"armorPenetrationPercent"`
	AttackDamage                 float64 `json:"attackDamage"`
	AttackRange                  float64 `json:"attackRange"`
	AttackSpeed                  float64 `json:"attackSpeed"`
	BonusArmorPenetrationPercent float64 `json:"bonusArmorPenetrationPercent"`
	BonusMagicPenetrationPercent float64 `json:"bonusMagicPenetrationPercent"`
	CooldownReduction            float64 `json:"cooldownReduction"`
	CritChance                   float64 `json:"critChance"`
	CritDamage                   float64 `json:"critDamage"`
	CurrentHealth                float64 `json:"currentHealth"`
	HealthRegenRate              float64 `json:"healthRegenRate"`
	LifeSteal                    float64 `json:"lifeSteal"`
	MagicLethality               float64 `json:"magicLethality"`
	MagicPenetrationFlat         float64 `json:"magicPenetrationFlat"`
	MagicPenetrationPercent      float64 `json:"magicPenetrationPercent"`
	MagicResist                  float64 `json:"magicResist"`
	MaxHealth                    float64 `json:"maxHealth"`
	MoveSpeed                    float64 `json:"moveSpeed"`
	PhysicalLethality            float64 `json:"physicalLethality"`
	ResourceMax                  float64 `json:"resourceMax"`
	ResourceRegenRate            float64 `json:"resourceRegenRate"`
	ResourceType                 string  `json:"resourceType"`
	ResourceValue                float64 `json:"resourceValue"`
	SpellVamp                    float64 `json:"spellVamp"`
	Tenacity                     float64 `json:"tenacity"`
}

type Rune struct {
	ID             int    `json:"id"`
	DisplayName    string `json:"displayName"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}

type StatRune struct {
	ID             int    `json:"id"`
	RawDescription string `json:"rawDescription"`
}

type FullRunes struct {
	GeneralRunes      []Rune     `json:"generalRunes"`
	Keystone          Rune       `json:"keystone"`
	PrimaryRuneTree   Rune       `json:"primaryRuneTree"`
	SecondaryRuneTree Rune       `json:"secondaryRuneTree"`
	StatRunes         []StatRune `json:"statRunes"`
}

type ActivePlayer struct {
	Abilities     Abilities     `json:"abilities"`
	ChampionStats ChampionStats `json:"championStats"`
	CurrentGold   float64       `json:"currentGold"`
	FullRunes     FullRunes     `json:"fullRunes"`
	Level         int           `json:"level"`
	SummonerName  SummonerName  `json:"summonerName"`
}

type Runes struct {
	Keystone          Rune `json:"keystone"`
	PrimaryRuneTree   Rune `json:"primaryRuneTree"`
	SecondaryRuneTree Rune `json:"secondaryRuneTree"`
}

type Scores struct {
	Assists    int     `json:"assists"`
	CreepScore int     `json:"creepScore"`
	Deaths     int     `json:"deaths"`
	Kills      int     `json:"kills"`
	WardScore  float64 `json:"wardScore"`
}

func (s Scores) AverageVisionScore(gameMinutes int) float64 {
	return s.WardScore / float64(gameMinutes)
}

func (s Scores) AverageCreepScore(gameMinutes int) int {
	return s.CreepScore / gameMinutes
}

type SummonerSpell struct {
	DisplayName    string `json:"displayName"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}

type SummonerSpells struct {
	SummonerSpellOne SummonerSpell `json:"summonerSpellOne"`
	SummonerSpellTwo SummonerSpell `json:"summonerSpellTwo"`
}

type Item struct {
	CanUse         bool   `json:"canUse"`
	Consumable     bool   `json:"consumable"`
	Count          int    `json:"count"`
	DisplayName    string `json:"displayName"`
	ItemID         int    `json:"itemID"`
	Price          int    `json:"price"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
	Slot           int    `json:"slot"`
}

type SummonerName string

const SummonerNameUnknown SummonerName = ""

type GamePlayer struct {
	ChampionName    ChampionName   `json:"championName"`
	IsBot           bool           `json:"isBot"`
	IsDead          bool           `json:"isDead"`
	Items           []Item         `json:"items"`
	Level           int            `json:"level"`
	Position        Position       `json:"position"`
	RawChampionName string         `json:"rawChampionName"`
	RespawnTimer    float64        `json:"respawnTimer"`
	Runes           Runes          `json:"runes"`
	Scores          Scores         `json:"scores"`
	SkinID          int            `json:"skinID"`
	SummonerName    SummonerName   `json:"summonerName"`
	SummonerSpells  SummonerSpells `json:"summonerSpells"`
	Team            TeamName       `json:"team"`
}

func (g *GameStats) Valid() error {
	if g == nil {
		return errors.New("GameStats is nil")
	}

	return nil
}

func (g *GameStats) GetActivePlayerSummonerName() SummonerName {
	return g.ActivePlayer.SummonerName
}

func (g *GameStats) GetOpposingTeam() ([]GamePlayer, error) {
	activeTeam, err := g.GetPlayersTeamName(g.ActivePlayer.SummonerName)
	if err != nil {
		return nil, err
	}

	opposingTeam := make([]GamePlayer, 0)
	for _, player := range g.AllPlayers {
		if player.Team != activeTeam {
			opposingTeam = append(opposingTeam, player)
		}
	}

	return opposingTeam, nil
}

func (g *GameStats) GetPlayersTeamName(summonerName SummonerName) (TeamName, error) {
	for _, player := range g.AllPlayers {
		if player.SummonerName == summonerName {
			return player.Team, nil
		}
	}

	return "", fmt.Errorf("could not find %s's team name", summonerName)
}

func (g *GameStats) IsSameTeam(player1, player2 SummonerName) bool {
	for _, p1 := range g.AllPlayers {
		if p1.SummonerName == player1 {
			for _, p2 := range g.AllPlayers {
				if p2.SummonerName == player2 {
					return p1.Team == p2.Team
				}
			}
		}
	}

	return false
}

func (g *GameStats) setSoulEvent() error {
	var chaosKills, orderKills int

	for _, event := range g.Events.IngameEvents {
		if event.EventName != EventDragonKill {
			continue
		}

		playersTeams, err := g.GetPlayersTeamName(event.KillerName)
		if err != nil {
			return err
		}

		if playersTeams == TeamNameChaos {
			chaosKills++
		} else if playersTeams == TeamNameOrder {
			orderKills++
		}

		if chaosKills == 4 || orderKills == 4 {
			event.EventName = EventSoulKill
		}
	}

	return nil
}

func (g *GameStats) SetEvents() error {
	g.setEventTeams()

	return g.setSoulEvent()
}

//sets the KillerTeam field for all events which have a KillerName.
func (g *GameStats) setEventTeams() {
	for _, event := range g.Events.IngameEvents {
		if event.KillerName == SummonerNameUnknown {
			continue
		}

		for _, player := range g.AllPlayers {
			if player.SummonerName == event.KillerName {
				event.KillerTeam = player.Team
				break
			}
		}
	}
}
