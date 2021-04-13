# go-template

application web tool for infra

## 配置注入

參數配置使用 viper 套件, 配置順序由高至低為 ([Viper uses the following precedence](https://github.com/spf13/viper#why-viper))

- explicit call to Set: 程式碼中使用 viper.Set() 配置設定
- flag: commaind line flag
- env: 使用環境變數, 必須為大寫
- config: 使用檔案配置 JSON, TOML, YAML, HCL, envfile
- key/value store: etcd或者consul
- default: 程式碼中使用 viper.SetDefault() 配置預設值


### 使用環境變數配置

宣告環境變數. 變數規範

- 前綴為 `ENV_` 
- 變數必須為大寫

範例:

```
# 配置 log 等級變數 `log.level: debug` 為 `export ENV_LOG_LEVEL=debug`
export ENV_LOG_LEVEL=debug
# 移除 log 時間搓記 `log.omittimekey: true` 為 `export ENV_LOG_OMITTIMEKEY=true`
export ENV_LOG_OMITTIMEKEY=true
go run main.go
```

### 使用檔案配置

參照 [confg.yml.sample](https://github.com/chiehting/go-template/tree/master/config/config.yml.sample), 依據範例改做修改.
