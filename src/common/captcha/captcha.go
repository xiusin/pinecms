package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Get(captchaType string) (string, string, error) {
	var driver base64Captcha.Driver

	switch captchaType {
	//case "string":
	//	driver = base64Captcha.NewDriverString(80, 240, 20, 100, 2, 5, nil, nil).ConvertFonts()
	case "math":
		driver = base64Captcha.NewDriverMath(34, 100, 5,
			base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
			nil, nil, []string{"3Dumb.ttf"})
	//case "chinese":
	//	driver = captchaDriverTmp.DriverChinese.ConvertFonts()
	default:
		driver = base64Captcha.DefaultDriverDigit
	}
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

func Verify(id, value string) bool {
	return store.Verify(id, value, true)
}
