package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"github.com/ajanicij/goduckgo"
)

func main() {
	// Read flags from command line
	fl_definition := flag.Bool("Definition", false, "Definition")
	fl_definition_source := flag.Bool("DefinitionSource", false, "Definition Source")
	fl_heading := flag.Bool("Heading", false, "Heading")
	fl_abstract_text := flag.Bool("AbstractText", false, "Abstract Text")
	fl_abstract := flag.Bool("Abstract", false, "Abstract")
	fl_abstract_source := flag.Bool("AbstractSource", false, "Abstract Source")
	fl_image := flag.Bool("Image", false, "Image")
	fl_type := flag.Bool("Type", false, "Type")
	fl_answer_type := flag.Bool("AnswerType", false, "Answer Type")
	fl_redirect := flag.Bool("Redirect", false, "Redirect")
	fl_definition_url := flag.Bool("DefinitionURL", false, "Definition URL")
	fl_answer := flag.Bool("Answer", false, "Answer")
	fl_abstract_url := flag.Bool("AbstractURL", false, "Abstract URL")
	fl_results := flag.Bool("Results", false, "Results")
	fl_related_topics := flag.Bool("RelatedTopics", false, "Related Topics")
	fl_all := flag.Bool("All", false, "All Fields")
	
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: simplequery [{flags}] <query>")
		os.Exit(0)
	}
	query := strings.Join(flag.Args(), " ")
	
	message, err := goduckgo.Query(query)
	CheckError(err)

	if *fl_all || *fl_definition {
		fmt.Println("Definition:", message.Definition)
	}
	if *fl_all || *fl_definition_source {
		fmt.Println("Definition Source:", message.DefinitionSource)
	}
	if *fl_all || *fl_heading {
		fmt.Println("Heading:", message.Heading)
	}
	if *fl_all || *fl_abstract_text {
		fmt.Println("Abstract Text:", message.AbstractText)
	}
	if *fl_all || *fl_abstract {
		fmt.Println("Abstract:", message.Abstract)
	}
	if *fl_all || *fl_abstract_source {
		fmt.Println("Abstract Source:", message.AbstractSource)
	}
	if *fl_all || *fl_image {
		fmt.Println("Image:", message.Image)
	}
	if *fl_all || *fl_type {
		fmt.Println("Type:", TypeDefinition(message.Type))
	}
	if *fl_all || *fl_answer_type {
		fmt.Println("Answer Type:", message.AnswerType)
	}
	if *fl_all || *fl_redirect {
		fmt.Println("Redirect:", message.Redirect)
	}
	if *fl_all || *fl_definition_url {
		fmt.Println("Definition URL:", message.DefinitionURL)
	}
	if *fl_all || *fl_answer {
		fmt.Println("Answer:", message.Answer)
	}
	if *fl_all || *fl_abstract_url {
		fmt.Println("Abstract URL:", message.AbstractURL)
	}
	if *fl_all || *fl_results {
		if message.Results != nil && len(message.Results) != 0 {
			for _, result := range message.Results {
				fmt.Println("Result")
				result.Show("  ")
			}
		}
	}
	if *fl_all || *fl_related_topics {
		if message.RelatedTopics != nil && len(message.RelatedTopics) != 0 {
			for _, topic := range message.RelatedTopics {
				fmt.Println("Related Topic")
				topic.Show("  ")
			}
		}
	}
}

func TypeDefinition(d string) string {
	switch d {
		case "A":
			return "Article"
		case "D":
			return "Disambiguation"
		case "C":
			return "Category"
		case "N":
			return "Name"
		case "E":
			return "Exclusive"
	}
	return "Unknown"
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(-1)
	}
}

