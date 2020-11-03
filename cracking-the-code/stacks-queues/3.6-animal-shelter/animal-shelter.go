package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type List struct {
	queueDog *list.List
	queueCat *list.List
}

type Animal struct {
	name   string
	time   int
	animal string
}

func main() {

	queueDog := list.New()
	queueCat := list.New()

	list := List{}

	list.queueDog = queueDog
	list.queueCat = queueCat

	list.AddAnimal("cat", "wilmot", 1)
	list.AddAnimal("cat", "garfield", 2)

	list.AddAnimal("dog", "charley", 3)
	list.AddAnimal("dog", "coconut", 4)

	//list.DequeueDog()
	//list.DequeueCat()
	list.DequeueAny()

	// for list.queueCat.Len() > 0 {
	// 	e := list.queueCat.Front() // First element
	// 	fmt.Print(e.Value)

	// 	list.queueCat.Remove(e) // Dequeue
	// 	//break
	// }

	// for list.queueDog.Len() > 0 {
	// 	e := list.queueDog.Front() // First element
	// 	fmt.Print(e.Value)

	// 	list.queueDog.Remove(e) // Dequeue
	// 	//break
	// }
}

func (l *List) AddAnimal(animal string, name string, time int) {

	var a Animal
	a.name = name
	a.animal = animal
	a.time = time

	if animal == "cat" {
		l.queueCat.PushBack(a)
	}

	if animal == "dog" {
		l.queueDog.PushBack(a)
	}
}

func (l *List) DequeueDog() {

	fmt.Println("dequeing dog")
	for l.queueDog.Len() > 0 {
		e := l.queueDog.Front() // First element
		fmt.Print(e.Value)

		l.queueDog.Remove(e) // Dequeue
		break
	}

}

func (l *List) DequeueCat() {

	fmt.Println("dequeing cat")
	for l.queueCat.Len() > 0 {
		e := l.queueCat.Front() // First element
		fmt.Print(e.Value)

		l.queueCat.Remove(e) // Dequeue
		break
	}

}

func (l *List) DequeueAny() {

	if l.queueCat.Len() > 0 && l.queueDog.Len() > 0 {

		c := l.queueCat.Front()
		d := l.queueDog.Front()

		vC := reflect.ValueOf(c.Value)
		vD := reflect.ValueOf(d.Value)
		timeDog := vD.FieldByName("time")
		timeCat := vC.FieldByName("time")

		if timeDog.Int() > timeCat.Int() {
			l.DequeueCat()
		} else {
			l.DequeueCat()
		}
	}
}
