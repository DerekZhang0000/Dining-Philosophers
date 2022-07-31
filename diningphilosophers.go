package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var host Host
var dinner sync.WaitGroup

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	name           string
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

type Host struct {
	sync.Mutex
	eating int
}

func (philo *Philosopher) waitToEat() bool {
	if host.allowEat() {
		return true
	} else {
		time.Sleep(1)
		return philo.waitToEat()
	}
}

func (philo *Philosopher) eat() {
	philo.waitToEat()
	fmt.Println(philo.name, "is eating")
	philo.leftChopstick.Lock()
	philo.rightChopstick.Lock()
	fmt.Println(philo.name, "finished eating")
	philo.leftChopstick.Unlock()
	philo.rightChopstick.Unlock()
	philo.finishEat()
	dinner.Done()
}

func (philo *Philosopher) finishEat() {
	host.finishEat()
}

func (h *Host) allowEat() bool {
	h.Lock()
	defer h.Unlock()
	if h.eating == 2 {
		return false
	}
	h.eating++
	return true
}

func (h *Host) finishEat() {
	h.Lock()
	defer h.Unlock()
	h.eating--
}

func main() {
	dinner.Add(5)

	Chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		Chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			name:           "Philosopher " + strconv.Itoa(i+1),
			leftChopstick:  Chopsticks[i],
			rightChopstick: Chopsticks[(i+1)%5],
		}
	}

	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}

	dinner.Wait()
}
