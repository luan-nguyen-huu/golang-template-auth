package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/luan-nguyen-huu/Adam/configs"
	"github.com/luan-nguyen-huu/Adam/internal/routers"
)

func main() {
	cfg, _ := configs.Load()

    router := routers.RegisterMainRoutes(cfg)

    addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
    log.Printf("Server running on %s", addr)

    http.ListenAndServe(addr, router)
}
