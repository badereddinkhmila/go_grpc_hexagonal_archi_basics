package cmd

import (
	"context"
	"fmt"
	"log"
	"product-service/config"
	"product-service/internal/adapters/db"
	"product-service/internal/adapters/grpc"
	"product-service/internal/adapters/rest"
	"product-service/internal/application/core/api"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const serviceId = "product_microservice_id"
const serviceName = "product_microservice"

func Serve() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Lunch server",
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("[Cmd] starting application")

		// Initialize database
		dbAdapter, err := db.NewPostgreSQLAdapter()
		if err != nil {
			log.Fatalf("Failed to connect to database. Error: %v", err)
		}
		fmt.Println("[Cmd] db started correctly")

		// Initialize application
		applicationAdapter := api.NewApplication(dbAdapter)
		fmt.Println("[Cmd] api started correctly")

		// Create errgroup with context
		group, ctx := errgroup.WithContext(context.Background())
		group.SetLimit(3)

		// Start gRPC server
		group.Go(func() error {
			fmt.Println("[Cmd] starting gRPC Server")
			srv := grpc.NewGrpcAdapter(applicationAdapter, config.APP().Port+100)
			if err := srv.Run(); err != nil {
				return fmt.Errorf("gRPC server failed: %v", err)
			}
			fmt.Println("[Cmd] gRPC server started correctly")
			return nil
		})

		// Start REST server
		group.Go(func() error {
			fmt.Println("[Cmd] starting REST server")
			restSrv := rest.NewRestAdapter(*applicationAdapter, config.APP().Port)
			if err := restSrv.InitServer(); err != nil {
				return fmt.Errorf("REST server failed: %v", err)
			}
			fmt.Println("[Cmd] REST server started correctly")
			return nil
		})

		// Register with Consul
		group.Go(func() error {
			client, err := consulapi.NewClient(consulapi.DefaultConfig())
			if err != nil {
				return fmt.Errorf("failed to create Consul client: %v", err)
			}

			restRegistration := &consulapi.AgentServiceRegistration{
				ID:      fmt.Sprintf("%s_%s", serviceId, "rest"),
				Name:    fmt.Sprintf("%s_%s", serviceName, "rest"),
				Port:    config.APP().Port,
				Address: config.APP().Host,
				Tags:    []string{"products", "categories", "traefik.enable=true"},
				Check: &consulapi.AgentServiceCheck{
					HTTP:          fmt.Sprintf("http://%s:%d/check", config.APP().Host, config.APP().Port),
					Interval:      "10s",
					Timeout:       "30s",
					TLSSkipVerify: true,
					CheckID:       "checkalive",
				},
				Meta: map[string]string{
					fmt.Sprintf("traefik_http_routers_%s_rule", fmt.Sprintf("%s_%s", serviceName, "rest")):                      fmt.Sprintf("Host(`%s.rest.local`)", fmt.Sprintf("%s_%s", serviceName, "rest")),
					fmt.Sprintf("traefik_http_routers_%s_rule", fmt.Sprintf("%s_%s", serviceName, "rest")):                      fmt.Sprintf("PathPrefix(`/api`) && Host(`%s.rest.local`)", fmt.Sprintf("%s_%s", serviceName, "rest")),
					fmt.Sprintf("traefik_http_services_%s_loadbalancer_server_port", fmt.Sprintf("%s_%s", serviceName, "rest")): "80",
				},
			}

			grpcRegistration := &consulapi.AgentServiceRegistration{
				ID:      fmt.Sprintf("%s_%s", serviceId, "grpc"),
				Name:    fmt.Sprintf("%s_%s", serviceName, "grpc"),
				Port:    config.APP().Port + 100,
				Address: config.APP().Host,
				Tags:    []string{"grpc", "products", "categories", "traefik.enable=true"},
				Check: &consulapi.AgentServiceCheck{
					TCP:                            fmt.Sprintf("%s:%d", config.APP().Host, config.APP().Port+100),
					Interval:                       "10s",
					Timeout:                        "3s",
					DeregisterCriticalServiceAfter: "5m",
					CheckID:                        "grpc-tcp-check",
				},
				Meta: map[string]string{
					fmt.Sprintf("traefik_http_routers_%s_rule", fmt.Sprintf("%s_%s", serviceName, "grpc")):                        fmt.Sprintf("Host(`%s.grpc.local`)", fmt.Sprintf("%s_%s", serviceName, "grpc")),
					fmt.Sprintf("traefik_http_services_%s_loadbalancer_server_scheme", fmt.Sprintf("%s_%s", serviceName, "grpc")): "h2c",
					fmt.Sprintf("traefik_http_services_%s_loadbalancer_server_port", fmt.Sprintf("%s_%s", serviceName, "grpc")):   "80",
				},
			}

			log.Println("[Cmd] Registering REST server with Consul...")
			if err := client.Agent().ServiceRegister(restRegistration); err != nil {
				fmt.Printf("failed to register rest service: %v", err)
				return fmt.Errorf("failed to register rest service: %v", err)
			}
			log.Println("[Cmd] Registered REST server with Consul...")
			if err := client.Agent().ServiceRegister(grpcRegistration); err != nil {
				fmt.Printf("failed to register grpc service: %v", err)
				return fmt.Errorf("failed to register grpc service: %v", err)
			}

			log.Printf("Successfully registered service: %s:%v", config.APP().Host, config.APP().Port)

			// Deregister on shutdown
			go func() {
				<-ctx.Done()
				log.Println("Shutting down server...")
				if err := client.Agent().ServiceDeregister(serviceId); err != nil {
					log.Printf("Failed to deregister service: %v", err)
				}
			}()

			return nil
		})

		// Wait for all goroutines to complete
		if err := group.Wait(); err != nil {
			log.Fatalf("Application failed: %v", err)
		}
		fmt.Println("[Cmd] started application")
	}
	return cmd
}
