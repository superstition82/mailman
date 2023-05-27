package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"probemail/server"
	_config "probemail/server/config"
)

const (
	greetingBanner = `
 ______     __    __     ______     __     __        
/\  ___\   /\ "-./  \   /\  __ \   /\ \   /\ \       
\ \  __\   \ \ \-./\ \  \ \  __ \  \ \ \  \ \ \____  
 \ \_____\  \ \_\ \ \_\  \ \_\ \_\  \ \_\  \ \_____\ 
  \/_____/   \/_/  \/_/   \/_/\/_/   \/_/   \/_____/ 
                                                     
`
)

var (
	config *_config.Config
	mode   string
	port   int
	data   string

	rootCmd = &cobra.Command{
		Use:   "probemail",
		Short: "An self-hosted email manager",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			s, err := server.New(ctx, config)
			if err != nil {
				cancel()
				fmt.Printf("Failed to create server, error: %+v\n", err)
				return
			}

			c := make(chan os.Signal, 1)
			// Trigger graceful shutdown on SIGINT or SIGTERM.
			// The default signal sent by the `kill` command is SIGTERM,
			// which is taken as the graceful shutdown signal for many systems, eg., Kubernetes, Gunicorn.
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				sig := <-c
				fmt.Printf("%s received.\n", sig.String())
				s.Shutdown(ctx)
				cancel()
			}()

			println(greetingBanner)
			fmt.Printf("Version %s has started at :%d\n", config.Version, config.Port)
			if err := s.Start(ctx); err != nil {
				if err != http.ErrServerClosed {
					fmt.Printf("failed to start server, error: %+v\n", err)
					cancel()
				}
			}

			// Wait for CTRL-C.
			<-ctx.Done()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "demo", `mode of server, can be "prod" or "dev" or "demo"`)
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "port of server")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "data directory")

	err := viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))
	if err != nil {
		panic(err)
	}

	viper.SetDefault("mode", "dev")
	viper.SetDefault("port", 8081)
	viper.SetEnvPrefix("memos")

	// setupCmd.Flags().String(setupCmdFlagHostUsername, "", "Owner username")
	// setupCmd.Flags().String(setupCmdFlagHostPassword, "", "Owner password")

	// rootCmd.AddCommand(setupCmd)
}

func initConfig() {
	viper.AutomaticEnv()
	var err error

	config, err = _config.GetConfig()
	if err != nil {
		fmt.Printf("failed to get config, error: %+v\n", err)
		return
	}

	println("---")
	println("Server config")
	println("dsn:", config.DSN)
	println("port:", config.Port)
	println("mode:", config.Mode)
	println("version:", config.Version)
	println("---")
}

const (
	setupCmdFlagHostUsername = "host-username"
	setupCmdFlagHostPassword = "host-password"
)
