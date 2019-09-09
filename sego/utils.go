package sego

import (
	"bytes"
	"fmt"
	"goParticiple/def"
	"goParticiple/bootstrap"
	"strings"
	)

// 输出分词结果为字符串
//
// 有两种输出模式，以"中华人民共和国"为例
//
//  普通模式（searchMode=false）输出一个分词"中华人民共和国/ns "
//  搜索模式（searchMode=true） 输出普通模式的再细致切分：
//      "中华/nz 人民/n 共和/nz 共和国/ns 人民共和国/nt 中华人民共和国/ns "
//
// 搜索模式主要用于给搜索引擎提供尽可能多的关键字，详情请见Token结构体的注释。
func SegmentsToStringUnique(segs []Segment, searchMode bool) (output string) {
	if searchMode {
		var uniqueMap = map[string]bool{}
		for _, seg := range segs {
			//output += tokenToString(seg.token)
			tokenStr := strings.TrimSpace(tokenToStringUnique(seg.token, bootstrap.ConfigData.TermDepth, uniqueMap))
			if len(tokenStr) > 0 {
				output += tokenStr
			}
		}
	} else {
		for _, seg := range segs {
			output += fmt.Sprintf(
				"%s/%s ", textSliceToString(seg.token.text), seg.token.pos)
		}
	}

	return
}

func SegmentsToString(segs []Segment, searchMode bool) (output string) {
	if searchMode {
		for _, seg := range segs {
			//output += tokenToString(seg.token)
			tokenStr := strings.TrimSpace(tokenToString(seg.token, bootstrap.ConfigData.TermDepth))
			if len(tokenStr) > 0 {
				output += tokenStr
			}
		}
	} else {
		for _, seg := range segs {
			output += fmt.Sprintf(
				"%s/%s ", textSliceToString(seg.token.text), seg.token.pos)
		}
	}

	return
}

func GetUniqueWords(text string) []string {
	// 分词
	segments :=  Segm.Segment([]byte(text))
	participleStr := SegmentsToStringUnique(segments, true)

	participleStr = strings.Trim(participleStr, ",")
	participleSlice := strings.Split(participleStr, ",")

	if len(participleSlice) == 0 {
		participleSlice[0] = text
	}

	returnData := []string{}
	for _, item:= range participleSlice {
		item := strings.TrimSpace(item)
		if len(item) > 0 {
			returnData = append(returnData, item)
		}
	}

	return returnData
}

func GetWords(text string) []string {
	// 分词
	segments :=  Segm.Segment([]byte(text))
	participleStr := SegmentsToString(segments, true)

	participleStr = strings.Trim(participleStr, ",")
	participleSlice := strings.Split(participleStr, ",")

	if len(participleSlice) == 0 {
		participleSlice[0] = text
	}

	returnData := []string{}
	for _, item:= range participleSlice {
		item := strings.TrimSpace(item)
		if len(item) > 0 {
			returnData = append(returnData, item)
		}
	}

	return returnData
}

func tokenToStringUnique(token *Token, depth int, uniqueMap map[string]bool) (output string) {
	if depth >= 0 {
		hasOnlyTerminalToken := true
		for _, s := range token.segments {
			if len(s.token.segments) > 1 {
				hasOnlyTerminalToken = false
			}
		}

		if !hasOnlyTerminalToken {
			depth--
			for _, s := range token.segments {
				if s != nil {
					tokenStr := tokenToStringUnique(s.token, depth, uniqueMap)
					if len(tokenStr) > 0 {
						output += tokenStr
					}
				}
			}
		}

		output = getTextStrUnique(GetTextStr(token.text), uniqueMap, output)
	}
	return
}

func tokenToString(token *Token, depth int) (output string) {
	if depth >= 0 {
		hasOnlyTerminalToken := true
		for _, s := range token.segments {
			if len(s.token.segments) > 1 {
				hasOnlyTerminalToken = false
			}
		}

		if !hasOnlyTerminalToken {
			depth--
			for _, s := range token.segments {
				if s != nil {
					tokenStr := tokenToString(s.token, depth)
					if len(tokenStr) > 0 {
						output += tokenStr
					}
				}
			}
		}

		output = getTextStr(GetTextStr(token.text), output)
	}
	return
}

