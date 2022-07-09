package internal

import (
	"crypto/tls"
	"net"
	"os"
	"os/signal"
	"syscall"

	iapi "github.com/rays-of-good/fasthttp-template/internal/api"
	idatabase "github.com/rays-of-good/fasthttp-template/internal/database"
	irenderer "github.com/rays-of-good/fasthttp-template/internal/renderer"
	iserver "github.com/rays-of-good/fasthttp-template/internal/server"

	"github.com/fasthttp/router"

	aheaders "github.com/go-asphyxia/http/headers"
	amethods "github.com/go-asphyxia/http/methods"
	amiddlewares "github.com/go-asphyxia/middlewares"
	atls "github.com/go-asphyxia/tls"
)

type (
	Configuration struct {
		Host  string
		Email string

		Database idatabase.Configuration
	}
)

func Main(configuration *Configuration) (err error) {
	d, err := idatabase.NewDatabase(&configuration.Database)
	if err != nil {
		return
	}
	defer d.Close()

	HSTS := amiddlewares.NewHSTS(31536000)
	CORS := amiddlewares.NewCORS(
		[]string{
			configuration.Host,
		},
		[]string{
			amethods.GET,
			amethods.POST,
			amethods.PUT,
			amethods.DELETE,
			amethods.OPTIONS,
		},
		[]string{
			aheaders.ContentType,
			aheaders.Accept,
			aheaders.Authorization,
		},
	)

	rr := irenderer.NewRenderer(d)

	API := iapi.NewAPI(d)

	r := router.New()

	r.ANY("/", rr.Main())

	api := r.Group("/api")

	users := api.Group("/users")
	users.GET("/", CORS.Middleware(API.GetUsers()))

	api.OPTIONS("/{any:*}", CORS.Handler())

	t, err := atls.NewTLS(atls.Version12)
	if err != nil {
		return
	}

	tlsConfiguration, err := t.Auto(configuration.Email, atls.DefaultCertificatesCachePath, configuration.Host, ("www." + configuration.Host))
	if err != nil {
		return
	}

	http, err := net.Listen("tcp", ":80")
	if err != nil {
		return
	}

	https, err := net.Listen("tcp", ":443")
	if err != nil {
		return
	}

	https = tls.NewListener(https, tlsConfiguration)

	s := iserver.NewServer(HSTS.Middleware(r.Handler), configuration.Host)
	defer s.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	errs := make(chan error)

	go func() {
		errs <- s.Serve(http)
	}()

	go func() {
		errs <- s.Serve(https)
	}()

	select {
	case err = <-errs:
	case <-signals:
	}

	return
}
