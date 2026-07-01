package main 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	queue := NewQueue[string]()

	for c := 'A'; c <= 'P'; c++ {
		queue.Enqueue(string(c))
	}

	scanner := bufio.NewScanner(os.Stdin)

	for queue.items.Len() > 1 {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		golsEsquerda, _ := strconv.Atoi(parts[0])
		golsDireita, _ := strconv.Atoi(parts[1])

		timeEsquerda := queue.Dequeue()
		timeDireita := queue.Dequeue()

		if golsEsquerda > golsDireita {
			queue.Enqueue(timeEsquerda)
		} else {
			queue.Enqueue(timeDireita)
		}
	}

	campeao := queue.Dequeue()
	fmt.Println(campeao)
}