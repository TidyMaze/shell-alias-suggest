package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var globalAliases = []alias{
	{"ll", "ls -l"},
	{"ls", "ls -F --color=auto --show-control-chars"},
	{"mkdir", "mkdir -pv"},
}

func TestExtractAliases(t *testing.T) {
	aliases := globalAliases

	tables := []struct {
		raw    string
		result []alias
	}{
		{"alias ll='ls -l'\nalias ls='ls -F --color=auto --show-control-chars'\nalias mkdir='mkdir -pv'", aliases},
	}

	for _, test := range tables {
		got := extractAliases(test.raw)
		assert.Equal(t, got, test.result)
	}
}

func TestRecommend(t *testing.T) {
	aliases := globalAliases

	tables := []struct {
		aliases []alias
		command string
		result  []alias
	}{
		{aliases, "ls -l file.txt", []alias{alias{"ll", "ls -l"}}},
		{aliases, "ls -F --color=auto --show-control-chars file.txt", []alias{{"ls", "ls -F --color=auto --show-control-chars"}}},
		{aliases, "BLA", []alias{}},
		{aliases, "ls -l | ls -F --color=auto --show-control-chars file.txt", []alias{{"ll", "ls -l"}, {"ls", "ls -F --color=auto --show-control-chars"}}},
	}

	for _, test := range tables {
		got := recommend(test.aliases, test.command)
		assert.Equal(t, got, test.result)
	}

}

func TestAliasString(t *testing.T) {
	aliasData := alias{"short", "long"}
	got := fmt.Sprintf("%s", aliasData)
	expected := "alias(short=\"short\", long=\"long\")"

	assert.Equal(t, got, expected)
}

func TestFancyPrintRecommendations(t *testing.T) {
	aliases := globalAliases

	tables := []struct {
		aliases []alias
		command string
		result  string
	}{
		{aliases, "ls -l file.txt", "You could use following aliases :\n\t-\t\"ls -l\" => \"ll\"\n\t=> result: \"ll file.txt\"\n"},
		{aliases, "ls -F --color=auto --show-control-chars file.txt", "You could use following aliases :\n\t-\t\"ls -F --color=auto --show-control-chars\" => \"ls\"\n\t=> result: \"ls file.txt\"\n"},
		{aliases, "ls -l | ls -F --color=auto --show-control-chars file.txt", "You could use following aliases :\n\t-\t\"ls -l\" => \"ll\"\n\t-\t\"ls -F --color=auto --show-control-chars\" => \"ls\"\n\t=> result: \"ll | ls file.txt\"\n"},
		{aliases, "BLA", ""},
	}

	for _, test := range tables {
		got := fancyPrintRecommendations(test.aliases, test.command)
		assert.Equal(t, test.result, got)
	}

}
