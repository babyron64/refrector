package auth

import (
	"net/url"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitAuth()
	os.Exit(m.Run())
}

func TestGetOAuthProvider1(t *testing.T) {
	url, _ := url.Parse("http://example.com/auth/callback/google")
	provider := getOAuthProvider(url)
	if provider.name != "google" {
		t.Logf("provider: %v", provider)
		t.Fail()
	}
}

func TestGetOAuthProvider2(t *testing.T) {
	url, _ := url.Parse("http://example.com/auth/google")
	provider := getOAuthProvider(url)
	if provider.name != "google" {
		t.Logf("provider: %v", provider)
		t.Fail()
	}
}
