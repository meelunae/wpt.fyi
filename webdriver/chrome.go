package webdriver

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var (
	chromeDriverPath = flag.String("chromedriver_path", "", "Path to the chromedriver binary")
	chromePath       = flag.String("chrome_path", "", "Path to the chrome binary")
)

// ChromeWebDriver starts up ChromeDriver on the given port.
func ChromeWebDriver(port int, options []selenium.ServiceOption) (*selenium.Service, selenium.WebDriver, error) {
	if *chromePath == "" {
		panic("-chrome_path not specified")
	}
	if *chromeDriverPath == "" {
		panic("-chromedriver_path not specified")
	}

	service, err := selenium.NewChromeDriverService(*chromeDriverPath, port, options...)
	if err != nil {
		panic(err)
	}

	// Connect to the WebDriver instance running locally.
	seleniumCapabilities := selenium.Capabilities{
		"browserName": "chrome",
	}

	ChromeCapabilities := chrome.Capabilities{
		Args: []string{"no-sandbox", "disable-dev-shm-usage"},
	}
	chromeAbsPath, err := filepath.Abs(*chromePath)
	if err != nil {
		panic(err)
	}
	ChromeCapabilities.Path = chromeAbsPath
	seleniumCapabilities.AddChrome(ChromeCapabilities)

	// selenium.NewChromeDriverService unconditionally specifies --url-base=wd/hub.
	wd, err := selenium.NewRemote(
		seleniumCapabilities,
		fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	return service, wd, err
}
