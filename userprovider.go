package testpage

func UserApiProvider() (b string) {
	session := GetSession()
	ok := Paginate(1, 10, session, nil)
	b = Success(ok)
	return string(b)
}
