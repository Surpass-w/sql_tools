package cmd

import (
	"develop/stac/zhaoshang_zhengquan/st_ha/sql_tools/internel/service"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen conf file",
}

var ymlCmd = &cobra.Command{
	Use:   "yml",
	Short: "gen docker compose yml file",
}

var cnfCmd = &cobra.Command{
	Use:   "cnf",
	Short: "gen mysql conf file",
}

func init() {
	var (
		mysqlDir      string
		serverId      int
		localAddress  string
		groupSeeds    string
		groupName     string
		port          int
		mysqlxPort    int
		containerName string
	)
	genCmd.AddCommand(ymlCmd, cnfCmd)
	ymlCmd.Flags().SortFlags = true
	cnfCmd.Flags().SortFlags = true
	ymlCmd.Flags().StringVarP(&mysqlDir, "mysql_dir", "d", "", "mysql_dir(ex: /mnt/mysql_node1/)")
	ymlCmd.Flags().StringVarP(&containerName, "container_name", "c", "mysql", "container_name(ex: mysql)")
	cnfCmd.Flags().IntVarP(&serverId, "server_id", "i", 1, "server_id(ex: 1)")
	cnfCmd.Flags().StringVarP(&localAddress, "local_address", "l", "", "local_address(ex: 192.168.152.21:33306)")
	cnfCmd.Flags().StringVarP(&groupSeeds, "group_seeds", "g", "", "group_seeds(ex: 192.168.152.21:33306,192.168.152.22:33306,192.168.152.22:33307)")
	cnfCmd.Flags().StringVarP(&groupName, "group_name", "n", "", "group_name(ex: 75223d08-ff89-4f70-966b-97edce050f17)")
	cnfCmd.Flags().IntVarP(&port, "port", "p", 3306, "port(ex: 3306)")
	cnfCmd.Flags().IntVarP(&mysqlxPort, "mysqlx-port", "x", 33060, "mysqlx-port(ex: 33060)")
	ymlCmd.RunE = func(cmd *cobra.Command, args []string) error {
		sourcePath, _ := cmd.Flags().GetString("source")
		targetPath, _ := cmd.Flags().GetString("target")
		if strings.HasSuffix(mysqlDir, "/") {
			mysqlDir = strings.TrimRight(mysqlDir, "/")
		}
		data := make(map[string]interface{})
		params := make(map[string]interface{})
		data["MYSQL_DIR"] = mysqlDir
		data["CONTAINER_NAME"] = containerName
		params["SOURCE_PATH"] = sourcePath
		params["TARGET_PATH"] = targetPath
		params["MYSQL_DIR"] = mysqlDir
		params["CONTAINER_NAME"] = containerName
		fmt.Println(params)
		err := service.RenderFile(sourcePath, targetPath, data)
		return err
	}
	cnfCmd.RunE = func(cmd *cobra.Command, args []string) error {
		sourcePath, _ := cmd.Flags().GetString("source")
		targetPath, _ := cmd.Flags().GetString("target")
		data := make(map[string]interface{})
		params := make(map[string]interface{})
		data["SERVER_ID"] = serverId
		data["LOCAL_ADDRESS"] = localAddress
		data["GROUP_SEEDS"] = groupSeeds
		data["GROUP_NAME"] = groupName
		data["PORT"] = port
		data["MYSQL_X_PORT"] = mysqlxPort
		params["SOURCE_PATH"] = sourcePath
		params["TARGET_PATH"] = targetPath
		params["SERVER_ID"] = serverId
		params["LOCAL_ADDRESS"] = localAddress
		params["GROUP_SEEDS"] = groupSeeds
		params["GROUP_NAME"] = groupName
		params["PORT"] = port
		params["MYSQL_X_PORT"] = mysqlxPort
		fmt.Println(params)
		err := service.RenderFile(sourcePath, targetPath, data)
		return err
	}
	rootCmd.AddCommand(genCmd)
}
