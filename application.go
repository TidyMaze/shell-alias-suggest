package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type alias struct {
	short string
	long  string
}

func (a alias) String() string {
	return "alias(short=\"" + a.short + "\", long=\"" + a.long + "\")"
}

func extractAliases(in string) []alias {
	re := regexp.MustCompile(`alias (\w+)='(.+)'`)
	matches := re.FindAllStringSubmatch(in, -1)
	aliases := make([]alias, len(matches))

	for i, match := range matches {
		aliases[i] = alias{match[1], match[2]}
	}

	fmt.Print("Parsed :")
	fmt.Println(aliases)
	return aliases
}

func recommend(aliases []alias, command string) []alias {
	matchingAliases := []alias{}

	for _, alias := range aliases {
		if strings.Contains(command, alias.long) {
			matchingAliases = append(matchingAliases, alias)
		}
	}

	return matchingAliases
}

func fancyPrintRecommendations(aliases []alias, command string) string {
	recommendations := recommend(aliases, command)
	if len(recommendations) == 0 {
		return ""
	}

	recoAsString := ""

	for _, alias := range recommendations {
		recoAsString += fmt.Sprintf("\t-\t\"%s\" => \"%s\"\n", alias.long, alias.short)
	}

	return fmt.Sprintf("You could use following aliases :\n%s", recoAsString)
}

func queryAliasCmd() string {
	out, err := exec.Command("C:\\Program Files\\Git\\bin\\bash.exe", "--login", "-i", "-c", "alias").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {
	command := "ls -l file"
	aliases := extractAliases(queryAliasCmd())

	fmt.Println(fancyPrintRecommendations(aliases, command))
}
