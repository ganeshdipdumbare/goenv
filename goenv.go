package goenv

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

// SyncEnvVar sync env variable in the input structure,
// if the env variable is not set then the default value
// will be kept as it is.
func SyncEnvVar(env interface{}) {
	bs, err := json.Marshal(&env)
	if err != nil {
		log.Fatal(err)
	}

	envVar := map[string]string{}
	err = json.Unmarshal(bs, &envVar)
	if err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv()
	for k := range envVar {
		val := viper.GetString(k)
		if val != "" {
			envVar[k] = val
		}
	}

	bs, err = json.Marshal(envVar)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bs, &env)
	if err != nil {
		log.Fatal(err)
	}
}
