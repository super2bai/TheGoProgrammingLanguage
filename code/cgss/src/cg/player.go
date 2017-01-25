package cg

import (
	"fmt"
)

/**
为了便于演示聊天系统，我们为每个玩家都起了一个独立的goroutine，监听所有发送给他们的聊天信息，
一旦收到就即时打印到控制台上。
*/
type Player struct {
	Name  string
	Level int
	Exp   int
	Room  int
	mq    chan *Message //等待收取的消息
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player)
	return player
}
