package config

type Config struct {
	Environment string `envconfig:"environment" required:"true"`
	Port        string `envconfig:"port" default:"9000"`
	AdminUser   string `envconfig:"admin_user" default:"odin"`
	AdmisPass   string `envconfig:"admin_pass" default:"odin123"`
}
