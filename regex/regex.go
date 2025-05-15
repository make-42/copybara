package regex

import (
	"net/url"
	"regexp"
)

type Rules struct {
	IsURLRule   bool
	Pattern     string
	Exceptions  []string
	ReplaceWith string
}

var ExtraRules = []Rules{}

func Clean(stringToClean string) (string, bool) {
	didSomething := false
	isURL := true
	_, err := url.Parse(stringToClean)
	if err != nil {
		isURL = false
	}
	for _, rule := range ExtraRules {
		if rule.IsURLRule == isURL {
			re := regexp.MustCompile(rule.Pattern)
			if re.MatchString(stringToClean) {
				exceptionFound := false
				for _, exception := range rule.Exceptions {
					rexception := regexp.MustCompile(exception)
					if rexception.MatchString(stringToClean) {
						exceptionFound = true
						break
					}
				}
				if !exceptionFound {
					stringToClean = re.ReplaceAllString(stringToClean, rule.ReplaceWith)
					didSomething = true
				}
			}
		}
	}
	return stringToClean, didSomething
}
