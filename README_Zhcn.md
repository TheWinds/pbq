# PBQ (Picture Bed QiNiu)
这个工具能够帮助你很方便的通过命令把图片上传到七牛云,上传完毕回返回外链并且复制生成的MarkDown代码到剪切板。这在写博客时是非常方便的

## 命令
|命令|描述
|:-|-|
|`pqb <fileName>` | 上传图片到你的七牛云相应bucket中 |
|`pqb account <AccessKey> <SecretKey> <BucketName>` | 配置账户信息|
|`pqb layout <Layout>` | 设置文件名布局|

## 关于文件名布局

|占位符|描述|例子|
|:-|-|-|
|`%FILENAME` | 文件名(必选,否则使用默认文件名布局) |test.png|
|`%YYYY` | 当前年 |2017|
|`%MM` | 当前月 |01|
|`%DD` | 当前天 |02|
|`%UNIX` | 时间戳 |1497407382|

- 如果不设置Layout,默认为 `%YYYY%MM%DD%UNIX-%FILENAME` => `201706141497407382-test.png`

[![asciicast](https://asciinema.org/a/6ga6ab4k5jp9g6wf90g3kh7a8.png)](https://asciinema.org/a/6ga6ab4k5jp9g6wf90g3kh7a8)
