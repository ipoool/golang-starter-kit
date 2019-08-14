package constants

const (
	ApiSuccessMessage      = 1000
	ApiAuthorizationFailed = 4001
)

var ApiReseponse = map[int]string{
	ApiSuccessMessage:      "SUCCESS",
	ApiAuthorizationFailed: "AUTHORIZATION FAILED",
}
