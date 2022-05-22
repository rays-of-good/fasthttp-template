package main

import (
	"os"

	iinternal "github.com/rays-of-good/fasthttp-template/internal"
	idatabase "github.com/rays-of-good/fasthttp-template/internal/database"

	"github.com/sakirsensoy/genv"
	"github.com/sakirsensoy/genv/dotenv"
)

var (
	envfile string = ".env"
)

func main() {
	if len(os.Args) > 1 {
		envfile = os.Args[1]
	}

	dotenv.Load(envfile)

	configuration := iinternal.Configuration{
		Host:  genv.Key("HOST").String(),
		Email: genv.Key("EMAIL").String(),
		Database: idatabase.Configuration{
			DSN: genv.Key("DATABASE_DSN").String(),
		},
	}

	err := iinternal.Main(&configuration)
	if err != nil {
		panic(err)
	}
}
