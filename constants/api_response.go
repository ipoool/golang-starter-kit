package constants

const (
	ApiSuccessMessage      = 1000
	ApiFailedMessage       = 2000
	ApiAuthorizationFailed = 4001
)

var ApiReseponse = map[int]string{
	ApiSuccessMessage:      "SUCCESS",
	ApiFailedMessage:       "FAILED",
	ApiAuthorizationFailed: "AUTHORIZATION FAILED",
}
