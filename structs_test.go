package testpage

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	var user User
	GetSession().Get(&user)
	b := Success(user)
	t.Log(string(b))
}

func TestPageSuccess(t *testing.T) {
	session := GetSession()
	ok := Paginate(1, 10, session, nil)
	b := Success(ok)

	t.Log(string(b))
}
