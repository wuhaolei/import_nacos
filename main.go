package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"os"
)

func main() {
	var (
		ipAddr      string
		port        uint64
		scheme      string
		contextPath string
		namespaceId string
		username    string
		password    string
		dataId      string
		group       string
		file        string
		configType  string
	)
	flag.StringVar(&username, "u", "", "账户")
	flag.StringVar(&password, "p", "", "密码")
	flag.StringVar(&ipAddr, "H", "", "地址")
	flag.Uint64Var(&port, "P", 8848, "端口")
	flag.StringVar(&scheme, "scheme", "http", "协议")
	flag.StringVar(&contextPath, "contextPath", "/nacos", "URI")
	flag.StringVar(&namespaceId, "n", "", "Namespace ID")
	flag.StringVar(&dataId, "d", "", "Data ID")
	flag.StringVar(&group, "g", "", "Group")
	flag.StringVar(&file, "f", "", "文件路径")
	flag.StringVar(&configType, "t", "yaml", "文件类型")
	flag.Parse()

	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId,
		TimeoutMs:           5 * 1000,
		NotLoadCacheAtStart: true,
		LogDir:              "./",
		LogLevel:            "debug",
		Username:            username,
		Password:            password,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      ipAddr,
			ContextPath: contextPath,
			Port:        port,
			Scheme:      scheme,
		},
	}
	if ipAddr != "" && dataId != "" && username != "" && password != "" && namespaceId != "" && group != "" && file != "" {
		client, err := clients.CreateConfigClient(map[string]interface{}{
			"clientConfig":  clientConfig,
			"serverConfigs": serverConfigs,
		})
		if err != nil {
			panic(err)
		}

		content, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		success, err := client.PublishConfig(vo.ConfigParam{
			DataId:  dataId,
			Group:   group,
			Content: string(content),
			Type:    configType,
		})
		fmt.Println(success)
	} else {
		fmt.Println("Usage:")
		fmt.Println("  -h or -help or --help")
	}
}
