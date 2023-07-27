package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/yiyid/dj/utils"
)

/*
若更改版本需要修改此处URL
*/
var urls = []string{
	"https://repo.huaweicloud.com/java/jdk/8u151-b12/jdk-8u151-linux-x64.tar.gz",
	"https://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-9/v9.0.76/bin/apache-tomcat-9.0.76.tar.gz",
	"https://mirrors.jenkins.io/war-stable/2.164.1/jenkins.war",
}

/*
若更改版本需要修改此处变量 jdk_dir 的值
*/
var (
	jdk_gz     = path.Base(urls[0])
	jdk_dir    = "jdk1.8.0_151"
	tomcat_gz  = path.Base(urls[1])
	tomcat_dir = strings.Split(tomcat_gz, ".tar.gz")[0]
)

func LinuxDownload() {

	// 记录不存在的文件
	missingFiles := []string{}

	// 检查文件是否存在
	for _, url := range urls {
		if !utils.FileExists(path.Base(url)) {
			missingFiles = append(missingFiles, url)
		}
	}

	// 打印提示信息
	if len(missingFiles) > 0 {
		for _, file := range missingFiles {
			fmt.Printf("文件不存在，请下载后放到本目录: %s\n", file)
		}
		os.Exit(1)
	}
}

func LinuxInstall() {

	fmt.Println("开始安装jdk...")
	fmt.Println("开始解压jdk...")
	utils.Exec(fmt.Sprintf("tar xf %s -C /usr/local/", jdk_gz))
	utils.Exec(fmt.Sprintf("mv /usr/local/%s /usr/local/java", jdk_dir))
	utils.Exec("echo 'JAVA_HOME=/usr/local/java; PATH=$JAVA_HOME/bin:$PATH; export JAVA_HOME PATH' >> /etc/profile; source /etc/profile;")

	fmt.Println("开始安装tomcat...")
	fmt.Println("开始解压tomcat...")
	utils.Exec(fmt.Sprintf("tar xf %s -C /usr/local/;", tomcat_gz))
	utils.Exec(fmt.Sprintf("mv /usr/local/%s /usr/local/tomcat", tomcat_dir))
	utils.Exec("echo 'CATALINA_HOME=/usr/local/tomcat; export CATALINA_HOME PATH' >> /etc/profile; source /etc/profile;")

	fmt.Println("开始安装jenkins...")
	utils.Exec("rm -rf /usr/local/tomcat/webapps/*")
	utils.Exec("cp jenkins.war /usr/local/tomcat/webapps/")
	utils.Exec("chmod +x /usr/local/tomcat/bin/*")
	fmt.Println("手动启动 jenkins 项目: source /etc/profile && /usr/local/tomcat/bin/startup.sh")

}

func LinuxDeploy() {
	// 下载 jenkins 相关安装包
	LinuxDownload()

	// 安装 jenkins 相关安装包
	LinuxInstall()
}
