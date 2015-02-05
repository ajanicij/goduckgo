package goduckgo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
	query := "New York City"
	expectedUrl := "https://api.duckduckgo.com/?q=New+York+City&format=json&pretty=1"
	expectedBody := `{
        "Heading" : "New York City"
    }`
	expectedMessage := &Message{}
	expectedMessage.Heading = "New York City"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, expectedBody)
	}))
	defer ts.Close()

	if url := EncodeUrl(query); url != expectedUrl {
		t.Errorf("Got %s, want %s", url, expectedUrl)
	}

	body, err := Do(ts.URL)
	if err != nil {
		t.Errorf("Got %v, want %v", err, nil)
	}

	if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedBody {
		t.Errorf("Got %s, want %s", trimmedBody, expectedBody)
	}

	message := &Message{}
	if err = message.Decode(body); err != nil {
		t.Errorf("Got %v, want %v", err, nil)
	}

	if !reflect.DeepEqual(message, expectedMessage) {
		t.Errorf("Got %v, want %v", message, expectedMessage)
	}
}

func TestEncodeUrl(t *testing.T) {
	query := "!gi New York City"
	expectedUrl := "https://api.duckduckgo.com/?q=%21gi+New+York+City&format=json&pretty=1&no_redirect=1"

	url := EncodeUrl(query)
	if url != expectedUrl {
		t.Errorf("Got %s, want %s", url, expectedUrl)
	}

	query = "New York City"
	expectedUrl = "https://api.duckduckgo.com/?q=New+York+City&format=json&pretty=1"
	url = EncodeUrl(query)
	if url != expectedUrl {
		t.Errorf("Got %s, want %s", url, expectedUrl)
	}

}
