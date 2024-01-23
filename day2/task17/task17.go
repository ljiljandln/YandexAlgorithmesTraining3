package main

import "fmt"

type Queue struct {
	data []byte
}

func (q *Queue) size() int {
	return len(q.data)
}

func (q *Queue) pop() byte {
	el := q.data[0]
	q.data = append(q.data[:0], q.data[1:]...)
	return el
}

func (q *Queue) push(num byte) {
	q.data = append(q.data, num)
}

func (q *Queue) push2Cards(card1, card2 byte) {
	q.push(card1)
	q.push(card2)
}

func getQueue() Queue {
	var q Queue
	for i := 0; i < 5; i++ {
		var num byte
		fmt.Scan(&num)
		q.push(num)
	}
	return q
}

func main() {
	pl1 := getQueue()
	pl2 := getQueue()

	count := 0

	for pl1.size() != 0 && pl2.size() != 0 {
		card1 := pl1.pop()
		card2 := pl2.pop()

		if card1 == 0 && card2 == 9 {
			pl1.push2Cards(card1, card2)
		} else if card1 == 9 && card2 == 0 {
			pl2.push2Cards(card1, card2)
		} else if card1 > card2 {
			pl1.push2Cards(card1, card2)
		} else {
			pl2.push2Cards(card1, card2)
		}

		count++
		if count == 1000001 {
			fmt.Print("botva")
			return
		}
	}

	if pl1.size() == 0 {
		fmt.Printf("second %d", count)
	} else {
		fmt.Printf("first %d", count)
	}
}
