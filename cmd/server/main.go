package main

import (
	"fmt"

	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/kelseyhightower/envconfig"
	"github.com/mtavano/ipfs-auth-proxy/config"
	"github.com/mtavano/ipfs-auth-proxy/internal/api"
	"github.com/mtavano/ipfs-auth-proxy/internal/app"
	"github.com/mtavano/ipfs-auth-proxy/internal/database"
)

func main() {
	fmt.Println("Running IPFS auth proxy")
	var conf config.Config
	envconfig.MustProcess("", &conf)

	storage := database.NewStorage(conf.AdminUser, conf.AdmisPass)

	var blockApi iface.BlockAPI
	if conf.Environment == "prod" {
		ipfs, err := httpapi.NewLocalApi()
		check(err)

		blockApi = ipfs.Block()
	}

	appCtx, err := app.NewContext(storage, blockApi)
	check(err)

	apiServer := api.New(appCtx, &conf)

	apiServer.Run()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
