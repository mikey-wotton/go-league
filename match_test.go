package go_league

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch_AllPlayers(t *testing.T) {
	kirky := &Player{
		TeamName:     TeamNameOrder,
		SummonerName: "kirky",
		ChampionName: ChampionNameBlitzcrank,
		FullScores: FullScores{
			DragonKills: 1,
			BaronKills:  2,
			Scores: Scores{
				Assists:    3,
				CreepScore: 4,
				Deaths:     5,
				Kills:      6,
				WardScore:  0,
			},
		},
	}
	mike := &Player{
		TeamName:     TeamNameChaos,
		SummonerName: "mikehpoo",
		ChampionName: ChampionNameShyvana,
		FullScores: FullScores{
			DragonKills: 0,
			BaronKills:  2,
			Scores: Scores{
				Assists:    3,
				CreepScore: 4,
				Deaths:     5,
				Kills:      6,
				WardScore:  7,
			},
		},
	}
	toBeRemoved := &Player{
		TeamName:     TeamNameOrder,
		SummonerName: "toBeRemoved",
	}

	match := Match{
		Order: []*Player{kirky, toBeRemoved},
		Chaos: []*Player{mike},
	}

	all := match.AllPlayers("mikehpoo", "kirky")

	assert.Equal(t, []*Player{kirky, mike}, all)
}

func TestGameStats_IsSameTeam(t *testing.T) {
	tests := map[string]struct {
		p1          SummonerName
		p2          SummonerName
		expSameTeam bool
	}{
		"Same team is true": {
			p1:          "kirky",
			p2:          "mikey",
			expSameTeam: true,
		},
		"Different team is false": {
			p1:          "kirky",
			p2:          "wezznco",
			expSameTeam: false,
		},
		"Player in no team is false": {
			p1:          "UNKNOWN",
			p2:          "mikey",
			expSameTeam: false,
		},
	}

	stats := &GameStats{
		AllPlayers: []GamePlayer{
			{
				SummonerName: "mikey",
				Team:         TeamNameChaos,
			},
			{
				SummonerName: "kirky",
				Team:         TeamNameChaos,
			},
			{
				SummonerName: "wezznco",
				Team:         TeamNameOrder,
			},
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			isSameTeam := stats.IsSameTeam(test.p1, test.p2)
			assert.Equal(t, test.expSameTeam, isSameTeam)
		})
	}
}

func TestMatch_ActivePlayerTeam(t *testing.T) {
	tests := map[string]struct {
		filterList []SummonerName
		match Match

		expTeam []*Player
	}{
		"ActivePlayer team is Order, returns entire team": {
			match: Match{
				ActivePlayerName: "mike",
				ActiveTeam:       TeamNameOrder,
				Order:            []*Player{
					{
						TeamName:     TeamNameOrder,
						SummonerName: "bob",
						ChampionName: ChampionNameEkko,
					},
					{
						TeamName:     TeamNameOrder,
						SummonerName: "mike",
						ChampionName: ChampionNameYorick,
					},
				},
				Chaos:            []*Player{
					{
						TeamName:     TeamNameChaos,
						SummonerName: "steve",
						ChampionName: ChampionNameYasuo,
					},
				},
			},
			expTeam: []*Player{
				{
					TeamName:     TeamNameOrder,
					SummonerName: "bob",
					ChampionName: ChampionNameEkko,
				},
				{
					TeamName:     TeamNameOrder,
					SummonerName: "mike",
					ChampionName: ChampionNameYorick,
				},
			},
		},
		"ActivePlayer's team is chaos, returns only filtered players": {
			filterList: []SummonerName{"mike"},
			match: Match{
				ActivePlayerName: "mike",
				ActiveTeam:       TeamNameChaos,
				Order:            []*Player{
					{
						TeamName:     TeamNameOrder,
						SummonerName: "bob",
						ChampionName: ChampionNameEkko,
					},
					{
						TeamName:     TeamNameOrder,
						SummonerName: "jerry",
						ChampionName: ChampionNameEkko,
					},
				},
				Chaos:            []*Player{
					{
						TeamName:     TeamNameChaos,
						SummonerName: "steve",
						ChampionName: ChampionNameYasuo,
					},
					{
						TeamName:     TeamNameChaos,
						SummonerName: "mike",
						ChampionName: ChampionNameYasuo,
					},
				},
			},
			expTeam: []*Player{
				{
					TeamName:     TeamNameChaos,
					SummonerName: "mike",
					ChampionName: ChampionNameYasuo,
				},
			},
		},
		"No team returns nil": {
			match: Match{
				ActivePlayerName: "mike",
				Order:            []*Player{
					{
						TeamName:     TeamNameOrder,
						SummonerName: "bob",
						ChampionName: ChampionNameEkko,
					},
					{
						TeamName:     TeamNameOrder,
						SummonerName: "jerry",
						ChampionName: ChampionNameEkko,
					},
				},
				Chaos:            []*Player{
					{
						TeamName:     TeamNameChaos,
						SummonerName: "steve",
						ChampionName: ChampionNameYasuo,
					},
					{
						TeamName:     TeamNameChaos,
						SummonerName: "mike",
						ChampionName: ChampionNameYasuo,
					},
				},
			},
		},

	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {

			team := test.match.ActivePlayerTeam(test.filterList...)

			assert.Equal(t, test.expTeam, team)
		})
	}
}
