package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 假设你有一个字符串表示的 Docker Compose YAML
	composeYAML := `
version: '3'
services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
  db:
    image: mysql:latest
    environment:
      - MYSQL_ROOT_PASSWORD=secret
`

	// 定义输出文件名
	outputFile := "docker-compose.yml"

	// 将字符串写入文件
	err := ioutil.WriteFile(outputFile, []byte(composeYAML), 0644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Docker Compose YAML has been written to %s\n", outputFile)
}
