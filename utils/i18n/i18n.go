package i18n

import (
	"fmt"
	"os"
	"strings"
)

var knownTranslations = map[string][]string{
	"dpos": {
		"default",
		"en_US",
		"fr_FR",
		"zh_CN",
		"ja_JP",
		"zh_TW",
		"it_IT",
		"de_DE",
		"ko_KR",
	},
	// only use for unit test
	"test": {
		"default",
		"en_US",
	},
}

func loadSystemLanguage() string {
	// Implement the following locale priority order: LC_ALL, LC_MESSAGE, LANG

	langStr := os.Getenv("LC_ALL")
	if langStr == "" {
		langStr = os.Getenv("LC_MESSAGES")
	}
	if langStr == "" {
		langStr = os.Getenv("LANG")
	}

	if langStr == "" {
		fmt.Println("Couldn't find the LC_ALL, LC_MESSAGE, or LANG environment variables, defaulting to en_US ")
		return "default"
	}
	pieces := strings.Split(langStr, ".")
	if len(pieces) != 2 {
		fmt.Println("")
		return "default"
	}
	return pieces[0]
}
