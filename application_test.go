package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtractAliases(t *testing.T) {
	aliases := []alias{
		alias{"ll", "ls -l"},
		alias{"ls", "ls -F --color=auto --show-control-chars"},
	}

	tables := []struct {
		raw    string
		result []alias
	}{
		{"alias ll='ls -l'\nalias ls='ls -F --color=auto --show-control-chars'", aliases},
	}

	for _, test := range tables {
		got := extractAliases(test.raw)
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("extractAliases(%s) is incorrect, got %s, expected %s", test.raw, got, test.result)
		}
	}
}

func TestRecommend(t *testing.T) {
	aliases := []alias{
		alias{"ll", "ls -l"},
		alias{"ls", "ls -F --color=auto --show-control-chars"},
	}

	tables := []struct {
		aliases []alias
		command string
		result  []string
	}{
		{aliases, "ls -l file.txt", []string{"ll"}},
		{aliases, "ls -F --color=auto --show-control-chars file.txt", []string{"ls"}},
		{aliases, "BLA", []string{}},
	}

	for _, test := range tables {
		got := recommend(test.aliases, test.command)
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("Recommend(%s) is incorrect, got %s, expected %s", test.command, got, test.result)
		}
	}

}

func TestAliasString(t *testing.T) {
	aliasData := alias{"short", "long"}
	got := fmt.Sprintf("%s", aliasData)
	expected := "alias(short=\"short\", long=\"long\")"

	if got != expected {
		t.Errorf("Wrong to string for alias %s, got %s instead of %s", aliasData, got, expected)
	}
}

func TestFancyPrintRecommendations(t *testing.T) {
	aliases := []alias{
		alias{"ll", "ls -l"},
		alias{"ls", "ls -F --color=auto --show-control-chars"},
	}

	tables := []struct {
		aliases []alias
		command string
		result  string
	}{
		{aliases, "ls -l file.txt", "You could use following aliases : ll"},
		{aliases, "ls -F --color=auto --show-control-chars file.txt", "You could use following aliases : ls"},
		{aliases, "BLA", ""},
	}

	for _, test := range tables {
		got := fancyPrintRecommendations(test.aliases, test.command)
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("Recommend(%s) is incorrect, got \"%s\", expected \"%s\"", test.command, got, test.result)
		}
	}

}
