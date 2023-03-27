package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	//ssh 時可以先以ssh -v連線看到詳細連訊息, 例如AuthMethod
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string,
				question []string,
				echos []bool) ([]string, error) {
				answers := make([]string, len(question))
				for i, _ := range question {
					answers[i] = "mypassword"
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", "192.168.10.230:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer connection.Close()

	session, err := connection.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	//一旦連線session被建立, 可以使用Run 這個方法來執行command
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls"); err != nil {
		log.Fatal("Failed to run: ", err)
	}
	fmt.Println(b.String())
}
