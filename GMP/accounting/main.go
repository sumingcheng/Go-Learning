package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	ledger := NewLedger()
	var wg sync.WaitGroup

	wg.Add(1)
	go ledger.ProcessCommands(&wg)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("欢迎使用记账本")
		fmt.Println("1. 添加账单")
		fmt.Println("2. 查看账单")
		fmt.Println("3. 退出")
		fmt.Print("请输入操作数字: ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("请输入金额: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			fmt.Print("请输入分类: ")
			scanner.Scan()
			category := strings.TrimSpace(scanner.Text())

			ledger.commands <- Command{"add", amount, category}
		case "2":
			ledger.commands <- Command{"list", 0, ""}
		case "3":
			close(ledger.commands)
			wg.Wait()
			return
		default:
			fmt.Println("无效输入，请重新输入")
		}
	}
}
