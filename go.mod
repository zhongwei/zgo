module github.com/zhongwei/zgo

require (
	github.com/go-xorm/xorm v0.7.1
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/mitchellh/go-homedir v1.0.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181212231659-93c0bb5c8393
	golang.org/x/net => github.com/golang/net v0.0.0-20181213202711-891ebc4b82d6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181213200352-4d1cda033e06
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181214171254-3c39ce7b6105
)
