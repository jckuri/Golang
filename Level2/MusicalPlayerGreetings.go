package main

import "fmt"

type MusicalPlayer interface {
	Greetings() string
}

type Trumpeter struct {
	Name string
}

type Violinist struct {
	Name string
}

func (trumpeter *Trumpeter) Greetings() string {
	return fmt.Sprintf("Hello trumpeter %s!", trumpeter.Name)
}

func (violinist *Violinist) Greetings() string {
	return fmt.Sprintf("Hello violinist %s!", violinist.Name)
}

func main() {
	players := make([]MusicalPlayer, 4)
	players[0] = &Trumpeter{"Jack White"}
	players[1] = &Trumpeter{"John Jordan"}
	players[2] = &Violinist{"Mary Zen"}
	players[3] = &Violinist{"Donald Kurzweil"}
	fmt.Println(players)
	for i := range players {
		fmt.Println(players[i].Greetings())
	}
}
