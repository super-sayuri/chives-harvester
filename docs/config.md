## 配置文件结构

配置文件分为两个部分：config 和 keys。其中 config 以 yml 格式储存通用的配置，而 keys 以 json 格式储存一些加密信息。

配置文件格式可以在 `config/local.yml` 里查看。加密信息的储存额外说明。

## config.yml

配置文件，其中包括：

- common: 通用配置
  - lang: 使用语言，目前只支持 `zh`
  - config_type: 配置类型，目前仅支持 `file`，即从文件读取
  - name: 该实例的名字

- service: 服务器相关配置，由于现在使用 rr 来读取 telegram 的信息，所以这里暂时不用
  - port: 端口号
  - gin_mode: gin mode，具体参看：https://github.com/gin-gonic/gin/blob/master/mode.go

- template: 对话模板的设置
  - basepath: 对话模板所在文件路径，现在所有的配置在 `assets/tpl` 文件夹内

- tgbot: telegram 机器人相关配置
  - owner: 你的 telegram id，会到 keys 里寻找同名字段
  - token: telegram 机器人的 token， 具体可参看 https://core.telegram.org/bots/api#authorizing-your-bot ，会到 keys 里寻找同名字段
  - call_gap: 操作时间间隔，单位为秒

- cron: 定时任务相关配置
  - crypto: 加密货币定时任务 cron

- log: 日志相关配置
  - level: 日志级别
  - output: 日志输出方法，目前只支持默认输出和 `FILE`，即输出到文件，除了 FILE 都是默认输出
  - format: 日志输出格式，目前只支持默认格式和 `JSON`，除了 JSON 格式现在都是默认格式
  - path: 日志文件路径，仅在输出方法为 `FILE` 时有效

- redis: redis 配置
  - nodes: 节点的地址，由于项目不大现在只支持单节点
  - password: 密码，会到 keys 里寻找同名字段
 
- crypto: 获得加密货币所需要的设置
  - app: 平台名称
  - currency: 输出显示货币，如 USD
  - token: 平台所对应的 token，会到 keys 里寻找同名字段

## keys.json

用于储存加密信息，加密信息为K-V对，其中K在 config 文件中设置。

目前只包括:
- tgbot.owner
- tgbot.token
- redis.password
- crypto.token
