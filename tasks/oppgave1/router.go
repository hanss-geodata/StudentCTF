package oppgave1

import (
	_ "embed"
	"math/rand"
	"net/http"
	"sync"

	"github.com/fingann/swat/pkg/flags"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

//go:embed openapi.yaml
var opeapi string

//var visitorCount int = 62651

var APIKEY = "f414d481-8f7b-41e5-ab5e-05814bb1f509"
var Flag = "flag{NEVER_HARD_CODE_SECRETS}"

var once sync.Once

/* type metrics struct {
	Visitors int
} */

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/oppgave1", func(c *gin.Context) {
		/* once.Do(func() {
			go func() {
				// increate visitorCount by random number
				for i := 0; i < 1000; i++ {
					randomVisitors := rand.Intn(10)
					visitorCount = randomVisitors
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

				}
			}()
		})
		visitorCount++ */
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Oppgave1Page()))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave1Page())
	})

	r.GET("/oppgave1/openapi", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(OpenAPIPage()))
			return
		}
		c.HTML(http.StatusOK, "", OpenAPIPage())
	})

	r.GET("/oppgave1/openapi.yaml", func(c *gin.Context) {
		c.String(http.StatusOK, opeapi)
	})

	r.GET("/oppgave1/rng/generateNumber", func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != APIKEY {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"visitors": rand.Intn(1000),
		})
	})
	r.GET("/oppgave1/rng/flag", func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != APIKEY {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"flag": flags.Oppgave1Flag,
		})
	})

	return nil
}
