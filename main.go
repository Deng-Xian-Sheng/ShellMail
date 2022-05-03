package main

import (
	"flag"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

// Config 配置文件的结构体
type Config struct {
	Host     string `yaml:"host"`     //服务器地址
	Port     int    `yaml:"port"`     //服务器端口
	Username string `yaml:"username"` //用户名
	Password string `yaml:"password"` //密码
}

// 用于获取配置文件的函数
func parseYaml(file string) (Config, error) {
	var config Config
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}
	return config, err
}

func main() {
	// -f 指定配置文件
	configFile := flag.String("f", "config.yaml", "config file")
	// -t 指定收件人
	to := flag.String("t", "", "指定收件人")
	// -s 指定主题
	subject := flag.String("s", "", "指定主题")
	// -b 指定内容
	body := flag.String("b", "", "指定内容")
	flag.Parse()

	// 判断没有输入任何flag
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	// 解析配置文件
	config, err := parseYaml(*configFile)
	if err != nil {
		log.Panic(err)
	}

	// 判断收件人是否为空
	if len(strings.TrimSpace(*to)) == 0 {
		log.Panic("收件人不能为空")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Username)
	m.SetHeader("To", *to)
	m.SetHeader("Subject", *subject)
	m.SetBody("text/html", *body)

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Panic(err)
	}
}
