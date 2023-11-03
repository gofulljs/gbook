# sync

```sh
gbook sync2 -h
NAME:
   gbook sync2 - sync2 gitbook, 不包含node_modules, suggest

USAGE:
   gbook sync2 [command options] [arguments...]

OPTIONS:
   --source value  gitbook数据源 (default: "https://github.com/gofulljs/gitbook/archive/refs/tags/3.2.3.tar.gz")
   --proxy1 value  自定义加速源(前缀+source), 不传采用以下\n[https://ghps.cc/ https://gh.api.99988866.xyz/ https://github.abskoop.workers.dev/]
   --proxy2 value  自定义加速源(替换https://github.com前缀)
   --help, -h      show help

```

拉取 gitbook 引擎，拉取完成后会自动调用 `npm install` 安装依赖
