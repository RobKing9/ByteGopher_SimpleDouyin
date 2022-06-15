# ByteGopher_SimpleDouyin

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

## 目录

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


## 技术栈

- [Gin](https://gin-gonic.com)
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