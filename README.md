# 极简版抖音服务端

## 一、介绍

- 1.基于RPC框架**Kitex**、HTTP框架**Gin**、ORM框架**GORM**的极简版抖音服务端项目

- 2.代码采用api层、service层、dal层三层结构

- 3.使用Kitex构建RPC微服务，Gin构建HTTP服务

- 4.GORM操作MySQL数据库，防止SQL注入，使用事务保证数据一致性，完整性

- 5.使用**ETCD**进行服务注册、服务发现

- 6.使用**MySQL**数据库进行数据存储，并建立索引

- 7.使用redis进行数据缓存，增大并发量

- 8.使用**OSS**进行视频对象存储，分片上传视频

- 9.使用**JWT**鉴权，**MD5**密码加密，**ffmpeg**获取视频第一帧当作视频封面

    

![](/public/技术框架图.png)

## 二、架构图



![](/public/项目架构图.png)



## 三、数据库ER图



## 四、文件目录结构

| idl       | proto 接口定义文件  |                                                            |
| --------- | ------------------- | ---------------------------------------------------------- |
| kitex_gen | Kitex自动生成的代码 |                                                            |
| pkg       | bound               | Kitex Transport Pipeline-Bound 拓展                        |
|           | constants           | 常量                                                       |
|           | errno               | 错误码                                                     |
|           | jwt                 | jwt认证                                                    |
|           | middleware          | Kitex的中间件                                              |
|           | oss                 | 对象存储                                                   |
|           | sign                | 个性签名                                                   |
| dal       | cache               | 封装了点赞的缓存逻辑                                       |
|           | db                  | 封装了其他数据库访问逻辑                                   |
|           | pack                | 数据打包/处理                                              |
|           | init.go             | 数据库初始化                                               |
| cmd       | api                 | 处理外部http请求，通过rpc客户端发送rpc请求                 |
|           | comment             | 评论微服务，支持查看评论、新增评论和删除评论等功能         |
|           | favorite            | 点赞微服务，支持点赞、取消点赞，个人主页显示点赞列表等功能 |
|           | feed                | 视频流微服务，支持获取视频流功能                           |
|           | message             | 消息微服务，支持查看好友关系列表、发送消息                 |
|           | publish             | 发布视频微服务，支持视频上传、视频列表等功能               |
|           | relation            | 关系微服务，支持关注、取消关注、查看关注和粉丝列表等功能   |
|           | user                | 用户微服务，支持用户注册、用户登录、用户信息等功能         |

## 五、项目运行

#### 1. 更改配置

```Plain
pkg/constants/constant.go
```

#### 2. 运行基础依赖

```Plain
docker-compose up
```

#### 3. 运行comment微服务

```Plain
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

#### 4. 运行favorite微服务

```Plain
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

#### 5. 运行feed微服务

```Plain
cd cmd/feed
sh build.sh
sh output/bootstrap.sh
```

#### 6. 运行message微服务

```Bash
cd cmd/message
sh build.sh
sh output/bootstrap.sh
```

#### 7. 运行publish微服务

```Plain
cd cmd/publish
sh build.sh
sh output/bootstrap.sh
```

#### 8. **运行relation微服务**

```Plain
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

#### 9. 运行user微服务

```Plain
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

#### 10. 运行api微服务

```Bash
cd cmd/api
chmod +x run.sh
./run.sh
```
