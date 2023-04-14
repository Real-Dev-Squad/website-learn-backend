package fixtures

func SendCookie() string {
	cookie := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJGc0VmZHhyRHpxWUJCV1phakM4ciIsImlhdCI6MTY4MDMwOTU0OSwiZXhwIjoxNjgyOTAxNTQ5fQ.YV4M44fvlg3g7S_skKQkY5-rJ5qESFVlunQyTF0K3KfG9I0xnWrg2WIx4M-SUYszwqkFq5WPln-qKd6O1UXWUJ9uUtx0ybSy0dg5NyWPuB3A4E0npNiBRliin3bCKgPXL2AfYhygX2ovEhuW14p7I74jgeVLz5SjwI6_gLKYEiY"

	return cookie
}

func Unauthorized() string {
	unauthorized := `{
		"message": "unauthenticated user"
	}`
	return unauthorized
}

func Authorized() string {
	authorized := `{
		"message": "Welcome to the authenticated route",
		"userId": "FsEfdxrDzqYBBWZajC8r"
	}`

	return authorized
}
