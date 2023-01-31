# assignment_bd

## 项目结构

##### 注：本次的小组项目的 go 的版本需要 1.18 及以上

### `api` 包

- 本来是准备放一些响应的结构体的（backend为后端，frontend为前），但是感觉太冗余了，就把所有的 response 结构体定义之后放在了 `model/response.go` 下
- 现已删除 qaq

### `config` 包

- 放一些项目的配置，不过现在里面只有 `DSN` 配置在 `config/default.go` 文件下
- 关于 `DSN` 的配置：因为大家本地数据库的用户名和密码都各不相同，所以建议大家把自己本地数据库的 `DSN` 加在这里面，方便大家以后对项目进行测试

### `consts` 包

- 下面定义的一些常量

### `controller` 包

- 根据 [青训营的api文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707524) 中的请求的二级目录进行分包
- ![image-20230131205557943](http://img.panker916.space/image-20230131205557943.png)

- 里面不同文件中的不同方法就对应着 api文档 中的不同的 api接口，大家可以去实现（其中 `user.go` 下的各个功能已经实现完毕（主要是登录功能），大家可以测试一下）

### `global` 包

- 里面的 `global/global.go` 文件下定义中项目所需要的一些全局变量

### `middleware` 包

- 定义一些项目中可能会用到的中间件

### `model` 包

- 将项目中可能会使用到的对象抽象成结构体
- ![image-20230131210031616](http://img.panker916.space/image-20230131210031616.png)

- 需要注意的是 `model/user.go` 里面的 `type User struct` 是对应着数据库中的结构体，而里面的 `type UserInfo struct` 则是因为部分 api 的 response 需要而创建出来的结构体， `UserInfo` 的对象不能直接从数据库中查出，需要通过组装字段来创建 `UserInfo` 对象。（ps：`model/video.go` 中的结构体同理）
-  `model/response.go` 定义了项目中所有可能用到的 `response` ，并且注释上有其对应的接口（具体可看代码注释）

### `service` 包

- 各种业务逻辑的实现（尽量与 `controller` 层对应？要不感觉太乱了）

### `utils` 包

- 包含了项目需要用到的工具

---

#### ps：大家写代码的时候尽量加点注释，并且尽量注意代码一下格式（根据代码的逻辑分片中间适当加点空行啥的），这样方便其它同学理解项目

---

## git示例操作

```
git branch -m main devlop
git fetch origin
git branch -u origin/devlop devlop
git remote set-head origin -a
```
