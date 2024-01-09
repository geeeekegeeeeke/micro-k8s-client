package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main2() {
	// 读取 JSON 文件内容
	data, err := ioutil.ReadFile("scan-results.json")
	if err != nil {
		log.Fatal(err)
	}

	// 解析 JSON 数据到 map[string]interface{} 类型
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	// 检查是否存在 package 键
	results, ok := jsonData["results"]
	if !ok {
		log.Fatal("results key not found")
	}
	// 将 package 转换为 map[string]interface{} 类型
	rsl, ok := results.(map[string]interface{})
	// 将 package 转换为 map[string]interface{} 类型
	//packageData, ok := pkg.(map[string]interface{})
	// 检查是否存在 package 键
	pkg, ok := rsl["packages"]
	if !ok {
		log.Fatal("package key not found")
	}

	// 将 package 转换为 map[string]interface{} 类型
	packageData, ok := pkg.(map[string]interface{})
	if !ok {
		log.Fatal("package is not an object")
	}

	// 检查是否存在 vulnerabilities 键
	vulnerabilities, ok := packageData["vulnerabilities"]
	if !ok {
		log.Fatal("vulnerabilities key not found")
	}

	// 将 vulnerabilities 转换为 []interface{} 类型
	vulnerabilitiesList, ok := vulnerabilities.([]interface{})
	if !ok {
		log.Fatal("vulnerabilities is not an array")
	}

	// 遍历 vulnerabilities 数组并打印属性
	for _, v := range vulnerabilitiesList {
		vulnerability, ok := v.(map[string]interface{})
		if !ok {
			log.Fatal("vulnerability is not an object")
		}

		// 打印属性
		for key, value := range vulnerability {
			fmt.Printf("%s: %v\n", key, value)
		}
		fmt.Println()
	}
}

func main() {
	//jsonData := `{
	//	"results": [
	//		{
	//			"source": {
	//				"path": "/src/go.mod",
	//				"type": "lockfile"
	//			},
	//			"packages": [
	//				{
	//					"package": {
	//						"name": "stdlib",
	//						"version": "1.20.10\n",
	//						"ecosystem": "Go",
	//						"commit": ""
	//					},
	//					"vulnerabilities": [
	//
	//					]
	//				}
	//			]
	//		}
	//	]
	//}`

	jsonData, err := ioutil.ReadFile("scan-results.json")
	if err != nil {
		log.Fatal(err)
	}

	//// 解析 JSON 数据到 map[string]interface{} 类型
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	//fmt.Println(data)
	resultsRaw, ok := data["results"]
	//resultsRaw, ok := resultsRaw["pacakges"]
	fmt.Println("JSON数据中缺少results字段", resultsRaw)
	if !ok {
		fmt.Println("JSON数据中缺少results字段", resultsRaw)
		return
	}
	// 访问解析后的数据
	results, ok := data["results"].([]interface{})
	if !ok {
		fmt.Println("无效的JSON1结构")
		return
	}

	for _, result := range results {
		resultMap, ok := result.(map[string]interface{})
		if !ok {
			fmt.Println("无效的JSON2结构")
			return
		}

		packages, ok := resultMap["packages"].([]interface{})
		if !ok {
			fmt.Println("无效的JSON3结构")
			return
		}

		for _, pkg := range packages {
			pkgMap, ok := pkg.(map[string]interface{})
			if !ok {
				fmt.Println("无效的JSON4结构")
				return
			}

			vulnerabilities, ok := pkgMap["vulnerabilities"].([]interface{})
			if !ok {
				fmt.Println("无效的JSON5结构")
				return
			}

			for _, v := range vulnerabilities {
				vMap, ok := v.(map[string]interface{})
				if !ok {
					fmt.Println("无效的JSON6结构")
					return
				}

				name, ok := vMap["modified"].(string)
				if !ok {
					fmt.Println("无效的JSON7结构")
					return
				}

				version, ok := vMap["published"].(string)
				if !ok {
					fmt.Println("无效的JSON8结构")
					return
				}

				ecosystem, ok := vMap["schema_version"].(string)
				if !ok {
					fmt.Println("无效的JSON9结构")
					return
				}

				commit, ok := vMap["summary"].(string)
				if !ok {
					fmt.Println("无效的JSON0结构")
					return
				}

				fmt.Println("Vulnerability Name:", name)
				fmt.Println("Vulnerability Version:", version)
				fmt.Println("Vulnerability Ecosystem:", ecosystem)
				fmt.Println("Vulnerability Commit:", commit)
			}
		}
	}
}

func main3() {
	// 读取 JSON 文件内容
	data, err := ioutil.ReadFile("scan-results.json")
	if err != nil {
		log.Fatal(err)
	}

	// 解析 JSON 数据到 RawMessage 类型
	var raw json.RawMessage
	err = json.Unmarshal(data, &raw)
	if err != nil {
		log.Fatal(err)
	}
	// 打印解析后的 JSON 数据
	fmt.Println(string(raw))
}
