package main

import (
	"context"
	"fmt"

	httpapi "github.com/ipfs/go-ipfs-http-client"

	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/kelseyhightower/envconfig"
	"github.com/mtavano/ipfs-auth-proxy/config"
)

func main() {
	fmt.Println("[OPS]")
	var conf config.Config
	envconfig.MustProcess("", &conf)
	fmt.Println(conf)

	ipfs, err := httpapi.NewLocalApi()
	check(err)
	fmt.Println(ipfs)

	reader, err := ipfs.Block().Get(context.Background(), path.New("QmU9D6ENLxzZojQVxiM3PUns8992crpUdzKLG5bSmEFNuu"))
	check(err)

	fmt.Println(reader)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
