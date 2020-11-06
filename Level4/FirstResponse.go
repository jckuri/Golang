package main

import (
	"fmt"
	"time"
)

type Conn struct {
	id int
	ns time.Duration
}

func (conn Conn) String() string {
	return fmt.Sprintf("Conn {id: %v, ms=%v}", conn.id, conn.ns)
}

func (conn *Conn) DoQuery() string {
	time.Sleep(conn.ns)
	return conn.String()
}

func Query(conns []Conn) string {
    ch := make(chan string)
    for _, conn := range conns {
        go func(c Conn) {
            select {
            case ch <- c.DoQuery():
            default:
            }
        }(conn)
    }
    return <-ch
}

func main() {
	conns := []Conn {
		Conn {1, 100},
		Conn {2, 150},
		Conn {3, 200},
		Conn {4, 55},
	}
	for _, conn := range conns {
		fmt.Println(conn)
	}
	fmt.Println("First response:", Query(conns))
}

/*
OUTPUT:

Conn {id: 1, ms=100ns}
Conn {id: 2, ms=150ns}
Conn {id: 3, ms=200ns}
Conn {id: 4, ms=55ns}
First response: Conn {id: 4, ms=55ns}
*/
