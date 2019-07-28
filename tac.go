package main

import "fmt"
import "strings"
import "os"
import "bufio"
import "strconv"

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type ttb struct {
	top    []string
	middle []string
	bottom []string
}

type coord struct {
	x, y int
}

func (t *ttb) addToken(c coord) {
	c.translate()

}
func (t ttb) printBoard() {
	fmt.Println(strings.Join(t.top, ""))
	fmt.Println("-----")
	fmt.Println(strings.Join(t.middle, ""))
	fmt.Println("-----")
	fmt.Println(strings.Join(t.bottom, ""))
}
func (c *coord) translate() {
	switch c.x {
	case 1:
		c.x = 2
	case 2:
		c.x = 4
	}

	switch c.y {
	case 1:
		c.y = 2
	case 2:
		c.y = 4
	}
	return
}
func createTtb() ttb {
	var t ttb
	t.top = []string{"-", "|", "-", "|", "-"}
	t.middle = []string{"-", "|", "-", "|", "-"}
	t.bottom = []string{"-", "|", "-", "|", "-"}
	return t
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	t := createTtb()
	t.printBoard()
	fmt.Print("row-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	row, err := strconv.Atoi(text)
	check(err)
	fmt.Print("col-> ")
	textt, _ := reader.ReadString('\n')
	textt = strings.Replace(textt, "\n", "", -1)
	col, err := strconv.Atoi(textt)
	check(err)
	token := coord{x: row, y: col}
	fmt.Println(token)
}
