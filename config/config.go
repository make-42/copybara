package config

import (
	"copybara/regex"
	"copybara/urlclean"
	"copybara/utils"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v2"
)

type ConfigS struct {
	NotificationsOnAppliedAutomations bool
	ExtraURLCleaningRulesAndOverrides map[string]urlclean.Provider
	ExtraRegexRules                   []regex.Rules
	EnableRegexAutomations            bool
	EnableURLCleaning                 bool
}

var DefaultConfig = ConfigS{
	NotificationsOnAppliedAutomations: true,
	ExtraRegexRules: []regex.Rules{
		{
			IsURLRule:   true,
			Pattern:     "^https?:\\/\\/(?:[a-z0-9-]+\\.)*?instagram\\.com\\/reel",
			ReplaceWith: "https://www.ddinstagram.com/reel",
		},
		{
			IsURLRule:   true,
			Pattern:     "^https?:\\/\\/(?:[a-z0-9-]+\\.)*?x\\.com",
			Exceptions:  []string{"^https?:\\/\\/(?:[a-z0-9-]+\\.)*?x\\.com$", "^https?:\\/\\/(?:[a-z0-9-]+\\.)*?x\\.com/$"},
			ReplaceWith: "https://fxtwitter.com",
		},
	},
	ExtraURLCleaningRulesAndOverrides: map[string]urlclean.Provider{"exampleoverride": urlclean.ClearURLsRules.Providers["amazon"]},
	EnableRegexAutomations:            true,
	EnableURLCleaning:                 true,
}

var Config ConfigS

func Init() {
	configPath := configdir.LocalConfig("ontake", "copybara")
	err := configdir.MakePath(configPath) // Ensure it exists.
	utils.CheckError(err)

	configFile := filepath.Join(configPath, "config.yml")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		fh, err := os.Create(configFile)
		utils.CheckError(err)
		defer fh.Close()

		encoder := yaml.NewEncoder(fh)
		encoder.Encode(&DefaultConfig)
		Config = DefaultConfig
	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		utils.CheckError(err)
		defer fh.Close()

		decoder := yaml.NewDecoder(fh)
		decoder.Decode(&Config)
	}
	for key := range Config.ExtraURLCleaningRulesAndOverrides {
		urlclean.ClearURLsRules.Providers[key] = Config.ExtraURLCleaningRulesAndOverrides[key]
	}
	regex.ExtraRules = Config.ExtraRegexRules
}
