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
	"strings"
)

var baseUrl = "http://api.duckduckgo.com/?q=%s&format=json&pretty=1"
var bangUrl = "http://api.duckduckgo.com/?q=%s&format=json&pretty=1&no_redirect=1"

// Message is a structure containing all the information returned by
// DDG for a query.
//   Abstract: topic summary (can contain HTML, e.g. italics)
//   AbstractText: topic summary (with no HTML)
//   AbstractSource: name of Abstract source
//   AbstractURL: deep link to expanded topic page in AbstractSource
//   Image: link to image that goes with Abstract
//   Heading: name of topic that goes with Abstract
//   Answer: instant answer
//   AnswerType: type of Answer, e.g. calc, color, digest, info, ip, iploc, phone, pw, rand, regexp, unicode, upc, or zip (see goodies & tech pages for examples).
//   Definition: dictionary definition (may differ from Abstract)
//   DefinitionSource: name of Definition source
//   DefinitionURL: deep link to expanded definition page in DefinitionSource
//   RelatedTopics: array of internal links to related topics associated with Abstract
//   Results: array of external links associated with Abstract
//   Type: response category, i.e. A (article), D (disambiguation), C (category), N (name), E (exclusive), or nothing.
//   Redirect: !bang redirect URL
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

// Decode a message given a HTTP response body
func (message *Message) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}

	return nil
}

// Show Result as standard output
func (result *Result) Show(prefix string) {
	result.fshow(os.Stdout, prefix)
}

func (result *Result) fshow(w io.Writer, prefix string) {
	fmt.Fprintln(w, prefix, "Result:", result.Result)
	if !result.Icon.IsEmpty() {
		fmt.Fprintln(w, prefix, "Icon:")
		result.Icon.fshow(w, prefix+prefix)
	}
	fmt.Fprintln(w, prefix, "First URL:", result.FirstURL)
	fmt.Fprintln(w, prefix, "Text:", result.Text)
}

type Results []Result

type RelatedTopics []RelatedTopic

// Result is an external link associated with Abstract
//   Result: HTML link(s) to external site(s)
//   FirstURL: first URL in Result
//   Icon: icon associated with FirstURL
//   Text: text from FirstURL
type Result RelatedTopic

// RelatedTopic is a internal link to related topics associated with Abstract
//	 Result: HTML link to a related topic
//   FirstURL: first URL in Result
//   Icon: icon associated with related topic
// 	 Text: text from first URL
type RelatedTopic struct {
	Result   string
	Icon     Icon
	FirstURL string
	Text     string
}

// Show RelatedTopic as standard output
func (topic *RelatedTopic) Show(prefix string) {
	topic.fshow(os.Stdout, prefix)
}

func (topic *RelatedTopic) fshow(w io.Writer, prefix string) {
	fmt.Fprintln(w, prefix, "Result:", topic.Result)
	if !topic.Icon.IsEmpty() {
		fmt.Fprintln(w, prefix, "Icon:")
		topic.Icon.fshow(w, prefix+prefix)
	}
	fmt.Fprintln(w, prefix, "First URL:", topic.FirstURL)
	fmt.Fprintln(w, prefix, "Text:", topic.Text)
}

// Icon associated with related topics
//   URL: URL of icon
//   Height: height of icon (px)
//   Width: width of icon (px)
type Icon struct {
	URL    string
	Height interface{} // can be string or number ("16" or 16)
	Width  interface{} // can be string or number ("16" or 16)
}

// IsEmpty if all Icon fields are empty
func (icon *Icon) IsEmpty() bool {
	return icon.URL == "" &&
		icon.Height == "" &&
		icon.Width == ""
}

// Show Show as standard output
func (icon *Icon) Show(prefix string) {
	icon.fshow(os.Stdout, prefix)
}

func (icon *Icon) fshow(w io.Writer, prefix string) {
	fmt.Fprintln(w, prefix, "URL:", icon.URL)
	fmt.Fprintln(w, prefix, "Height:", icon.Height)
	fmt.Fprintln(w, prefix, "Width:", icon.Width)
}

// EncodeUrl given a text query
func EncodeUrl(query string) string {
	queryEnc := url.QueryEscape(query)
	if strings.Contains(query, "!") && strings.Index(query, "!") == 0 {
		return fmt.Sprintf(bangUrl, queryEnc)
	} else {
		return fmt.Sprintf(baseUrl, queryEnc)
	}
}

// Do the HTTP requests against API and handle errors
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

// Query the API given a text query, returns a Message
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
