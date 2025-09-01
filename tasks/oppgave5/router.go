package oppgave5

import (
	_ "embed"
	"net/http"

	"github.com/fingann/swat/pkg/flags"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

func SetCookieIfNotSet(c *gin.Context, name, value string) {
	_, err := c.Cookie(name)
	if err != nil {
		c.SetCookie(name, value, 0, "", "", false, true)
	}
}

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/oppgave5", func(c *gin.Context) {
		SetCookieIfNotSet(c, "username", "cato")
		SetCookieIfNotSet(c, "admin", "0")
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Oppgave5Page()))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave5Page())
	})

	r.GET("/oppgave5/admin", func(c *gin.Context) {
		admin, err := c.Cookie("admin")
		if err != nil {
			c.HTML(http.StatusOK, "", public.Base(Oppgave5AdminInfo("")))
			return
		}
		if admin == "1" {
			c.HTML(http.StatusOK, "", Oppgave5AdminInfo(flags.Oppgave5Flag))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave5AdminInfo(""))
	})

	return nil
}
