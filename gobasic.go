package gobasic

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
