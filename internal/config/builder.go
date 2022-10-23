package config

import (
	"strings"

	"github.com/spf13/viper"
)

type LoggerConfig struct {
	Level    string
	Encoding string
}

type GrpcServerConfig struct {
	Port string
}

type FirestoreConfig struct {
	ProjectName    string
	CollectionName string
}

type ProjectConfig struct {
	Interval int
}
type Config struct {
	LoggerConfig     LoggerConfig
	GrpcServerConfig GrpcServerConfig
	FirestoreConfig  FirestoreConfig
	ProjectConfig    ProjectConfig
}

// New returns a new instance of config
func New() *Config {
	confer := viper.New()

	confer.AutomaticEnv()
	confer.SetEnvPrefix("scheduler")
	confer.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	conf := &Config{
		LoggerConfig: LoggerConfig{
			Level:    confer.GetString("log.level"),
			Encoding: confer.GetString("log.encoding"),
		},
		GrpcServerConfig: GrpcServerConfig{
			Port: confer.GetString("grpc.port"),
		},
		FirestoreConfig: FirestoreConfig{
			ProjectName:    confer.GetString("firestore.projectname"),
			CollectionName: confer.GetString("firestore.collection"),
		},
		ProjectConfig: ProjectConfig{
			Interval: confer.GetInt("project.interval"),
		},
	}

	return conf
}
