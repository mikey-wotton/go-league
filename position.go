package go_league

import "fmt"

// Position represents the different lanes people can be in within regular games, not set during ARAM/URF/TFT.
type Position string

const (
	PositionUnknown Position = ""
	PositionTop     Position = "TOP"
	PositionJungle  Position = "JUNGLE"
	PositionMiddle  Position = "MIDDLE"
	PositionBottom  Position = "BOTTOM"
	PositionSupport Position = "UTILITY"
)

var positions = []Position{PositionTop, PositionJungle, PositionMiddle, PositionBottom, PositionSupport}

func (p Position) Valid() error {
	for _, position := range positions {
		if position == p {
			return nil
		}
	}

	return fmt.Errorf("unknown position '%s' provided", p)
}
