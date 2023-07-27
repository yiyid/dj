package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/yiyid/dj/utils"
)

var urls = []string{
	"https://repo.huaweicloud.com/java/jdk/8u151-b12/jdk-8u151-linux-x64.tar.gz",
	"https://dlcdn.apache.org/tomcat/tomcat-9/v9.0.78/bin/apache-tomcat-9.0.78.zip",
	"https://mirrors.jenkins.io/war-stable/2.164.1/jenkins.war",
}

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
	utils.Exec("tar xvf jdk-8u151-linux-x64.tar.gz -C /usr/local/")
	utils.Exec("mv /usr/local/jdk1.8.0_151/ /usr/local/java")
	utils.Exec("echo 'JAVA_HOME=/usr/local/java; PATH=$JAVA_HOME/bin:$PATH; export JAVA_HOME PATH' >> /etc/profile; source /etc/profile;")
	fmt.Println("重启shell后执行 java -verion")
	fmt.Println("开始安装tomcat...")
	tomcat_zip := path.Base(urls[1])
	tomcat_dir := strings.Split(tomcat_zip, ".zip")[0]
	err := utils.Exec("yum install unzip -y")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = utils.Exec(fmt.Sprintf("unzip %s; mv %s /usr/local/tomcat/", tomcat_zip, tomcat_dir))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	utils.Exec("echo 'CATALINA_HOME=/usr/local/tomcat/; export CATALINA_HOME PATH' >> /etc/profile; source /etc/profile;")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	utils.Exec("rm -rf /usr/local/tomcat/webapps/*")
	utils.Exec("cp jenkins.war /usr/local/tomcat/webapps/")
	utils.Exec("chmod +x /usr/local/tomcat/bin/*")
	utils.Exec("/usr/local/tomcat/bin/startup.sh")

}

func LinuxDeploy() {
	// 下载 jenkins 相关安装包
	LinuxDownload()

	// 安装 jenkins 相关安装包
	LinuxInstall()
}
