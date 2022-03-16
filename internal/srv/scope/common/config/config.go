package config

import "flag"

type PRBConfigMode int

const (
	PRBModeNone PRBConfigMode = 0
	PRBModeAPI  PRBConfigMode = 1
	PRBModeWeb  PRBConfigMode = 2
	PRBModeMono PRBConfigMode = 3
)

func (mode PRBConfigMode) IsNone() bool {
	return mode == PRBModeNone || (!mode.IsAPI() && !mode.IsWeb())
}

func (mode PRBConfigMode) IsAPI() bool  { return mode&PRBModeAPI != 0 }
func (mode PRBConfigMode) IsWeb() bool  { return mode&PRBModeWeb != 0 }
func (mode PRBConfigMode) IsMono() bool { return mode.IsAPI() && mode.IsWeb() }

func NewPRBConfigMode(value string) (mode PRBConfigMode) {
	switch value {
	case "mono":
		mode = PRBModeMono
	case "web":
		mode = PRBModeWeb
	case "api":
		mode = PRBModeAPI
	}

	return
}

type PRBConfig struct {
	BasePath                   string
	APIInterface, WebInterface string
	PRBConfigMode
}

func NewPRBConfig() (config PRBConfig) {
	const (
		basePathName      = "base-path"
		modeName          = "mode"
		webInterfaceName  = "web-interface"
		apiInterfaceName  = "api-interface"
		monoInterfaceName = "mono-interface"
	)

	basePath := flag.String(basePathName, "", "")
	mode := flag.String(modeName, "mono", "mono, api, web")

	webInterface := flag.String(webInterfaceName, ":8080", "")
	apiInterface := flag.String(apiInterfaceName, ":7070", "")
	monoInterface := flag.String(monoInterfaceName, ":8080", "")

	flag.Parse()

	config.BasePath = *basePath
	config.PRBConfigMode = NewPRBConfigMode(*mode)
	switch {
	case config.IsMono() && isFlagPassed(monoInterfaceName):
		config.APIInterface, config.WebInterface = *monoInterface, *monoInterface
	case !config.IsMono() && config.IsWeb() && isFlagPassed(webInterfaceName) && isFlagPassed(apiInterfaceName):
		config.APIInterface, config.WebInterface = *apiInterface, *webInterface
	case !config.IsMono() && config.IsAPI() && isFlagPassed(apiInterfaceName):
		config.APIInterface = *apiInterface
	default:
		panic("bad config")
	}

	return
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
