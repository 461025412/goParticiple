package tool

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
	"goParticiple/def"
	"strconv"

	)

var DictionaryLoad def.DictionaryLoad

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

//判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 生成随机数
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func WriteToFile(fileName string, data interface{}) {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		logs.Error("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		data, err := json.Marshal(data)
		if err != nil {
			logs.Error("json化失败")
		}

		data = append(data, '\n')
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt(data, n)

	}

}

func PrintStack(r interface{}) string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	result := fmt.Sprintf("r:%v==> %s\n", r, string(buf[:n]))
	return result
}

func GetString(i interface{}, d string) string {
	if i == nil {
		return d
	}
	switch i.(type) {
	case string:
		return i.(string)
	case float64:
		s2 := strconv.FormatFloat(i.(float64), 'E', -1, 64)
		return s2
	case []uint8:
		var s3 []byte
		s2 := i.([]uint8)
		for _, item := range s2 {
			s3 = append(s3, byte(item))
		}
		return string(s3)
	}
	return d
}

