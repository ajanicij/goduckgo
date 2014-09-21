// Package goduckgo provides the functionality for using
// DuckDuckGo API. For the description of the API, visit
// http://duckduckgo.com/api.html.
package goduckgo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var baseUrl = "http://api.duckduckgo.com/?q=%s&format=json&pretty=1"

// Type Message is a structure containing all the information returned by
// DDG for a query.
type Message struct {
	Definition       string
	DefinitionSource string
	Heading          string
	AbstractText     string
	Abstract         string
	AbstractSource   string
	Image            string
	Type             string
	AnswerType       string
	Redirect         string
	DefinitionURL    string
	Answer           string
	AbstractURL      string
	Results          Results
	RelatedTopics    RelatedTopics
}

func (message *Message) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}

	return nil
}

type Result struct {
	Result   string
	FirstURL string
	Text     string
}

// Method Show of struct Result writes Result to standard output
func (result *Result) Show(prefix string) {
	result.Fshow(os.Stdout, prefix)
}

func (result *Result) Fshow(w io.Writer, prefix string) {
	fmt.Fprintln(w, prefix, "Result:", result.Result)
	fmt.Fprintln(w, prefix, "First URL:", result.FirstURL)
	fmt.Fprintln(w, prefix, "Text:", result.Text)
}

type Results []Result

type RelatedTopics []RelatedTopic

type RelatedTopic struct {
	Result   string
	Icon     Icon
	FirstURL string
	Text     string
}

func (topic *RelatedTopic) Show(prefix string) {
	topic.Fshow(os.Stdout, prefix)
}

func (topic *RelatedTopic) Fshow(w io.Writer, prefix string) {
	fmt.Fprintln(w, prefix, "Result:", topic.Result)
	fmt.Fprintln(w, prefix, "Icon:")
	topic.Icon.Fshow(w, prefix+prefix)
	fmt.Fprintln(w, prefix, "FirstURL:", topic.FirstURL)
	fmt.Fprintln(w, prefix, "Text:", topic.Text)
}

type Icon struct {
	URL    string
	Height interface{} // can be string or number ("16" or 16)
	Width  interface{} // can be string or number ("16" or 16)
}

func (icon *Icon) Show(prefix string) {
	icon.Fshow(os.Stdout, prefix)
}

func (icon *Icon) Fshow(w io.Writer, prefix string) {
	fmt.Fprintln(w, prefix, "URL:", icon.URL)
	fmt.Fprintln(w, prefix, "Height:", icon.Height)
	fmt.Fprintln(w, prefix, "Width:", icon.Width)
}

func EncodeUrl(query string) string {
	queryEnc := url.QueryEscape(query)
	return fmt.Sprintf(baseUrl, queryEnc)
}

func Do(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Query(query string) (*Message, error) {
	ddgUrl := EncodeUrl(query)

	body, err := Do(ddgUrl)
	if err != nil {
		return nil, err
	}

	message := &Message{}
	if err = message.Decode(body); err != nil {
		return nil, err
	}

	return message, nil
}
