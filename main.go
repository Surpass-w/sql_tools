package main

import (
	"develop/stac/zhaoshang_zhengquan/st_ha/sql_tools/cmd"
	"fmt"
)

var (
	Version   = "v1.0.0"
	CommitID  = ""
	BuildTime = ""
)

func main() {
	cmd.SetVersion(fmt.Sprintf("  %s\nCommitID:  %s\nBuildTime: %s\n"+
		"Author:    wangpeng@moresec.cn\n", Version, CommitID, BuildTime))
	cmd.Execute()
	return
}
