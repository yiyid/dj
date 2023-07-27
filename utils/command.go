package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func Exec(command string) error {
	cmd := exec.Command("bash", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("无法获取标准输出管道: %s\n", err)
		return err
	}

	err = cmd.Start()
	if err != nil {
		log.Printf("无法启动命令: %s\n", err)
		return err
	}

	reader := bufio.NewReader(stdout)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("读取输出时发生错误: %s\n", err)
			break
		}
		fmt.Printf("%s", line)
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("ERROR 命令执行错误: %s\n", err)
		return err
	}
	return nil
}
