package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"mails/server"
	_config "mails/server/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config *_config.Config
	mode   string
	port   int
	data   string

	rootCmd = &cobra.Command{
		Use:   "mailman",
		Short: `이메일 검증, 이메일 템플릿 관리, 발송 기능을 지원하는 웹 서비스입니다`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			s, err := server.NewServer(ctx, config)
			if err != nil {
				cancel()
				fmt.Printf("failed to create server, error: %+v\n", err)
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

	/**
	setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Make initial setup for memos",
		Run: func(cmd *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			hostUsername, err := cmd.Flags().GetString(setupCmdFlagHostUsername)
			if err != nil {
				fmt.Printf("failed to get owner username, error: %+v\n", err)
				return
			}

			hostPassword, err := cmd.Flags().GetString(setupCmdFlagHostPassword)
			if err != nil {
				fmt.Printf("failed to get owner password, error: %+v\n", err)
				return
			}

			db := db.NewDB(config)
			if err := db.Open(ctx); err != nil {
				fmt.Printf("failed to open db, error: %+v\n", err)
				return
			}

			store := store.New(db.DBInstance, config)
			if err := setup.Execute(ctx, store, hostUsername, hostPassword); err != nil {
				fmt.Printf("failed to setup, error: %+v\n", err)
				return
			}
		},
	}
	*/
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "prod", `mode of server, can be "prod" or "dev" or "demo"`)
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

	viper.SetDefault("mode", "prod")
	viper.SetDefault("port", 8081)
	viper.SetEnvPrefix("memos")

	// setupCmd.Flags().String(setupCmdFlagHostUsername, "", "Owner username")
	// setupCmd.Flags().String(setupCmdFlagHostPassword, "", "Owner password")

	// rootCmd.AddCommand(setupCmd)
}

func initConfig() {
	viper.AutomaticEnv()
	var err error
	config, err = _config.Getconfig()
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

// const (
// 	setupCmdFlagHostUsername = "host-username"
// 	setupCmdFlagHostPassword = "host-password"
// )
