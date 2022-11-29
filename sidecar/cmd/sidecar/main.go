package main

import (
	"fmt"
	"sidecar/config"
	"sidecar/infra/db"
)

func main(){
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	if err := db.InitDB(db.URI(cfg.Database), cfg.IsLocal()); err != nil {
		panic(err)
	}
}

