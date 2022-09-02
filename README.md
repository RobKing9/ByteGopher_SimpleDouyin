# ByteGopher_SimpleDouyin

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

## 目录

- [项目介绍](#项目介绍)
- [上手指南](#上手指南)
    - [开发前的配置要求](#开发前的配置要求)
    - [安装](#安装)
- [文件目录说明](#文件目录说明)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)

## 项目介绍

项目启动的首先会连接mysql数据库，之后启动路由监听端口。

### 用户注册接口

服务器获取用户输入的用户名和密码，先通过用户名在数据库中查找用户，如果已经存在的话返回信息用户名已存在，然后进行密码的验证，密码必须超过六位，为了信息安全，对用户的密码进行加密，使用的是`golang.org/x/crypto/bcrypt`包生成加密密码，之后生成随机id，将用户信息存放到数据库中。同时给用户发送token，使用的是JWT认证，存放的信息有用户的id

JWT认证，jwt主要包括三个部分

- Header：描述jwt的元数据，主要包括签名使用的算法
- PayLoad：用来存放实际要传递的数据，包括token过期时间，签发人，主题等信息
- Signature：对前两部分的签名，防止数据被篡改，首先会指定一个密钥，使用Header中的签名算法

生成token之后用户可以直接登录而不需要进行登录信息的输入

### 用户登录接口

服务器获取用户输入的用户名和密码，通过用户名在数据库中查找用户，将数据库中的密码进行解密和输入的密码进行比对，比对成功之后返回token给用户。

### 用户信息接口

想要通过信息接口查看信息必需要通过中间件的认证，中间件实现的细节如下：获取用户的token，如果token为空则用户没有权限查看 信息，否则对用户的token进行解析生成用户的id，通过这个id可以在数据库中查找到该用户，然后将用户信息放入到上下文中，通过认证之后即可查看用户的信息。

## 上手指南


### 开发前的配置要求
go1.16以上


### 安装


```sh
git clone https://github.com/RobKing9/ByteGopher_SimpleDouyin.git
```

## 文件目录说明

```
.
|-- Dockerfile
|-- README.md
|-- controller
|   |-- comment.go
|   |-- favorite.go
|   |-- relation.go
|   |-- user.go
|   `-- video.go
|-- dao
|   |-- comment.go
|   |-- favorite.go
|   |-- mysql.go
|   |-- relation.go
|   |-- user.go
|   `-- video.go
|-- main.go
|-- middleware
|   `-- AuthMiddleware.go
|-- model
|   |-- comment.go
|   |-- common.go
|   |-- favorite.go
|   |-- follow.go
|   |-- user.go
|   `-- video.go
|-- router
|   `-- router.go
|-- tree.txt
`-- utils
    |-- convertVideoModelListToVideoList.go
    |-- deleteFile.go
    |-- format.go
    |-- jwt.go
    |-- qiniuUpload.go
    |-- randNum.go
    `-- timetool.go



```


## :yum: 技术栈

- [Gin](https://gin-gonic.com)
- JWT
- [Gorm](https://gorm.io)
- [MySQL](https://mysql.com)

## 贡献者

请阅读**README.md** 查阅为该项目做出贡献的开发者。

### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



## 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

## 作者

曾祥文  QQ邮箱：2768817839@qq.com

*您也可以在贡献者名单中参看所有参与该项目的开发者。*

## 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/RobKing9/ByteGopher_SimpleDouyin/blob/master/LICENSE.txt)

## 鸣谢

- [字节跳动后端青训营](https://youthcamp.bytedance.com/)

<!-- links -->
[your-project-path]:RobKing9/ByteGopher_SimpleDouyin
[contributors-shield]: https://img.shields.io/github/contributors/RobKing9/ByteGopher_SimpleDouyin.svg?style=flat-square
[contributors-url]: https://github.com/RobKing9/ByteGopher_SimpleDouyin/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/RobKing9/ByteGopher_SimpleDouyin.svg?style=flat-square
[forks-url]: https://github.com/RobKing9/ByteGopher_SimpleDouyin/network/members
[stars-shield]: https://img.shields.io/github/stars/RobKing9/ByteGopher_SimpleDouyin.svg?style=flat-square
[stars-url]: https://github.com/RobKing9/ByteGopher_SimpleDouyin/stargazers
[issues-shield]: https://img.shields.io/github/issues/RobKing9/ByteGopher_SimpleDouyin.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/RobKing9/ByteGopher_SimpleDouyin.svg
[license-shield]: https://img.shields.io/github/license/RobKing9/ByteGopher_SimpleDouyin.svg?style=flat-square
[license-url]: https://github.com/RobKing9/ByteGopher_SimpleDouyin/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/RobKing9