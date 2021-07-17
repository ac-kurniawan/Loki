package apiControllerV1

type SetAuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}