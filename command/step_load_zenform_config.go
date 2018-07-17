package command

import (
	"io/ioutil"
	"os"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/xflagstudio/zenform/config"
	"gopkg.in/yaml.v2"
)

const zenformConfigFileName = "zenform.yml"

func stepLoadZenformConfig(exe *StepExecutor, zfconfig *config.ZenformConfig, zd *zendesk.Client) func() error {
	return func() error {
		if _, err := os.Stat(zenformConfigFileName); os.IsNotExist(err) {
			exe.Error(zenformConfigFileName + " Not Found")
			os.Exit(1)
		}
		exe.Success("Found " + zenformConfigFileName)

		configYmlStr, _ := ioutil.ReadFile(zenformConfigFileName)
		err := yaml.Unmarshal(configYmlStr, zfconfig)
		if err != nil {
			exe.Error(err.Error())
			os.Exit(1)
		}

		exe.Step("Configure Zendesk subdomain and credentials", func() error {
			if !isValidCredential(zfconfig) {
				exe.Error("Invalid config")
				os.Exit(1)
			}

			zd.SetSubdomain(zfconfig.Zendesk.Subdomain)
			zd.SetCredential(zendesk.NewAPITokenCredential(zfconfig.Zendesk.Email, zfconfig.Zendesk.Token))
			exe.Success("OK")
			return nil
		})
		return nil
	}
}

func isValidCredential(zfconfig *config.ZenformConfig) bool {
	if zfconfig.Zendesk.Subdomain == "" ||
		zfconfig.Zendesk.Email == "" ||
		zfconfig.Zendesk.Token == "" {
		return false
	}
	return true
}
