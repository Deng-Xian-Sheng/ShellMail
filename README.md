# Shell Mail 通过 Shell 命令发送邮件

## 为什么开发 Shell Mail

因为我发现使用Linux发送邮件很麻烦。

我并不是要部署邮件服务器；而是想发送一封简单的邮件；因为国内大多云服务器厂商封禁了25端口，所以我需要使用Linux的mail命令并且指定外部的smtp服务器。

问题就出现在这儿，虽然我不需要部署邮件服务器；但是我仍然需要安装它们的依赖包，因为mail命令需要；且Centos、Ubuntu各个发行版的mail命令还不一样；依赖包也不一样。

我感到头疼，这浪费了我的时间，我想让这更加简单。

## 什么时候需要 Shell Mail

Shell Mail顾名思义，例如你有一个Shell脚本，在某种情况下(例如运行错误或者需要通知等)，你需要发送邮件，那么你可以使用Shell Mail。

## 如何使用 Shell Mail

你需要提供一个config.yaml配置文件(除了后缀，名称随意);它需要包含以下内容：

这些内容你可以在你的邮件平台上获取

```yaml
host: smtp.sina.com #smtp服务器地址
port: 587 #smtp端口
username: x17615848429@sina.com #用户名
password: xxxxxxx #smtp授权码
```

发送邮件：

```bash
./shellMail -f config.yaml -t x17615848429@sina.com -s "like" -b "I like 邓文怡"
# -f 配置文件
# -t 收件人
# -s 主题
# -b 内容
```