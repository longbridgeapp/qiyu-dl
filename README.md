# qiyu-dl
七鱼会话信息下载器

## Quickstart

_安装 [golang](https://go.dev/doc/install) 环境_

_下载 qiyu-dl_

```
go install github.com/longbridgeapp/qiyu-dl@latest
```

_设置环境变量(MacOS/Linux)_

```bash
export QIYU_APP_KEY="七鱼 APP KEY"                       
export QIYU_APP_SECRET="七鱼 APP SECRET"
export QIYU_HOST="七鱼 API 域名"
```

_设置环境变量(Windows)_

```powershell
setx QIYU_APP_KEY "七鱼 APP KEY"
setx QIYU_APP_SECRET "七鱼 APP SECRET"
setx QIYU_HOST "七鱼 API 域名"
```

_使用_
```
qiyu-dl -date 2023-06-19

```




