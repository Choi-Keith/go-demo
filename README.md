# go-demo
go-demo主要是用来练习开发中常用的一些库


## 技术
- web框架：gin
- orm: gorm
- database: mysql
- logger：logrus or zap

## 内容
- 用户登录注册：用户可以采用oauth2进行登录或注册，或则采用用户名和密码的方式进行登录注册   
- jwt鉴权：目前只开发了token的生成和认证, TODO: 需要使用redis存储，并且校验过期时间和refresh_token

