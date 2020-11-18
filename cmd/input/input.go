package gobasic

import "github.com/chzyer/readline"

func Input(prompt string, preinput string) string {
        stdin, _ := readline.New(prompt);
        defer stdin.Close();
        stdin.WriteStdin([]byte(preinput));
        value, _ := stdin.Readline()
        return value
}

