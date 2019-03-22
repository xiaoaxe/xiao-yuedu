# xiao-yuedu
构建一个下载电子书的网站


### feature
- [x] item & list
- [x] search
- [ ] hot
- [ ] save
- [ ] category
- [ ] download
- [ ] tag
- [ ] css分页和样式

### 目前支持的方法列表
1. [book](http://localhost:8000/books/1)
1. [page](http://localhost:8000/pages/1)
1. [search](http://localhost:8000/search?text=围城&p=1)

### 引用的类库
1. web框架：github.com/gin-gonic/gin
1. sqlite3存储：github.com/mattn/go-sqlite3
1. yaml配置解析：github.com/go-yaml/yaml 
1. 数据库orm：github.com/jinzhu/gorm
1. 全局搜索：github.com/huichen/wukong
