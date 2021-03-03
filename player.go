package go_league

import (
	"fmt"
)

type Player struct {
	TeamName     TeamName
	SummonerName SummonerName
	ChampionName ChampionName

	FullScores FullScores
}

type FullScores struct {
	DragonKills int
	BaronKills  int
	Scores      Scores
}

func newPlayer(gamer GamePlayer, drakeKills, baronKills int) *Player {
	return &Player{
		TeamName:     gamer.Team,
		SummonerName: gamer.SummonerName,
		ChampionName: gamer.ChampionName,
		FullScores: FullScores{
			DragonKills: drakeKills,
			BaronKills:  baronKills,
			Scores:      gamer.Scores,
		},
	}
}

func newMatch(activePlayer SummonerName, activePlayerTeamName TeamName) *Match {
	return &Match{
		ActivePlayerName: activePlayer,
		ActiveTeam:       activePlayerTeamName,
		Order:            []*Player{},
		Chaos:            []*Player{},
	}
}

func GetMatch(stats GameStats) (*Match, error) {
	activePlayerTeam, err := stats.GetPlayersTeamName(stats.ActivePlayer.SummonerName)
	if err != nil {
		return nil, err
	}

	match := newMatch(stats.ActivePlayer.SummonerName, activePlayerTeam)

	drakeKills := getDrakeKills(stats)

	baronKills := getBaronKills(stats)

	for _, player := range stats.AllPlayers {
		switch player.Team {
		case TeamNameChaos:
			match.Chaos = append(match.Chaos, newPlayer(player, drakeKills, baronKills))
		case TeamNameOrder:
			match.Order = append(match.Order, newPlayer(player, drakeKills, baronKills))
		default:
			return nil, fmt.Errorf("unknown team name provided: %s", player.Team)
		}
	}

	return match, nil
}

func getDrakeKills(stats GameStats) int {
	count := 0
	for _, event := range stats.Events.IngameEvents {
		if event.EventName == EventDragonKill {
			count++
		}
	}

	return count
}

func getBaronKills(stats GameStats) int {
	count := 0
	for _, event := range stats.Events.IngameEvents {
		if event.EventName == EventBaronKill {
			count++
		}
	}

	return count
}
