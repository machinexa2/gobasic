package gobasic

import "bufio"
import "fmt"
import "os"

import "github.com/chzyer/readline"

func InArray(check string, list []string) bool {
        for _, individual := range list {
                if individual == check {
                        return true
                }
        }
        return false
}

func Input(prompt string, preinput string) string {
        stdin, _ := readline.New(prompt);
        defer stdin.Close();
        stdin.WriteStdin([]byte(preinput));
        value, _ := stdin.Readline()
        return value
}

func ReadFile(filename string, c chan string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	close(c)
	file.Close()
}