// 返回格式化的term信息
func SegmentsToFormatTerm(segs []Segment) (output []def.FormatTerm) {
	var uniqueMap = map[string]bool{}
	for _, seg := range segs {
		output = append(output, tokenToFormatTerm(seg.token, bootstrap.ConfigData.TermDepth, uniqueMap)...)
	}
	return
}

// 获取结构化的词汇
func getFormatTerm(textStr string, token *Token, uniqueMap map[string]bool, output []def.FormatTerm) []def.FormatTerm {
	if len(textStr) > 0 {
		if _,ok := uniqueMap[textStr]; !ok {
			formatTerm := def.FormatTerm{}
			formatTerm.Term = textStr
			formatTerm.Pos = token.pos
			formatTerm.Frequency = token.frequency
			output = append(output, formatTerm)
			uniqueMap[textStr] = true
		}
	}

	return output
}

func tokenToFormatTerm(token *Token, depth int, uniqueMap map[string]bool) (output []def.FormatTerm) {
	if depth > 0 {
		hasOnlyTerminalToken := true
		for _, s := range token.segments {
			if len(s.token.segments) > 1 {
				hasOnlyTerminalToken = false
			}
		}

		if !hasOnlyTerminalToken {
			depth--
			for _, s := range token.segments {
				if s != nil {
					output = append(output, tokenToFormatTerm(s.token, depth, uniqueMap)...)
				}
			}
		}
		output = getFormatTerm(GetTextStr(token.text), token, uniqueMap, output)
	}
	return
}

func GetTextStr(text []Text) string {
	var str = ""
	if len(text) > 0 {
		for _, item := range text {
			str += strings.TrimSpace(fmt.Sprintf("%s", item))
		}
	}

	return str
}

func getTextStrUnique(textStr string, uniqueMap map[string]bool, output string) string {
	textStr = strings.TrimSpace(textStr)
	if len(textStr) > 0 {
		if _,ok := uniqueMap[textStr]; !ok {
			uniqueMap[textStr] = true
			output += "," + textStr
		}
	}

	return output
}

func getTextStr(textStr string, output string) string {
	textStr = strings.TrimSpace(textStr)
	if len(textStr) > 0 {
		output += "," + textStr
	}

	return output
}


// 输出分词结果到一个字符串slice
//
// 有两种输出模式，以"中华人民共和国"为例
//
//  普通模式（searchMode=false）输出一个分词"[中华人民共和国]"
//  搜索模式（searchMode=true） 输出普通模式的再细致切分：
//      "[中华 人民 共和 共和国 人民共和国 中华人民共和国]"
//
// 搜索模式主要用于给搜索引擎提供尽可能多的关键字，详情请见Token结构体的注释。

func SegmentsToSlice(segs []Segment, searchMode bool) (output []string) {
	if searchMode {
		for _, seg := range segs {
			output = append(output, tokenToSlice(seg.token)...)
		}
	} else {
		for _, seg := range segs {
			output = append(output, seg.token.Text())
		}
	}
	return
}

func tokenToSlice(token *Token) (output []string) {
	hasOnlyTerminalToken := true
	for _, s := range token.segments {
		if len(s.token.segments) > 1 {
			hasOnlyTerminalToken = false
		}
	}
	if !hasOnlyTerminalToken {
		for _, s := range token.segments {
			output = append(output, tokenToSlice(s.token)...)
		}
	}
	output = append(output, textSliceToString(token.text))
	return output
}

// 将多个字元拼接一个字符串输出
func textSliceToString(text []Text) string {
	return Join(text)
}

func Join(a []Text) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return string(a[0])
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return string(a[0]) + string(a[1])
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return string(a[0]) + string(a[1]) + string(a[2])
	}
	n := 0
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], s)
	}
	return string(b)
}

// 返回多个字元的字节总长度
func textSliceByteLength(text []Text) (length int) {
	for _, word := range text {
		length += len(word)
	}
	return
}

func textSliceToBytes(text []Text) []byte {
	var buf bytes.Buffer
	for _, word := range text {
		buf.Write(word)
	}
	return buf.Bytes()
}
