package config

type ZenformConfig struct {
	Zendesk struct {
		Subdomain string `yaml:"subdomain"`
		Email     string `yaml:"email"`
		Token     string `yaml:"token"`
	} `yaml:"zendesk"`
}
