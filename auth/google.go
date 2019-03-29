package auth

import (
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

var googleOAuth *provider

func googleInitAuth() {
	b, err := ioutil.ReadFile("auth/secret/google_client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json
	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/drive.appdata")
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/drive.metadata.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	googleOAuth = &provider{
		name:            "google",
		config:          config,
		tokFile:         "auth/token/google_token.json",
		entryHandler:    GoogleOAuthHandler,
		callbackHandler: GoogleOAuthCallbackHandler}
}

// GoogleOAuthHandler handles a request for google OAuth authorization.
func GoogleOAuthHandler(w http.ResponseWriter, r *http.Request) {
	stateToken := "state-token"
	authURL := googleOAuth.config.AuthCodeURL(stateToken, oauth2.AccessTypeOffline)
	http.SetCookie(w, &http.Cookie{Name: "state-token", Value: stateToken})
	http.Redirect(w, r, authURL, 303)
}

// GoogleOAuthCallbackHandler handles a callback from google OAuth server, and
// retrieves a code for further authorization process.
func GoogleOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if ck, err := r.Cookie("state-token"); err != nil {
		log.Fatalf("Unable to read state-token cookie %v", err)
	} else if ck.Value != r.FormValue("state") {
		log.Fatalf("state-token is not valid %v", err)
	}

	authCode := r.FormValue("code")
	if authCode == "" {
		log.Fatalf("Unable to read authorization code")
	}

	tok, err := googleOAuth.config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	// saveToken(googleOAuth.tokFile, tok)

	googleOAuth.client = googleOAuth.config.Client(context.Background(), tok)
	http.Redirect(w, r, "/home", 303)
}

// CreateGoogleDriveClient creates a new google drive api client
// from googleOAuth.client.
//
// googleOAuth must be configured previously.
func CreateGoogleDriveClient() (*drive.Service, error) {
	return drive.New(googleOAuth.client)
}
