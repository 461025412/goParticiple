package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"io"
		"net/http"
	"runtime"
	"goParticiple/sego"
		"goParticiple/bootstrap"
	"github.com/astaxie/beego/logs"
	"goParticiple/tool"
	)

var (
	host      = flag.String("host", "", "HTTP服务器主机名")
	configPath      = flag.String("configPath", "goParticiple.conf", "配置文件地址，默认为 goParticiple.conf")
)

type Segment struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
}
// 只返回切出来的词
func JsonTermRpcServer(w http.ResponseWriter, req *http.Request) {
	// 得到要分词的文本
	text := req.URL.Query().Get("keyword")
	if text == "" {
		text = req.PostFormValue("keyword")
	}

	if text == "" {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "")
	}

	returnData := sego.GetWords(text)
	response, _ := json.Marshal(returnData)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(response))
}

// 返回切出来的结构化的词
func JsonFormatTermRpcServer(w http.ResponseWriter, req *http.Request) {
	// 得到要分词的文本
	text := req.URL.Query().Get("keyword")
	if text == "" {
		text = req.PostFormValue("keyword")
	}

	// 空的时候返回空s
	if text == "" {
		response, _ := json.Marshal([]int{})
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(response))
	}

	// 分词
	segments := sego.Segm.Segment([]byte(text))
	participleTerms := sego.SegmentsToFormatTerm(segments)

	response, _ := json.Marshal(participleTerms)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(response))
}

func main() {
	flag.Parse()

	// 将线程数设置为CPU数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//初始化服务
	err := bootstrap.InitService(*configPath)
	if err != nil {
		panic(err)
	}

	// 热加载词典
	go func() {
		sego.HotLoadDictionary()
	}()

	http.HandleFunc("/getTerms", errWrapper(JsonTermRpcServer))
	http.HandleFunc("/getFormatTerms", errWrapper(JsonFormatTermRpcServer))
	logs.Debug("开始监听网络端口:%s", bootstrap.ConfigData.HttpPort)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *host, bootstrap.ConfigData.HttpPort), nil)
}

/**
http 请求统一错误处理
*/
func errWrapper(handler func(writer http.ResponseWriter,
	request *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				runtimeErr := tool.PrintStack(r)
				res := tool.NewResponse(writer, request)
				logs.Error(fmt.Sprintf("服务器处理请求失败 panic: %v\n runtime error:%v", r, runtimeErr))
				res.ReturnError(http.StatusBadRequest, 200002, "服务器处理请求失败")
			}
		}()
		handler(writer, request)
	}
}
