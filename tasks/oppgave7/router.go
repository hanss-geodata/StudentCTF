package oppgave7

import (
	_ "embed"
	"net/http"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

//go:embed stego.jpg
var stego string

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/oppgave7", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Oppgave7Page()))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave7Page())
	})

	r.GET("/oppgave7/stego.jpg", func(c *gin.Context) {
		c.String(http.StatusOK, stego)
	})

	return nil
}
