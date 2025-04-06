package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/gtlyy/mycoin"
)

func main() {
	// 获取命令行参数
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <numWorkers> <iterations>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1]) // 协程数量
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers. Please provide a positive integer.")
		return
	}

	iterations, err := strconv.Atoi(os.Args[2]) // 循环次数
	if err != nil || iterations <= 0 {
		fmt.Println("Invalid number of iterations. Please provide a positive integer.")
		return
	}

	addrRight := "16HUcQitGbVcFtXCkV6W9VkSEYfJiaZW1f" // 10 BTC
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				// 生成随机 BTC 私钥
				privateKey, _ := mycoin.RandomBtcPrivateKeyMath(64)
				// fmt.Println(privateKey)

				// 根据私钥生成地址
				addr := mycoin.ToAddress2To9(privateKey)

				// 检查地址是否匹配
				if addr == addrRight {
					// 找到匹配的私钥，写入文件
					err := mycoin.WriteToFile("right.txt", privateKey)
					if err != nil {
						fmt.Println("Error: writeToFile().")
					}

					// 立即退出程序
					os.Exit(0)
				}
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
}
