package fixtures

func UnauthorizedError() string {
	unauthorized := `{
		"message": "unauthenticated user"
	}`
	return unauthorized
}
