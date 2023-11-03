# install

```sh
gbook install -h
NAME:
   gbook install - install plugin
                    `install`: install all plugins from gitbook
                    `install [plugins...]`: install plugin you want, eg: `gbook install code ga`

USAGE:
   gbook install [command options] [arguments...]

OPTIONS:
   --help, -h  show help
```

- `gbook install [plugins...]` 会默认调用 `gbook ready` 和 `gbook sync2`命令，确保依赖正常后才会继续安装依赖。
- 使用前请确保设置好 npm 加速源
