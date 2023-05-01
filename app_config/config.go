package app_config

import (
	ConfigPkg "config/poc/config_pkg"
	"time"
)

type SearchPlatformAPIConfig struct {
	Host            string        `mapstructure:"host"`
	Timeout         time.Duration `mapstructure:"max_related_items"`
	MaxRelatedItems int32         `mapstructure:"time_out"`
	MaxRetries      int           `mapstructure:"max_retries"`
}

type Config struct {
	SearchPlatformAPI SearchPlatformAPIConfig `mapstructure:"search_platform_api"`
	// FavouritesDynamoDBTable string
	// MaxMultiSearchQueries   int
	// EcomModuleConfig
	// MerchantLabelRetriever MerchantLabelRetrieverParams
	// SearchAppConfig
	// ItemRuleRetriever ItemRuleRetrieverConfig
	// AppAdsConfig
	// MerchantLocation
	// // inherits all receiver functions from go-flippmall/config
	// config.Config
	// // MerchantsForceShowItemDetails is a list of merchant ids that must show item details
	// // even if they don't meet specific criteria (i.e. non-indexed flyer)
	// MerchantsForceShowItemDetails []int64
}

func LoadConfig() Config {
	configPath := "."
	conf := &Config{}
	ConfigPkg.Loader(configPath, conf)

	return *conf
}
