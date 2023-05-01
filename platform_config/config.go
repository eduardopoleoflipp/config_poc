/*
	This package also leverages the config_pkg to unmarshal the config
	from a yml file into config structs. It should be it's own package
	centralized outside flipp-mall
*/

package platform_config

import (
	ConfigPkg "config/poc/config_pkg"
	"fmt"

	"github.com/imdario/mergo"
)

type MultiQuerySearchRequestConfig struct {
	SearchQueryResultLimit int    `mapstructure:"search_query_result_limit"`
	SearchQuerySpellcheck  bool   `mapstructure:"search_query_spell_check"`
	Host                   string `mapstructure:"search_query_re_host"`
}

type WatchlistRequestConfig struct {
	Host string `mapstructure:"host"`
}

type AuctionHouseRequestConfig struct {
	Host string `mapstructure:"host"`
}

// Config is used by the app to hold environment config variables
type Config struct {
	MultiQuerySearchRequestConfig MultiQuerySearchRequestConfig `mapstructure:"multi_query_search_request"`
	WatchlistRequestConfig        WatchlistRequestConfig        `mapstructure:"watch_list_request_config"`
	AuctionHouseRequestConfig     AuctionHouseRequestConfig     `mapstructure:"auction_house_request_config"`
}

func LoadConfig() Config {
	// Loads the config
	configPath := "./platform_config"
	var config Config
	ConfigPkg.Loader(configPath, &config)

	overridesPath := "./platform_overrides"
	var overrideConfig Config
	ConfigPkg.Loader(overridesPath, &overrideConfig)
	mergo.Merge(&overrideConfig, config)

	return overrideConfig

}

func PrintConfig(config Config) {
	fmt.Printf(
		"multi_query_search_request: \n  search_query_result_limit: %d\n  search_query_spell_check: %t\n  search_query_re_host: %s\n",
		config.MultiQuerySearchRequestConfig.SearchQueryResultLimit,
		config.MultiQuerySearchRequestConfig.SearchQuerySpellcheck,
		config.MultiQuerySearchRequestConfig.Host,
	)

	fmt.Printf(
		"watch_list_request_config: \n  host: %s\n",
		config.WatchlistRequestConfig.Host,
	)

	fmt.Printf(
		"auction_house_request_config: \n  host: %s\n",
		config.AuctionHouseRequestConfig.Host,
	)
}
