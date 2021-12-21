package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Environment is for labeling the environment which the application runs in.
type Environment string

// Application environments
const (
	Local       Environment = "local"
	Development Environment = "development"
	Production  Environment = "production"
)

type envVars struct {
	Environment Environment `envconfig:"environment" default:"local"`

	MongoDatabaseURI  string `split_words:"true" default:"mongodb://mongoadmin:secret@localhost:27017"`
	MongoDatabaseName string `split_words:"true" default:"microservice"`

	RedisURL            string `split_words:"true" default:"localhost:6379"`
	ZipkinEndpoint      string `split_words:"true" default:"http://localhost:9411/api/v1/spans"`
	ZipkinSchemaVersion string `split_word:"true" default:"v1"`
	ZipkinServiceName   string `split_words:"true" default:"microservice"`
	TraceMaxBytes       int    `split_words:"true" default:"4096"`

	LogLevel                string `split_words:"true" default:"debug"`
	DefaultSearchResultSize int    `split_words:"true" default:"20"`

	EnableAsyncCaching         bool `split_words:"true" default:"true"`
	EnableZipkinLogs           bool `split_words:"true" default:"true"`
	LogZipkinErrors            bool `split_words:"true" default:"false"`
	EnableCarrefourRequestLogs bool `split_words:"true" default:"false"`

	APIMetricsVersion string `envconfig:"api_metrics_version" default:"v1"`

	Port int `split_words:"true" default:"8888"`
}

// Vars are all available config variables in application environment.
var Vars envVars

// Init parses and prepares all config variables.
func Init() {
	override()
	envconfig.MustProcess("", &Vars)

}

// override loads config file to override environment vars.
// This small feature targets local development environment.
func override() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}

	var configVars map[string]string
	if err := json.Unmarshal(b, &configVars); err != nil {
		panic(err)
	}

	for k, v := range configVars {
		os.Setenv(k, v)
	}
}
