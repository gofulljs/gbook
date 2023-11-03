# 命令概览

```sh
$ gbook -h
NAME:
   gbook - uniswap tick update, other command will forward gitbook *

USAGE:
   gbook [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   ready       check env is ready
   sync        sync gitbook
   sync2       sync2 gitbook, 不包含node_modules, suggest
   install, i  install plugin
                `install`: install all plugins from gitbook
                `install [plugins...]`: install plugin you want, eg: `gbook install code ga`
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --bookVersion value, --bv value  (default: "3.2.3") [%BOOK_VERSION%]
   --bookHome value, --bh value     gitbook path, default is $HOME/.gitbook/versions/
   --nodePath value                 nodejs home, if not specified, use current node [%BOOK_NODE_HOME%]
   --help, -h                       show help
   --version, -v                    print the version
```

- `bookVersion`: 默认使用 gitbook 引擎为 3.2.3
- `bookHome`: 默认 gitbook 引擎存放位置
- `nodePath`: `gbook` 使用指定路径的 node 版本，建议将`BOOK_NODE_HOME`配置到环境变量中，这样切换 node 版本不将影响到`gbook`
