package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type User struct {
	id   string
	name string
	msg  chan string
}

// 存储全部用户的map
var allUsers = make(map[string]*User)

// 广播管道
var message = make(chan string, 10)

// 全局锁
var lock sync.RWMutex

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 启动广播go程
	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	user := User{
		id:   remoteAddr,
		name: remoteAddr,
		msg:  make(chan string),
	}
	//加入用户到map
	lock.Lock()
	allUsers[user.id] = &user
	lock.Unlock()
	//启动接收消息并发送给客户的go程
	go writeBackToClient(&user, conn)
	//监听退出的信号
	isQuit := make(chan bool)
	//刷新时间的信号
	refresh := make(chan bool)
	go isStop(&user, conn, isQuit, refresh)
	//广播上线通知
	loginInfo := fmt.Sprintf("[%s]:[%s] 已上线...\n", user.id, user.name)
	message <- loginInfo
	//业务逻辑
	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		//用户ctrl+c退出
		if cnt == 0 {
			isQuit <- true
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		userInput := string(buf[:cnt-1])
		var userInfos []string
		if len(userInput) == 3 && userInput == "who" {
			lock.Lock()
			for _, val := range allUsers {
				userInfo := fmt.Sprintf("user_name:%s", val.name)
				userInfos = append(userInfos, userInfo)
			}
			lock.Unlock()
			user.msg <- strings.Join(userInfos, "\n")
		} else if len(userInput) >= 8 && userInput[:6] == "rename" {
			user.name = strings.Split(userInput, " ")[1]
			user.msg <- "rename success!"
		} else {
			message <- userInput
		}
		refresh <- true
	}
}

func broadcast() {
	for {
		for info := range message {
			//向全体用户发送广播消息
			for _, user := range allUsers {
				user.msg <- info
			}
		}
	}
}

func writeBackToClient(user *User, conn net.Conn) {
	for data := range user.msg {
		_, _ = conn.Write([]byte(data))
	}
}

func isStop(user *User, conn net.Conn, isQuit <-chan bool, refresh <-chan bool) {
	defer fmt.Println(user.name, " stopped")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s logout!", user.name)
			message <- logoutInfo
			delete(allUsers, user.id)
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout!", user.name)
			message <- logoutInfo
			delete(allUsers, user.id)
			conn.Close()
			return
		case <-refresh:
			fmt.Println("refresh timer!")
		}

	}
}
