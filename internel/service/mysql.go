package service

import "develop/stac/zhaoshang_zhengquan/st_ha/sql_tools/pkg"

func RenderFile(from, to string, data map[string]interface{}) error {
	return pkg.RenderFile(from, to, data)
}
