package config

import "github.com/spf13/viper"

func getString(vp *viper.Viper, key string, defaultVal ...string) string {
	if len(defaultVal) > 0 {
		vp.SetDefault(key, defaultVal[0])
	}

	return vp.GetString(key)
}

func getBool(vp *viper.Viper, key string, defaultVal ...bool) bool {
	if len(defaultVal) > 0 {
		vp.SetDefault(key, defaultVal[0])
	}

	return vp.GetBool(key)
}

func getInt(vp *viper.Viper, key string, defaultVal ...int) int {
	if len(defaultVal) > 0 {
		vp.SetDefault(key, defaultVal[0])
	}

	return vp.GetInt(key)
}
