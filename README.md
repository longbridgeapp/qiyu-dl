# qiyu-dl
qiyu session message downloader, 七鱼会话信息下载器

## Quickstart

_1. 安装 [golang](https://go.dev/doc/install) 环境_

_2. 下载 qiyu-dl_

```
go install github.com/longbridgeapp/qiyu-dl@latest
```

_3. 设置环境变量_

```bash
## MacOS/LINUX
export QIYU_APP_KEY="七鱼 APP KEY"                       
export QIYU_APP_SECRET="七鱼 APP SECRET"
export QIYU_HOST="七鱼 API 域名"
```

```powershell
## Windows 
setx QIYU_APP_KEY "七鱼 APP KEY"
setx QIYU_APP_SECRET "七鱼 APP SECRET"
setx QIYU_HOST "七鱼 API 域名"
```

_使用 qiyu-dl 下载某天信息_
```
qiyu-dl -date 2023-06-19

```




