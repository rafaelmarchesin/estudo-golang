package system

import (
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// WebDriverInit inicializa o driver do Google Chrome

func WebDriverInit() selenium.WebDriver {
	var wd selenium.WebDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			// "--headless",
			"--no-sandbox",
			"--start-maximized",
			"--window-size=1920,1080",
			"--disable-crash-reporter",
			"--hide-scrollbars",
			"--disable-gpu",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
		},
		ExcludeSwitches: []string{
			"--enable-automation", // Remove a mensagem de que o Chrome está rodando uma automação
		},
	}
	caps.AddChrome(chromeCaps)
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Println(err)
	}

	return wd
}
