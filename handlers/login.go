package handlers

import (
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
)

var loginOauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func Login(w http.ResponseWriter, r *http.Request) {
	callback := "http://" + r.host + "/callback"
	tempCred, err := loginOauthClient.RequestTemporaryCredentials(nil,callback,nil)

	//todo
}
