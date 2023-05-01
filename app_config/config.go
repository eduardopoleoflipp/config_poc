package app_config

import (
	ConfigPkg "config/poc/config_pkg"
	"fmt"
	"time"
)

type SearchPlatformAPIConfig struct {
	Host            string        `mapstructure:"host"`
	Timeout         time.Duration `mapstructure:"max_related_items"`
	MaxRelatedItems int32         `mapstructure:"time_out"`
	MaxRetries      int           `mapstructure:"max_retries"`
}

type ItemRuleRetrieverConfig struct {
	BaseUrl         string        `mapstructure:"base_url"`
	RefreshInterval time.Duration `mapstructure:"refresh_interval"`
}

type EcomModuleConfig struct {
	URI                 string        `mapstructure:"uri"`
	BaseURL             string        `mapstructure:"base_url"`
	ConfigSyncInterval  time.Duration `mapstructure:"config_sync_interval"`
	MaxEcomItemsPerPage int           `mapstructure:"max_ecom_items_per_page"`
}

type MerchantLabelRetrieverParams struct {
	Bucket           string        `mapstructure:"bucket"`
	Filename         string        `mapstructure:"filename"`
	RuleSyncInterval time.Duration `mapstructure:"rule_sync_interval"`
}

type EventSearches map[EventQuery]EventStack
type EventQuery string
type EventStack map[string]int

type SearchAppConfig struct {
	MixFlyerAndEcomTerms []string `mapstructure:"mix_flyer_and_econ_terms"`
	EventSearches        `mapstructure:"event_searches"`
}

type MerchantLocation struct {
	Table             string `mapstructure:"table"`
	NationalMinFSAsUS int    `mapstructure:"national_min_fsas_us"`
	NationalMinFSAsCA int    `mapstructure:"national_min_fsas_ca"`
}

type AppAdsConfig struct {
	Host    string        `mapstructure:"host"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Config struct {
	SearchPlatformAPI             SearchPlatformAPIConfig      `mapstructure:"search_platform_api"`
	FavouritesDynamoDBTable       string                       `mapstructure:"favourites_dynamo_db_table"`
	MaxMultiSearchQueries         int                          `mapstructure:"max_multi_search_queries"`
	ItemRuleRetriever             ItemRuleRetrieverConfig      `mapstructure:"item_rule_retriever"`
	EcomModuleConfig              EcomModuleConfig             `mapstructure:"ecom_module_config"`
	MerchantLabelRetriever        MerchantLabelRetrieverParams `mapstructure:"merchant_label_retriever"`
	SearchAppConfig               SearchAppConfig              `mapstructure:"search_app_config"`
	MerchantsForceShowItemDetails []int64                      `mapstructure:"merchants_force_show_item_details"`
	MerchantLocation              MerchantLocation             `mapstructure:"merchant_location"`
	AppAdsConfig                  AppAdsConfig                 `mapstructure:"app_ads_config"`
}

func LoadConfig() Config {
	// path relative to the root of the app
	configPath := "./app_config"
	conf := &Config{}
	ConfigPkg.Loader(configPath, conf)

	return *conf
}

func PrintConfig(config Config) {
	fmt.Printf(
		"search_platform_api: \n  host: %s\n  max_related_items: %d\n  time_out: %d\n  max_retries: %d \n",
		config.SearchPlatformAPI.Host,
		config.SearchPlatformAPI.Timeout,
		config.SearchPlatformAPI.MaxRelatedItems,
		config.SearchPlatformAPI.MaxRetries,
	)

	fmt.Printf(
		"favourites_dynamo_db_table: %s\n",
		config.FavouritesDynamoDBTable,
	)

	fmt.Printf(
		"max_multi_search_queries: %d\n",
		config.MaxMultiSearchQueries,
	)

	fmt.Printf(
		"item_rule_retriever: \n  base_url: %s\n  refresh_interval:  %s\n",
		config.ItemRuleRetriever.BaseUrl,
		config.ItemRuleRetriever.RefreshInterval,
	)

	fmt.Printf(
		"ecom_module_config: \n  uri: %s\n  base_url:  %s\n  config_sync_interval:  %d\n  max_ecom_items_per_page: %d\n",
		config.EcomModuleConfig.URI,
		config.EcomModuleConfig.BaseURL,
		config.EcomModuleConfig.ConfigSyncInterval,
		config.EcomModuleConfig.MaxEcomItemsPerPage,
	)

	fmt.Printf(
		"merchant_label_retriever: \n  bucket: %s\n  filename:  %s\n  rule_sync_interval:  %d\n",
		config.MerchantLabelRetriever.Bucket,
		config.MerchantLabelRetriever.Filename,
		config.MerchantLabelRetriever.RuleSyncInterval,
	)

	fmt.Printf(
		"search_app_config:\n  mix_flyer_and_econ_terms %v\n  event_searches %v\n",
		config.SearchAppConfig.MixFlyerAndEcomTerms,
		config.SearchAppConfig.EventSearches,
	)
}
