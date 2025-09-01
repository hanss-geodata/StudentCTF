package oppgave8

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) error {
	db, err := SetupDb()
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	r.GET("/oppgave8", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Oppgave8Page("")))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave8Page(""))
	})

	r.POST("/oppgave8/login", func(c *gin.Context) {
		username, _ := c.GetPostForm("username")
		password, _ := c.GetPostForm("password")

		message, err := Login(db, username, password)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.HTML(http.StatusOK, "", Oppgave8Page("Invalid username or password"))
				return
			}
			c.Status(http.StatusInternalServerError)
			if err := Oppgave8Page(err.Error()).Render(c, c.Writer); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			return
		}

		c.HTML(http.StatusOK, "", LoginSuccess(message))
		return
	})
	return nil
}
