package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:           "sql_tools",
	Long:          `stac mysql tools`,
	SilenceErrors: true,
	SilenceUsage:  true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func SetVersion(version string) {
	rootCmd.SetVersionTemplate(`{{printf "Version: %s" .Version}}`)
	rootCmd.Version = version
}

func init() {
	var (
		source, target string
	)
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "source file(ex: /mnt/mysql_node1/docker-compose.yml.tpl)")
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "target file(ex: /mnt/mysql_node1/docker-compose.yml)")
}
