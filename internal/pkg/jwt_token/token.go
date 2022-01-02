package jwt_token

type jwtToken struct {
	header map[string]string
	payload map[string]interface{}
	signature string
}


