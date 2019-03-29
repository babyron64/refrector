// Package auth and its subpackages provide OpenID connect authentication
// and authorization utilities.
package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type provider struct {
	name            string
	config          *oauth2.Config
	tokFile         string
	entryHandler    http.HandlerFunc
	callbackHandler http.HandlerFunc
	client          *http.Client
}

// InitAuth initialize the internal state of the auth clients.
func InitAuth() {
	http.HandleFunc("/auth/", OAuthHandler)
	http.HandleFunc("/auth/callback/", OAuthCallbackHandler)

	// Initalize provider specific settings.
	googleInitAuth()
}

// re is a regular expression used in getOAuthProvider.
var re, _ = regexp.Compile(`^/auth/(.*/)?(?P<provider>\w*)$`)

// getOAuthProvider cuts off an OAuth provider name from
// a url.
//
// An OAuth provider is the last part of a url path.
func getOAuthProvider(url *url.URL) *provider {
	path := url.Path
	res := re.FindStringSubmatch(path)

	if len(res) <= 2 {
		log.Fatalf("Invalid url %v", url)
	}
	name := res[2]

	var provider *provider
	switch name {
	case "google":
		provider = googleOAuth
	default:
		log.Fatalf("Unknown OAuth provider %v", provider)
	}

	return provider
}

// OAuthHandler handles requests for OAuth authorization.
func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	provider := getOAuthProvider(r.URL)

	tok, err := tokenFromFile(provider.tokFile)
	if err == nil {
		provider.client = provider.config.Client(context.Background(), tok)
		http.Redirect(w, r, "/home", 303)
		return
	}

	provider.entryHandler(w, r)
}

// OAuthCallbackHandler handles a callback from OAuth server.
func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	provider := getOAuthProvider(r.URL)
	provider.callbackHandler(w, r)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to : %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
