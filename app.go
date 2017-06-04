package testpage

import (
	"io"
	"net/http"
)

func RunHttpSvr() (err error) {
	http.HandleFunc("/user", UserApi)
	return http.ListenAndServe("localhost:8080", nil)
}

func UserApi(w http.ResponseWriter, r *http.Request) {
	userResponse := UserApiProvider()
	io.WriteString(w, userResponse)
}
