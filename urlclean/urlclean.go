package urlclean

import (
	"embed"
	"encoding/json"
	"regexp"

	"net/url"
)

type Provider struct {
	URLPattern        string
	CompleteProvider  bool
	Rules             []string
	ReferralMarketing []string
	Exceptions        []string
	RawRules          []string
	Redirections      []string
	ForceRedirection  bool
}

type URLRules struct {
	Providers map[string]Provider
}

// Steal data format and data from
// https://gitlab.com/ClearURLs/rules/-/raw/master/data.min.json
//

var ClearURLsRules URLRules

//go:embed assets/*
var embeddedFilesystem embed.FS

func Init() {
	data, _ := embeddedFilesystem.ReadFile("assets/data.min.json")
	json.Unmarshal(data, &ClearURLsRules)
}

func CleanWithProvider(provider Provider, stringToClean string) (string, bool) {
	didSomething := false
	parsedURL, _ := url.Parse(stringToClean)
	query := parsedURL.Query()
	for key := range query {
		for _, list := range [][]string{provider.Rules, provider.ReferralMarketing} {
			for _, rule := range list {
				re := regexp.MustCompile("(?i)^" + rule + "$")
				if re.MatchString(key) {
					query.Del(key)
					didSomething = true
					break
				}
			}
		}
	}

	parsedURL.RawQuery = query.Encode()

	// Apply rawRules to the URL path or entire string
	cleanedURL := parsedURL.String()
	for _, rawRule := range provider.RawRules {
		re := regexp.MustCompile(rawRule)
		cleanedURL = re.ReplaceAllString(cleanedURL, "")
	}
	return cleanedURL, didSomething
}

func CleanURLs(stringToClean string) (string, bool) {
	parsedURL, err := url.Parse(stringToClean)
	if err != nil {
		return stringToClean, false // Return original if parsing fails
	}
	unescapeAtEnd := false
	unescaped, _ := url.QueryUnescape(parsedURL.String())
	if stringToClean == unescaped {
		unescapeAtEnd = true
	}
	matchingProvider := Provider{}

	for providerName, provider := range ClearURLsRules.Providers {
		if providerName != "globalRules" {
			pattern := provider.URLPattern
			re := regexp.MustCompile(pattern)
			if re.MatchString(stringToClean) {
				// Check exceptions
				for _, ex := range provider.Exceptions {
					re := regexp.MustCompile(ex)
					if re.MatchString(stringToClean) {
						continue // Exception matched, skip key
					}
				}
				matchingProvider = provider
				break
			}
		}
	}
	cleanedURL, didSomethingA := CleanWithProvider(ClearURLsRules.Providers["globalRules"], stringToClean)
	didSomethingB := false
	if matchingProvider.URLPattern != "" {
		cleanedURL, didSomethingB = CleanWithProvider(matchingProvider, stringToClean)
	}
	if unescapeAtEnd {
		cleanedURL, _ = url.QueryUnescape(cleanedURL)
	}
	if !didSomethingA && !didSomethingB {
		return stringToClean, false
	}
	return cleanedURL, (didSomethingA || didSomethingB)
}
