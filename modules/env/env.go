package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port            string `envconfig:"PORT"`
	LineAccessToken string `envconfig:"LINE_ACCESS_TOKEN"`
}

var Env Environment

func init() {
	loadEnvironment()
}

func loadEnvironment() {
	if err := envconfig.Process("", &Env); err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %s", err.Error())
	}
}
