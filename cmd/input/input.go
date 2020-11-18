package gobasic

import "bufio"
import "fmt"
import "os"
import "strings"
import "github.com/chzyer/readline"

func InputRead() string {
        reader := bufio.NewReader(os.Stdin);
        fmt.Print("> ");
        text, _ := reader.ReadString('\n');
        text = strings.Replace(text, "\n", "", -1);
        return text
}

func PrefilledInputRead(preinput string) string {
        stdin, _ := readline.New("> ");
        defer stdin.Close();
        stdin.WriteStdin([]byte(preinput));
        value, _ := stdin.Readline()
        return value
}

