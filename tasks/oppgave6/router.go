package Oppgave6

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/fingann/swat/pkg/flags"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) error {

	r.GET("/Oppgave6", func(c *gin.Context) {
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Oppgave6HintPage()))
			return
		}
		c.HTML(http.StatusOK, "", Oppgave6HintPage())
	})

	r.GET("/.well-known/Oppgave6.txt", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf(`-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

Mailto:Oppgave6@geodata.no
Contact: https://geodata.no
Expires: 2024-03-14T00:00:00.000Z
Acknowledgments: %s 
Preferred-Languages: en, fr, de
Canonical: https://Oppgave6txt.org/.well-known/Oppgave6.txt
-----BEGIN PGP SIGNATURE-----

iHUEARYKAB0WIQSsP2kEdoKDVFpSg6u3rK+YCkjapwUCY9qRaQAKCRC3rK+YCkja
pwALAP9LEHSYMDW4h8QRHg4MwCzUdnbjBLIvpq4QTo3dIqCUPwEA31MsEf95OKCh
MTHYHajOzjwpwlQVrjkK419igx4imgk=
=KONn
-----END PGP SIGNATURE-----`, flags.Oppgave6Flag)))
	})

	return nil
}
