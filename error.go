package go_league

import "errors"

const (
	apiDownErrMessage = `Get "` + allGameDataURL + `": dial tcp 127.0.0.1:2999: connectex: No connection could be made because the target machine actively refused it.`
)

var (
	//The error we receive when the league api goes down without warning
	errConnRefused = errors.New(apiDownErrMessage)
)

type Error struct {
	ErrorCode  string `json:"errorCode"`
	HttpStatus int    `json:"httpStatus"`
	Message    string `json:"message"`
}

func IsServerRefusedErr(err error) bool {
	return err.Error() == errConnRefused.Error()
}
