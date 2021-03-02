package go_league

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameStats_GetOpposingTeam(t *testing.T) {
	tests := map[string]struct{
		stats *GameStats

		expTeam []GamePlayer
		expErr	error
	}{
		"Simplest": {
			stats: statsHelper(),
			expTeam: []GamePlayer{
				{
					Team: TeamNameChaos,
					SummonerName: "bob",
				},
				{
					Team: TeamNameChaos,
					SummonerName: "steve",
				},
				{
					Team: TeamNameChaos,
					SummonerName: "tim",
				},
			},
			expErr: nil,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			opposingTeam, err := test.stats.GetOpposingTeam()

			assert.Equal(t, test.expErr, err)
			assert.Equal(t, test.expTeam, opposingTeam)
		})
	}
}

func statsHelper() *GameStats {
	return &GameStats{
		ActivePlayer: ActivePlayer{
			SummonerName:  "mike",
		},
		AllPlayers: []GamePlayer{
			{
				Team: TeamNameChaos,
				SummonerName: "bob",
			},
			{
				Team: TeamNameChaos,
				SummonerName: "steve",
			},
			{
				Team: TeamNameChaos,
				SummonerName: "tim",
			},
			{
				Team: TeamNameOrder,
				SummonerName: "eric",
			},
			{
				Team: TeamNameOrder,
				SummonerName: "john",
			},
			{
				Team: TeamNameOrder,
				SummonerName: "mike",
			},
		},
	}
}