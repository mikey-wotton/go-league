package go_league

import "fmt"

type TeamName string

const (
	TeamNameUnknown TeamName = ""
	TeamNameChaos   TeamName = "CHAOS"
	TeamNameOrder   TeamName = "ORDER"
)

type Match struct {
	ActivePlayerName SummonerName
	ActiveTeam       TeamName
	Order            []*Player
	Chaos            []*Player
}

//AllPlayers returns all players of a match but can be provided a filter list which will only return the
//players within the filter list.
func (m Match) AllPlayers(filterList ...SummonerName) []*Player {
	all := append(m.Order, m.Chaos...)

	if len(filterList) >= 0 {
		all = filter(all, filterList)
	}

	return all
}

func (m Match) ActivePlayerTeam(filterList ...SummonerName) []*Player {
	var all []*Player
	switch m.ActiveTeam {
	case TeamNameOrder:
		all = make([]*Player, len(m.Order))
		copy(all, m.Order)
	case TeamNameChaos:
		all = make([]*Player, len(m.Chaos))
		copy(all, m.Chaos)
	default:
		return nil
	}

	if len(filterList) > 0 {
		all = filter(all, filterList)
	}

	return all
}

func (t TeamName) Valid() error {
	switch t {
	case TeamNameOrder:
		return nil
	case TeamNameChaos:
		return nil
	default:
		return fmt.Errorf("unknown TeamName %s", t)
	}
}

func (t TeamName) OpposingTeam() (TeamName, error) {
	switch t {
	case TeamNameChaos:
		return TeamNameOrder, nil
	case TeamNameOrder:
		return TeamNameChaos, nil
	default:
		return "", fmt.Errorf("teamName '%s' has no opposing team", t)
	}
}

func filter(all []*Player, filterList []SummonerName) []*Player {
	if len(filterList) <= 0 {
		return all
	}

	filtered := make([]*Player, 0, len(all))

	for _, p := range all {
		var allowed bool

		for _, f := range filterList {
			if p.SummonerName == f {
				allowed = true
				break
			}
		}

		if allowed {
			filtered = append(filtered, p)
		}
	}

	return filtered
}
