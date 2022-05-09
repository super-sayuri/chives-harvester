# 小百合的韭菜收割机

提供各类金融产品数据咨询(目前只有BTC和ETH),自动分发到各个平台(目前只有telegram bot),以及各类告警(下个版本更新,如果下个版本没更新参考上半句)

## 执行

### build
```bash
./cmd/build.sh
```

### 运行
```bash
./bin/syr_btc_bot.sh -c {配置文件地址} -k {重要信息文件地址}
```

## 准备

1. 安装redis （略）
2. 写入路由信息
```
hset syr_api_conf api_ping PING_API
hset syr_api_conf api_prefix API_PREFIX
hset syr_api_conf api_meta META_API
hset syr_api_conf api_command COMMAND_API
hset syr_api_conf api_webhook WEBHOOK_API
```
3. 根据./env/conf.tpl.yml和keys.tpl.json的模版完成配置文件
4. GO!

## todo
* 支持在不同群发消息
* template和item的cache
* ~~用redis储存object信息太tm弱智了！我要换mongoDB!~~
* 后来想想我买的机器也不怎么好，还是换sqlite凑合吧
* 等我有钱了，我一定……
* 增加用户系统，以后会有管理网站的
---
* 同比环比警告
* 模拟账户

天灾再起!