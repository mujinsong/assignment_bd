# assignment_bd

### 项目结构介绍

| 名          | 用处                                                |
|------------|---------------------------------------------------|
| api        | 主要写请求回复的结构体(backend为后端，frontend为前)                |
| config     | 一些连接定义，如数据库连接，redis连接，CDN连接                       |
| consts     | 常量                                                |
| controller | 承上启下，路由里调用该层函数，controller层负责接受请求，然后调用service层函数处理 |
| dao        | 数据层，定义数据库结构体                                      |
| global     | 全局变量，一般放像Gorm.DB                                  |
| service    | 核心处理函数，一般被controller调用                            |
| middleware | 中间件                                               |
| utils      | 工具函数，写一些不相关的                                      |

- 大作业

- git示例操作
  ```
  git branch -m main devlop
  git fetch origin
  git branch -u origin/devlop devlop
  git remote set-head origin -a
  ```
