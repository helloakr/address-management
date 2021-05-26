package controllers

import (
	"hummingbird/servicesOption"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 96, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	session.Save()
	servicesOption.Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	captchaId := session.Get("captcha")
	if captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		session.Delete("captcha")
		_ = session.Save()
		return false
	}
}

func VerifyOption(c *gin.Context) {
	value := c.Query("verCode")
	if CaptchaVerify(c, value) {
		c.JSON(http.StatusOK, gin.H{"status": 1})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0})
	}
}

func CaptOption(c *gin.Context) {
	Captcha(c, 4)
}
