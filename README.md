# apiGo-template

application web tool for infra

## Prerequisites

* [MySQL 8.0.24](https://hub.docker.com/_/mysql?tab=tags&page=1&ordering=last_updated&name=8.0.24)
* [Redis 5.0.8](https://hub.docker.com/_/redis?tab=tags&page=1&ordering=last_updated&name=5.0.8-alpine3.11)

## 配置注入

參數配置使用 viper 套件, 配置順序由高至低為 ([Viper uses the following precedence](https://github.com/spf13/viper#why-viper))

* explicit call to Set: 程式碼中使用 viper.Set() 配置設定
* flag: commaind line flag
* env: 使用環境變數, 必須為大寫
* config: 使用檔案配置 JSON, TOML, YAML, HCL, envfile
* key/value store: etcd或者consul
* default: 程式碼中使用 viper.SetDefault() 配置預設值

### 使用環境變數配置

宣告環境變數. 變數規範:

* 前綴為 `ENV_`
* 變數必須為大寫

範例:

```bash
# 配置 log 等級變數 `log.level: debug` 為 `export ENV_LOG_LEVEL=debug`
export ENV_LOG_LEVEL=debug
# 移除 log 時間搓記 `log.omittimekey: true` 為 `export ENV_LOG_OMITTIMEKEY=true`
export ENV_LOG_OMITTIMEKEY=true
go run main.go
```

### 使用檔案配置

參照 [confg.yml.sample](./config/config.yml.sample), 依據範例改做修改.

## 啟動服務

```bash
$ cp ./config/config.yml.sample ./config/config.yml
$ go run main.go
2021-08-11T01:31:20Z    info    log/log.go:66   log.level:debug
Applied 2 migrations
2021-08-11T01:31:20Z    info    apiGo-template/main.go:15       http server listening port :80
```

啟動後可以開啟 [http://localhost](http://localhost) 開啟測試頁面.
