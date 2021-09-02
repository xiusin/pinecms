package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

var captchaDrivers = struct {
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}{}

func Get(captchaType string) (string, string, error) {
	var driver base64Captcha.Driver
	switch captchaType {
	case "string":
		driver = captchaDrivers.DriverString.ConvertFonts()
	case "math":
		driver = captchaDrivers.DriverMath.ConvertFonts()
	case "chinese":
		driver = captchaDrivers.DriverChinese.ConvertFonts()
	default:
		driver = captchaDrivers.DriverDigit
	}
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

func Verify(id, value string) bool {
	return store.Verify(id, value, true)
}
