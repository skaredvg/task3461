package arithmetic

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Statement struct {
	op1  int64
	op2  int64
	sign string
}

func Parsing(expr string, tmpl string) (Statement, error) {
	st := Statement{}
	re, err := regexp.Compile(tmpl)
	if err != nil {
		return st, err
	}
	sl := re.FindStringSubmatch(expr)
	if len(sl) != 6 {
		return st, fmt.Errorf(fmt.Sprintf("Неверное выражение %s", expr))
	}
	n, err := strconv.ParseInt(sl[1], 10, 64)
	if err != nil {
		return st, fmt.Errorf(fmt.Sprintf("Неверный операнд %s", sl[1]))
	}
	st.op1 = n

	st.sign = strings.TrimSpace(sl[2])
	n, err = strconv.ParseInt(sl[3], 10, 64)
	if err != nil {
		return st, fmt.Errorf(fmt.Sprintf("Неверный оеранд %s", sl[3]))
	}
	st.op2 = n
	return st, nil
}

func Calculate(st Statement) (res int64, err error) {
	switch st.sign {
	case "+":
		res = st.op1 + st.op2
	case "-":
		res = st.op1 - st.op2
	case "*":
		res = st.op1 * st.op2
	case "/":
		{
			if st.op2 == 0 {
				err = fmt.Errorf(fmt.Sprintf("Деление на %d недопустимо", st.op2))
				return
			}
			res = int64(st.op1 / st.op2)
		}
	default:
		err = fmt.Errorf(fmt.Sprintf("Операция %d %s %d не поддрживается", st.op1, st.sign, st.op2))
	}
	return
}
