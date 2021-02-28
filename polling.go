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
		<-time.Tick(pollingWaitTime)
		_, err := GetLiveData()
		if err == nil {
			break
		}
	}
}

func PollUntilGameEndChannel() (chan *GameStats, chan error) {
	statsChan := make(chan *GameStats, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(statsChan)
		defer close(errChan)

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

			<-time.Tick(15 * time.Second)
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
