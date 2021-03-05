package go_league

import (
	"fmt"
	"log"
	"time"
)

const (
	pollingWaitTime        = 1 * time.Second
	numberOfGameEndRetries = 5
)

func PollUntilGameStart() {
	for {
		ticker := time.NewTicker(pollingWaitTime)
		<-ticker.C

		game, err := GetLiveData()
		if err == nil {
			//sometimes the api's can come up but not be ready to respond with data so default values are set
			if game.GameData.GameMode != GameModeUnknown {
				break
			}
		}
	}
}

func PollUntilGameEndChannel() (chan *GameStats, chan error) {
	statsChan := make(chan *GameStats, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(statsChan)
		defer close(errChan)
		ticker := time.NewTicker(time.Second * 15)

		gameHasEnded := GameEndErrCounter()

		for {
			gameData, err := GetLiveData()
			if err != nil {
				gameOngoing, e := gameHasEnded(err)
				if e != nil {
					errChan <- e
					return
				}

				if !gameOngoing {
					break
				}
			}

			if gameData != nil {
				statsChan <- gameData
			}

			<-ticker.C
		}

	}()

	return statsChan, errChan
}

func GameEndErrCounter() func(error) (bool, error) {
	count := 0
	return func(err error) (bool, error) {
		if IsServerRefusedErr(err) {
			count++
			log.Println(fmt.Sprintf("Game ended %d/%d", count, numberOfGameEndRetries))
			if count == numberOfGameEndRetries {
				return true, nil
			}
			return false, nil
		}

		//game has ended by some unknown error
		return true, err
	}
}
