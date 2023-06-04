package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pocketmail/server"
	"pocketmail/server/config"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg *config.Config

	mode string
	port int
	data string

	rootCmd = &cobra.Command{
		Use:   "pocketmail",
		Short: `이메일 검증, 이메일 템플릿 관리, 발송 기능을 지원하는 웹 서비스입니다`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			s, err := server.NewServer(ctx, cfg)
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

			fmt.Printf("Version %s has started at :%d\n", cfg.Version, cfg.Port)
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

			db := db.NewDB(cfg)
			if err := db.Open(ctx); err != nil {
				fmt.Printf("failed to open db, error: %+v\n", err)
				return
			}

			store := store.New(db.DBInstance, cfg)
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

	viper.SetDefault("mode", "demo")
	viper.SetDefault("port", 8081)
	viper.SetEnvPrefix("memos")

	// setupCmd.Flags().String(setupCmdFlagHostUsername, "", "Owner username")
	// setupCmd.Flags().String(setupCmdFlagHostPassword, "", "Owner password")

	// rootCmd.AddCommand(setupCmd)
}

func initConfig() {
	viper.AutomaticEnv()
	var err error
	cfg, err = config.Getconfig()
	if err != nil {
		fmt.Printf("failed to get cfg, error: %+v\n", err)
		return
	}

	println("---")
	println("Server config")
	println("dsn:", cfg.DSN)
	println("port:", cfg.Port)
	println("mode:", cfg.Mode)
	println("version:", cfg.Version)
	println("---")
}

// const (
// 	setupCmdFlagHostUsername = "host-username"
// 	setupCmdFlagHostPassword = "host-password"
// )
