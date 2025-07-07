package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// 打印父进程的启动信息，包括父进程的PID
	parentPID := os.Getpid()
	fmt.Println("父进程启动，PID:", parentPID)

	// 创建一个新的命令，用于启动子进程。这里使用当前程序作为子进程并传递参数"child"
	cmd := exec.Command(os.Args[0], "child")
	// 将子进程的标准输出重定向到父进程的标准输出
	cmd.Stdout = os.Stdout
	// 将子进程的标准错误重定向到父进程的标准错误
	cmd.Stderr = os.Stderr

	// 启动子进程，如果启动失败，打印错误信息并退出父进程
	if err := cmd.Start(); err != nil {
		fmt.Println("启动子进程失败:", err)
		return
	}

	// 打印子进程的启动信息，包括子进程的PID
	childPID := cmd.Process.Pid
	fmt.Println("父进程已启动子进程，子进程PID:", childPID)

	// 父进程在一个新的goroutine中继续工作，每隔一秒打印一次工作信息，共打印三次
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("父进程正在工作...")
			time.Sleep(1 * time.Second)
		}
	}()

	// 等待子进程结束，如果子进程异常退出，打印错误信息
	if err := cmd.Wait(); err != nil {
		fmt.Println("子进程异常退出:", err)
	} else {
		// 如果子进程正常退出，打印正常退出的信息
		fmt.Println("子进程正常退出")
	}

	// 打印父进程结束的信息
	fmt.Println("父进程结束")
}

// childProcess函数包含子进程的逻辑
func childProcess() {
	// 打印子进程的启动信息，包括子进程的PID
	fmt.Println("子进程启动，PID:", os.Getpid())
	// 子进程每隔500毫秒打印一次工作信息，共打印五次
	for i := 0; i < 5; i++ {
		fmt.Println("子进程正在工作...")
		time.Sleep(500 * time.Millisecond)
	}
	// 打印子进程结束的信息
	fmt.Println("子进程结束")
}

// init函数在main函数之前执行，检查命令行参数以确定是否作为子进程运行
func init() {
	// 如果命令行参数存在且第二个参数为"child"，则执行子进程的逻辑
	if len(os.Args) > 1 && os.Args[1] == "child" {
		childProcess()
		// 子进程执行完毕后退出
		os.Exit(0)
	}
}
