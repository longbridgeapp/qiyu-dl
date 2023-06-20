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
```

```powershell
## Windows 
setx QIYU_APP_KEY "七鱼 APP KEY"
setx QIYU_APP_SECRET "七鱼 APP SECRET"
```

_4. 使用 qiyu-dl 下载某天信息到当前目录_
```
qiyu-dl -date 2023-06-19
```

_5. 解压 zip 格式的文件, 解压密码为七鱼 APP KEY 前12位_

## Document

See [wiki](https://github.com/longbridgeapp/qiyu-dl/wiki/%E8%8E%B7%E5%8F%96%E4%B8%83%E9%B1%BC%E4%BC%9A%E8%AF%9D%E4%BF%A1%E6%81%AF%E6%96%87%E6%A1%A3)
