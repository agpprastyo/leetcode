import (
	"strconv"
)

func maximum69Number (num int) int {
    s := strconv.Itoa(num)

    bs := []byte(s)

    for i := 0; i < len(bs); i++ {
        if bs[i] == '6' {
            bs[i] = '9'
            break
        }
    }

    res, err := strconv.Atoi(string(bs))
    if err != nil {
        panic(err)
    }
    return res
 }