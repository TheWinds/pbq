# PBQ (Picture Bed QiNiu)
This tool can help you to easily upload pictures to QiNiu Cloud, you can also use it to generate Markdown code.

## [👉 中文文档](https://github.com/TheWinds/pbq/blob/master/README-zhCN.md)

## Command
|Command|Description|
|:-|-|
|`pqb <fileName>` | Upload picture to your bucket |
|`pqb account <AccessKey> <SecretKey> <BucketName>` | Configure the account|
|`pqb layout <Layout>` | Set the file name layout|

## About File Name Layout

|Placeholder|Description|Examle|
|:-|-|-|
|`%FILENAME` | filename(Required,Otherwise use the default file name layout) |test.png|
|`%YYYY` |current year |2017|
|`%MM` | current month |01|
|`%DD` | current day |02|
|`%UNIX` | timestamp |1497407382|

- If you do not set the Layout, the default is `%YYYY%MM%DD%UNIX-%FILENAME` => `201706141497407382-test.png`

[![asciicast](https://asciinema.org/a/6ga6ab4k5jp9g6wf90g3kh7a8.png)](https://asciinema.org/a/6ga6ab4k5jp9g6wf90g3kh7a8)

## Download
|Platform|Link|
|-|-|
|Linux 64bit|[Download](http://ohrkcds09.bkt.clouddn.com/pbq-1.0-linux64.zip)|
|Darwin 64bit|[Download](http://ohrkcds09.bkt.clouddn.com/pbq-1.0-darwin64.zip)|
|Windows 64bit|[Download](http://ohrkcds09.bkt.clouddn.com/pbq-1.0-win64.zip)|
|Windows 32bit|[Download](http://ohrkcds09.bkt.clouddn.com/pbq-1.0-win32.zip)|
