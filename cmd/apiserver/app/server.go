package app

import (
	"net/http"
	"os"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/shopspring/decimal"
	"golang.org/x/text/language"

	"mgw/mgw-resi/cmd/apiserver/app/routes"
	"mgw/mgw-resi/cmd/apiserver/app/store"
	"mgw/mgw-resi/config"
	"mgw/mgw-resi/pkg/utils/lang"
	"mgw/mgw-resi/pkg/utils/logs"
	timeutil "mgw/mgw-resi/pkg/utils/time"
)

const (
	defReadTimeount   = 10 * time.Second
	defWirteTimeout   = 30 * time.Second
	defMaxHeaderBytes = 1 << 20
)

// Run boot the application server
func Run() {
	decimal.MarshalJSONWithoutQuotes = true
	os.Setenv("TZ", config.Config.System.TimeZone)

	// setup i18n
	bundle := i18n.NewBundle(language.Indonesian)
	bundle.MustLoadMessageFile(config.Config.APIServer.LocalePath + "id.json")
	bundle.MustLoadMessageFile(config.Config.APIServer.LocalePath + "en-US.json")

	// init components
	store.InitDI()
	lang.Init(bundle)
	logs.Init("")
	timeutil.Init(config.Config.System.TimeZone)

	logs.Log.Infof("[Server:Addr]: %s%s\n", config.Config.System.AppServer, config.Config.System.AppAddr)
	logs.Log.Fatal(gracehttp.Serve(initServer()))
}

func initServer() *http.Server {
	server := http.Server{
		Addr:           config.Config.System.AppAddr,
		Handler:        routes.Init(config.Config.System.Mode),
		ReadTimeout:    defReadTimeount,
		WriteTimeout:   defWirteTimeout,
		MaxHeaderBytes: defMaxHeaderBytes,
	}

	return &server
}
