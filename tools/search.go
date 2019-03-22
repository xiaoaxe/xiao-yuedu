// 搜索
// author: baoqiang
// time: 2019/3/22 下午12:59
package tools

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"fmt"
	"log"
	"github.com/githubao/xiao-yuedu/helper"
)

var searcher engine.Engine

func InitSearch() engine.Engine {
	searcher = engine.Engine{}
	searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries:   helper.GetDictionaryFile(),
		UsePersistentStorage:    true,
		PersistentStorageFolder: helper.GetConfig().Wukong.IndexPath,
	})

	// import books
	//books := dal.FindBookAll()
	//for _, book := range books {
	//	AddDocument(book.Id, book.Name, book.Content)
	//}

	// flush
	searcher.FlushIndex()

	return searcher
}

func AddDocument(bid int64, name, body string) {
	content := fmt.Sprintf("%s: %s", name, body)
	log.Printf("Add search content: %d -> %s", bid, content)
	searcher.IndexDocument(uint64(bid), types.DocumentIndexData{Content: content}, false)
}

func SearchDocument(text string) types.SearchResponse {
	return searcher.Search(types.SearchRequest{Text: text})
}
