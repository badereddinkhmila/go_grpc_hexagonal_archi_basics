package cmd

import (
	"context"
	"fmt"
	"log"
	"payment-service/config"
	"payment-service/internal/adapters/db"
	"payment-service/internal/adapters/grpc"
	"payment-service/internal/adapters/rest"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func Serve() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Lunch server",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("[Cmd] starting application")
		dbAdapter, err := db.NewAdapter()
		if err != nil {
			log.Fatalf("Failed to connect to database. Error: %v", err)
		}
		group, _ := errgroup.WithContext(context.Background())
		group.SetLimit(2)
		group.Go(func() error {
			fmt.Println("[Cmd] starting Grpc Server")
			srv := grpc.NewGrpcAdapter(dbAdapter, config.APP().Port+10)
			srv.Run()
			return nil
		})
		group.Go(func() error {
			fmt.Println("[Cmd] starting Rest server")
			restSrv := rest.NewRestAdapter(rest.Routes(), config.APP().Port)
			restSrv.Run()
			return nil
		})

		if err := group.Wait(); err != nil {
			fmt.Printf("errgroup tasks ended up with an error: %v\n", err)
		}
		fmt.Println("[Cmd] started application")
	}

	return cmd
}
