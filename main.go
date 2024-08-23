package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"skillfact/task3461/arithmetic"
	"strconv"
	"strings"
)

func main() {
	rgt := `^([0-9]+)[[:space:]]*([\+\-*/]{1})[[:space:]]*([0-9]+)[[:space:]]*(=)[[:space:]]*([?])[[:space:]]*$`

	fmt.Print("Введите имя входного файла: ")
	nf1 := ""
	if _, err := fmt.Scanln(&nf1); err != nil {
		log.Fatal(err.Error())
	}

	rin, err := os.OpenFile(nf1, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rin.Close()
	bin := bufio.NewReader(rin)

	fmt.Print("Введите имя выходного файла: ")
	nf2 := ""
	if _, err := fmt.Scanln(&nf2); err != nil {
		log.Fatal(err.Error())
	}

	rout, err := os.OpenFile(nf2, os.O_CREATE+os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rout.Close()
	bout := bufio.NewWriter(rout)

	sb := strings.Builder{}
	for {
		line, _, err := bin.ReadLine()
		if err == io.EOF {
			break
		}
		st, err := arithmetic.Parsing(strings.TrimSpace(string(line)), rgt)
		if err != nil {
			continue
		}
		res, err := arithmetic.Calculate(st)
		if err != nil {
			continue
		}
		sb.WriteString(strings.Replace(string(line), "?", strconv.FormatInt(res, 10), -1) + "\n")
	}
	bout.WriteString(sb.String())
	bout.Flush()
}
