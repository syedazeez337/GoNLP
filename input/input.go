package input

import (
	"bufio"
	"log"
	"os"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)

	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	message = message[:len(message)-1]

	return message
}
