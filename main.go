package main

import (
	"goParticiple/sego"
	"fmt"
)

func main() {
	// 载入词典
	var segmenter = sego.Segm
	segmenter.LoadDictionary("../data/word.dic")
	// 热加载词典
	//go func() {
	//	sego.HotLoadDictionary()
	//}()

	// 分词
	text := []byte("洗手间门")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	fmt.Println(sego.SegmentsToString(segments, true))
}
