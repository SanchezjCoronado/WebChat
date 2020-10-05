package chat

import "fmt"

type Chat struct {
}

func (c *Chat) Run() {
	for {
		fmt.Println("Chat running")
	}
}

func Start() {
	c := &Chat{}
	c.Run()
}
