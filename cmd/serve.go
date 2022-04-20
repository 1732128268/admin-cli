/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"admin-cli/global"
	"admin-cli/initialize"
	"admin-cli/serve"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//	初始化配置
		if err := initialize.InitConfig(); err != nil {
			panic(err)
		}
		//	初始化日志
		initialize.InitLog()
		//	连接Mysql数据库
		if err := initialize.InitMysql(); err != nil {
			panic(err)
		}
		//初始化管理员
		//if err := initialize.Admin(); err != nil {
		//	panic(err)
		//}
		//	初始化redis
		if global.Config.HttpConfig.OpenRedis {
			if err := initialize.InitRedis(); err != nil {
				panic(err)
			}
		}
		//	启动http服务
		srv, router := serve.StartHttp()
		//路由
		initialize.Router(router)
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		logrus.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logrus.Fatal("Server Shutdown:", err)
		}
		logrus.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
