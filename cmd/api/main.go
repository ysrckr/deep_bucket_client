package main

import (
	"context"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ysrckr/deep_bucket_client/internal/server"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	server.StartServer(ctx)

}
