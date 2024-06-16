/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dns-host/client/service"
	"dns-host/pkg/color"
	"fmt"

	"github.com/spf13/cobra"
)

// hostRestartCmd represents the hostRestart command
var hostRestartCmd = &cobra.Command{
	Use:   "hostRestart",
	Short: "Перезапускает службу systemd-hostnamed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ok, err := service.RestartHost()
		if err != nil && !ok {
			print(fmt.Sprintf("Сервис hostnamed не перезапущен, ошибка %s", err), color.CRed)
			return
		}
		print("Сервис hostnamed перезапущен", color.CGreen)
	},
}

func init() {
	rootCmd.AddCommand(hostRestartCmd)
}