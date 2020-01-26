package api

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"server/templates"
)

// sessionValue stores a random string.
var sessionValue string

//store the cookie store which is going to store session data in the cookie
var store = sessions.NewCookieStore([]byte("secret-password"))
var session *sessions.Session

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
	"abcdefghijklmnopqrstuvwxyzåäö" +
	"0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
	sessionValue = generateValue(16)
}

//generateValue generates a random string
func generateValue(length uint) string {
	var b strings.Builder
	for i := uint(0); i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

//IsLoggedIn will check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, "session")
	if err == nil && (session.Values["loggedin"] == sessionValue) {
		return true
	}
	return false
}

//LogoutFunc handles "/logout"
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	if IsLoggedIn(r) {
		session, err := store.Get(r, "session")
		if err == nil && (session.Values["loggedin"] == sessionValue) {
			session.Values["loggedin"] = nil
		}
	}
	return
}

// LoginFunc handles "/login"
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	randInt := rand.Intn(1000) + 1000
	time.Sleep(time.Duration(randInt) * time.Millisecond)

	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	switch r.Method {
	case "GET":
		templates.LoginTemplate.Execute(w, nil)
	case "POST":
		log.Print("Inside POST")
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if (username != "") && validUser(username, password) {
			session, _ := store.Get(r, "session")
			session.Values["loggedin"] = sessionValue
			session.Values["username"] = username
			session.Save(r, w)
			log.Print("user ", username, " is authenticated")
			http.Redirect(w, r, "/", 302)
			return
		}
		log.Print("Invalid user " + username)
		templates.LoginTemplate.Execute(w, nil)
	default:
		http.Redirect(w, r, "/login/", http.StatusUnauthorized)
	}
}

func validUser(username, password string) bool {
	combination := username + "{" + password + "}256"
	hasher := sha256.New()
	hasher.Write([]byte(combination))
	if hex.EncodeToString(hasher.Sum(nil)) == "d77c2632fbd17b3341f64462d6d22b573871d62eb13b457f74e42c21af2bcee9" {
		return true
	}
	return false
}
