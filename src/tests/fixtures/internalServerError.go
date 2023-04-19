package fixtures

func InternalServerError() string {
	serverError := `{
		"message":"An internal server error occurred"
	}`
	return serverError
}
