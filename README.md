goduckgo
========

Go package for DuckDuckGo API

[DuckDuckGo](http://duckduckgo.com) is a search engine that:
* emphasizes privacy
* does not record user information
* breaks out of the filter bubble

The API is described [here](http://duckduckgo.com/api.html). For
example, the URL for querying about New York City is

[http://api.duckduckgo.com/?q=New+York+City&format=json&pretty=1](http://api.duckduckgo.com/?q=New+York+City&format=json&pretty=1)

The previous query causes DuckDuckGo to return the result in JSON format.

Function goduckgo.Query declared as

    func Query(query string) (*Message, error)

generates the URL, sends it to DuckDuckGo, receives the result, unmarshals from
JSON format to Message structure and returns a pointer to the structure.

## Installation

At this moment, there is no automatic installation procedure. It is assumed that
the package path is `github.com/ajanicij/goduckgo`, which means that the source
file duckduck.go should be manually copied to the directory

`$(GOROOT)/src/github.com/ajanicij/goduckgo/`

and build by running command

`go install github./com/ajanicij/goduckgo`.

On an x86 Linux system, that will
generate the binary `$(GOROOT)/pkg/linux_386/github.com/ajanicij/goduckgo.a`.

## Usage

Look at the source for the command-line utility, `askduck.go`. It imports
package `github.com/ajanicij/goduckgo`, generates the query in the variable
`query` (for example, "New York City") and passes it to function
`goduckgo.Query`. That function returns two values: `*Message` and `error`.

## Command-line utility

The source code of the command-line utility is `askduck.go`. It builds to command
`askduck`. Its usage is:
`askduck [{flags}] <query>`

Flags determine which fields we will see in the result.
For example, if we want to search for "New York City," we can issue command
`askduck -All New York City`

Flag -All tells the command that we want all fields.

Command
`askduck -help`

(or `askduck` without any flags) will give us a help string that lists
all available options:

    Usage of ./askduck:
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

For example, query

`askduck -Definition Europe`

produces the following:

    Definition: The sixth-largest continent, extending west from the Dardanelles, Black Sea, and Ural Mountains.

