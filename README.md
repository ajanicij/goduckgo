goduckgo [![GoDoc](https://godoc.org/github.com/ajanicij/goduckgo/goduckgo?status.png)](http://godoc.org/github.com/ajanicij/goduckgo/goduckgo)
========

Go package for DuckDuckGo Instant Answer API.

[DuckDuckGo](http://duckduckgo.com) is a search engine that:

* Emphasizes privacy
* Does not record user information
* Breaks out of the filter bubble

The Instant Answer API is described [here](http://duckduckgo.com/api.html). For
example, the URL for querying about New York City is

[http://api.duckduckgo.com/?q=New+York+City&format=json&pretty=1](http://api.duckduckgo.com/?q=New+York+City&format=json&pretty=1)

The previous query causes DuckDuckGo to return the result in JSON format.

Function goduckgo.Query declared as

```
func Query(query string) (*Message, error)
```

generates the URL, sends it to DuckDuckGo, receives the result, unmarshals from
JSON format to Message structure and returns a pointer to the structure.

Installation
------------

```
go get -u github.com/ajanicij/goduckgo
```

Usage
-----

Look at the source for the command-line utility, `main.go`. It imports
package `github.com/ajanicij/goduckgo/goduckgo`, generates the query in the variable
`query` (for example, "New York City") and passes it to function
`goduckgo.Query`. That function returns two values: `*Message` and `error`.

Command-line utility
--------------------

The source code of the command-line utility is `main.go`. It builds `goduckgo`
command. Its usage is:

`goduckgo [{flags}] <query>`

Flags determine which fields we will see in the result.
For example, if we want to search for "New York City," we can issue command

`goduckgo -All New York City`

Flag -All tells the command that we want all fields.

Command

`goduckgo -help`

(or `goduckgo` without any flags) will give us a help string that lists
all available options:

```
Usage of ./goduckgo:
  -Abstract=false: Abstract
  -AbstractSource=false: Abstract Source
  -AbstractText=false: Abstract Text
  -AbstractURL=false: Abstract URL
  -All=false: All Fields
  -Answer=false: Answer
  -AnswerType=false: Answer Type
  -Definition=false: Definition
  -DefinitionSource=false: Definition Source
  -DefinitionURL=false: Definition URL
  -Heading=false: Heading
  -Image=false: Image
  -Redirect=false: Redirect
  -RelatedTopics=false: Related Topics
  -Results=false: Results
  -Type=false: Type
```

For example, query

`goduckgo -Abstract DuckDuckgo`

produces the following:

```
Abstract: DuckDuckGo is an Internet search engine that emphasizes protecting searchers'
privacy and avoiding the "filter bubble" of personalized search results. DuckDuckGo
distinguishes itself from other search engines by not profiling its users and by deliberately
showing all users the same search results for a given search term. DuckDuckGo also
emphasizes getting information from the best sources rather than the most sources,
generating its search results from key crowdsourced sites such as Wikipedia and from
partnerships with other search engines like Yandex, Yahoo!, Bing, Wolfram Alpha and Yummly.
```
