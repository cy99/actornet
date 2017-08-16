# ActorNet
[![MIT licensed][11]][12] [![GoDoc][1]][2]

[1]: https://godoc.org/github.com/davyxu/actornet?status.svg
[2]: https://godoc.org/github.com/davyxu/actornet
[11]: https://img.shields.io/badge/license-MIT-blue.svg
[12]: LICENSE

基于Actor模式游戏服务器框架

# 目标

- 方便的编写跨服逻辑及强状态逻辑

- 异步/同步,同进程/跨进程的消息和逻辑处理使用同样的方式

- 灰度热更新, 精确到actor, 做到某玩家的模块更新

- 更健壮的服务器逻辑, 逻辑崩溃或异常时,可以使用多种策略替换原逻辑

# 开发进度
- [x] PID
- [x] Process
- [x] Mailbox
- [x] 同进程异步消息
- [x] 跨进程异步消息
- [x] 同进程同步消息
- [x] 跨进程同步消息
- [x] 同步消息使用Future
- [x] 网关
- [ ] 同步消息Timeout
- [ ] 序列化
- [ ] 热更新架构
- [ ] 消息响应便捷化
- [ ] 进程通信抽象
- [ ] 灰度更新架构

# 备注

感觉不错请star, 谢谢!

博客: http://www.cppblog.com/sunicdavy

知乎: http://www.zhihu.com/people/sunicdavy
