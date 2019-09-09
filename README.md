## 基于 https://github.com/huichen/sego 这个项目的修改。

### 添加了 字典热更新操作 定期（默认60秒，可以修改配置文件里的参数来控制监测时间）检查热词文件是否被更新。对于词典加了读写锁，确保不会在更新词典数据时导致服务不可用。
### 添加了 切割深度的功能 默认深度为2 可以通过配置文件设置切割深度
### 开启server服务，新增接口


###配置文件参数
[Log] // 日志设置
log_path=goParticiple.log // 存放日志地址
log_level=debug //  日志级别
[WordFilePath]  // 词典位置
word_file_path=../data/word.dic
[TermDepth] // 切割深度
term_depth=2
[UpdateInterval] 更新间隔
update_interval=60
[HttpPort] 服务监听端口
http_port=8080

## 用法 

### 进入server目录，运行 go build -mod=vendor ，启动服务 ./server

### 服务域名/getFormatTerms 例如 172.16.4.24:8080/getFormatTerms 返回格式化的分词信息

    [
        {
            "Term": "大理石电视背景墙",
            "Pos": "n",
            "Frequency": 7
        },
        {
            "Term": "大理石",
            "Pos": "n",
            "Frequency": 7
        },
        {
            "Term": "大",
            "Pos": "n",
            "Frequency": 7
        },
        {
            "Term": "理石",
            "Pos": "n",
            "Frequency": 7
        },
        {
            "Term": "电视背景墙",
            "Pos": "n",
            "Frequency": 7
        },
        {
            "Term": "电视",
            "Pos": "n",
            "Frequency": 7
        },
        {
            "Term": "背景墙",
            "Pos": "n",
            "Frequency": 7
        }
    ]


### 服务域名/getTerms  例如 172.16.4.24:8080/getTerms 返回分词信息

    [
        "大理石电视背景墙",
        "大理石",
        "大",
        "理石",
        "电视背景墙",
        "电视",
        "背景墙"
    ]
