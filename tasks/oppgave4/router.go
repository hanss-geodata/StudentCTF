package oppgave4

import (
	_ "embed"
	"net/http"

	"github.com/fingann/swat/pkg/codes"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

var oppgave4 string
var Code = "code{}"

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/oppgave4", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Oppgave4Page()))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave4Page())
	})

	r.POST("/oppgave4/unlock", func(c *gin.Context) {
		codeValue := c.PostForm("code")
		taskName, ok := codes.CheckCode(codeValue)
		if ok {
			c.JSON(http.StatusOK, gin.H{"name": taskName})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Code"})
	})

	return nil
}
