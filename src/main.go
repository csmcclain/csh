package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Execution struct {
	output  string
	success bool
}

func executor(c chan Execution, program string, args ...string) {
	out, err := exec.Command(program, args...).Output()
	var result Execution

	if err != nil {
		result = Execution{
			output:  err.Error(),
			success: false,
		}

	} else {
		result = Execution{
			output:  string(out[:]),
			success: true,
		}
	}

	c <- result
}

func main() {
	c := make(chan Execution)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("CSH: Please input command: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		brokenUpInput := strings.Split(text, " ")
		if text == "exit" {
			break
		}
		args := brokenUpInput[1:]
		go executor(c, brokenUpInput[0], args...)

		result := <-c
		if result.success {
			fmt.Println(result.output)
		} else {
			fmt.Errorf(result.output)
		}
	}

}
