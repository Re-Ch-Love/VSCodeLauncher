package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("获取用户目录时发生错误：", err)
	}
	homeDir := user.HomeDir
	logFile, err := os.OpenFile(filepath.Join(homeDir, ".vscode_launcher.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal("创建log文件时发生错误：", err)
	}
	stdout := logFile
	log.SetOutput(stdout)

	confFile, err := ioutil.ReadFile(filepath.Join(homeDir, ".vscode_launcher.conf"))
	if err != nil {
		log.Fatal("配置文件读取失败：", err)
	}

	tmp := strings.SplitN(string(confFile), "\n", 2)
	if len(tmp) < 2 {
		log.Fatal("配置文件不合法")
	}
	exe, confArgsString := tmp[0], tmp[1]
	if exe[len(exe)-1] == '\r' {
		exe = exe[:len(exe)-1]
	}
	confArgsArray := ParseArgs(confArgsString)
	allArgs := append(confArgsArray, os.Args[1:]...)
	cmd := exec.Command(exe, allArgs...)
	log.Print(cmd.Args)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	if err := cmd.Run(); err != nil {
		log.Fatal("启动vscode时发生错误：", err)
	}
	log.Print("exit vscode")
}

func ParseArgs(str string) []string {
	var cmd []string
	var buffer bytes.Buffer
	var isInQuotes bool
	for i, r := range str {
		if r == ' ' && !isInQuotes {
			if buffer.Len() != 0 {
				cmd = append(cmd, buffer.String())
				buffer.Reset()
			}
		} else if r == '"' {
			isInQuotes = !isInQuotes
			if buffer.Len() != 0 || str[i+1] == '"' {
				cmd = append(cmd, buffer.String())
				buffer.Reset()
			}
		} else {
			buffer.WriteRune(r)
		}
		if i == len(str)-1 {
			cmd = append(cmd, buffer.String())
			buffer.Reset()
		}
	}
	return cmd
}
