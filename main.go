package opshins

import (
	"fmt"
	"github.com/SupremeERG/colorPrint"
	"golang.org/x/exp/slices"
	"log"
	"strings"
)

func prep() {
	log.SetPrefix("opshins: ")
	log.SetFlags(0)
}


// Function for yes or no prompts
func PromptYN(question string, def string, promptor string) string {
	prep()
	if def != "yes" && def != "no" {
		log.Fatal("The default must be \"yes\" or \"no\", not " + fmt.Sprintf("\"%s\"", def))
	}

	var opt2 string
	switch def {
	case "yes":
		opt2 = "no"

	case "no":
		opt2 = "yes"
	}
	optionPrompt := fmt.Sprintf("%s [%s/%s]", question, colorPrint.Color("white", def+colorPrint.Default()), opt2)

	var answer string
	fmt.Print(optionPrompt, promptor)
	fmt.Scanln(&answer)

	yesAnswers := []string{"yes", "y", "yup"}
	noAnswers := []string{"no", "n", "nope"}

	if answer == "" {
		return def
	} else if slices.Contains(yesAnswers, strings.ToLower(answer)) == false && slices.Contains(noAnswers, strings.ToLower(answer)) == false {
		PromptYN(question, def, promptor)
	} else {
		if slices.Contains(yesAnswers, strings.ToLower(answer)) == true {
			return "yes"
		} else {
			return "no"
		}
	}
	return ""
}

var inputs []string
func WhilePrompt(question string, end string, promptor string) []string {
	prep()
	var answer string
	prompt := fmt.Sprintf("%s [\"%s\" to stop]", question, colorPrint.Color("white", end + colorPrint.Default()))

	fmt.Print(prompt, promptor)
	fmt.Scanln(&answer)

	if answer == end {
		return inputs
	} else {
		inputs = append(inputs, answer)
		return WhilePrompt(question, end, promptor)
		
	}
}