package common

import (
	"fmt"
	"github.com/xiusin/pinecms/src/application/models"
	"strconv"
	"unicode/utf8"
)

// 处理文章列表信息数据. 补全一些cms生成
func HandleArtListInfo(list []map[string]string, titlelen int) {
	m := models.NewCategoryModel()
	for i, art := range list {
		catid, _ := strconv.Atoi(art["catid"])
		prefix := m.GetUrlPrefix(int64(catid))
		if art["type"] != "2" {
			art["caturl"] = fmt.Sprintf("/%s/", prefix)
			art["typeurl"] = art["caturl"]
		}
		id, _ := strconv.Atoi(art["id"])
		art["arcurl"] = fmt.Sprintf("/%s/%d.html", prefix, id)
		art["arturl"] = art["arcurl"]
		art["click"] = art["visit_count"]
		art["fulltitle"] = art["title"]
		if titlelen > 0 {
			if utf8.RuneCountInString(art["title"]) > titlelen {
				titleRune := []rune(art["title"])
				art["title"] = string(titleRune[:titlelen])
			}
		}
		list[i] = art
	}
}
