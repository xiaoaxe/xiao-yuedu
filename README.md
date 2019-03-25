# xiao-yuedu
构建一个下载电子书的网站


### feature
- [x] item & list
- [x] search
- [x] category
- [x] download
- [x] hot
- [ ] save
- [ ] tag
- [ ] css分页和样式

### 目前支持的方法列表
1. [book](http://localhost:8000/books/2)
1. [page](http://localhost:8000/pages/2)
1. [search](http://localhost:8000/search?text=go&p=2)
1. [download](http://localhost:8000/download?book_id=2)
1. [hot](http://localhost:8000/hot?p=2)

### 引用的类库
1. web框架：github.com/gin-gonic/gin
1. sqlite3存储：github.com/mattn/go-sqlite3
1. yaml配置解析：github.com/go-yaml/yaml 
1. 数据库orm：github.com/jinzhu/gorm
1. 全局搜索：github.com/huichen/wukong

### 规范
1. 代码尽量使用全称呼不使用缩写，书名命名只有英文中文和下划线
2. 

### 依赖关系图
- depth
    - download: go get github.com/KyleBanks/depth/cmd/depth
    - run: depth github.com/githubao/xiao-yuedu > dep.txt
- godepgraph
    - download: go get github.com/kisielk/godepgraph
    - run: godepgraph -s github.com/githubao/xiao-yuedu > dep.dot
    - run: dot dep.dot -T png -o dep.png

### 代码质量检查
- gometalinter
    - download: go get github.com/alecthomas/gometalinter
    - install: gometalinter --install --update
    - run: gometalinter ./... -e 'should have comment' -e 'should be of the form' -e 'return value not checked' -e 'Errors unhandled' --skip vendor --deadline 600s > warn.txt
- golangci-lint
    - download: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
    - run: golangci-lint run
