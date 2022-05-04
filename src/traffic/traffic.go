package traffic

import (
	"OSProject1/src/controllers"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var roads []Road
var cars []Car

type roadMap struct {
	check map[Pair]bool
	mux   sync.Mutex
}

type carsCounter struct {
	num int
	mux sync.Mutex
}

var RoadMap roadMap
var carsOnRoad carsCounter = carsCounter{num: 0}
var crashCarsOnRoad carsCounter = carsCounter{num: 0}
var end = false

func Display(c *gin.Context) {
	var liveMap string
	for i := 0; i < roadLen; i++ {
		liveMap += "."
	}
	RoadMap.mux.Lock()
	for position, hasCar := range RoadMap.check {
		if hasCar == true {
			liveMap = liveMap[:position.x] + "C" + liveMap[position.x+1:]
		}
	}
	RoadMap.mux.Unlock()
	crashCarsOnRoad.mux.Lock()
	if crashCarsOnRoad.num > 0 {
		liveMap += "\n出事啦阿伯..."
	}
	crashCarsOnRoad.mux.Unlock()
	c.String(http.StatusOK, liveMap)
}

func Start(c *gin.Context) {
	crashCarsOnRoad.mux.Lock()
	crashCarsOnRoad.num = 0
	crashCarsOnRoad.mux.Unlock()
	end = false
	RoadMap.mux.Lock()
	RoadMap.check = make(map[Pair]bool)
	RoadMap.mux.Unlock()
	roads = append(roads, Road{startAt: Pair{x: 0, y: 0}, direction: Pair{x: 1, y: 0}})
	roads = append(roads, Road{startAt: Pair{x: 500, y: 0}, direction: Pair{x: 1, y: 1}})

	var newCars []Car
	newCars = append(newCars, newCar())
	cars = newCars
	go cars[0].Depart()
	c.Redirect(http.StatusFound, "/traffic")
}

func End(c *gin.Context) {
	RoadMap.mux.Lock()
	end = true
	for position := range RoadMap.check {
		RoadMap.check[position] = false
	}
	RoadMap.mux.Unlock()

	c.Redirect(http.StatusFound, "/traffic")
}

func AddCar(c *gin.Context) {
	cars = append(cars, newCar())
	go cars[len(cars)-1].Depart()
	c.Redirect(http.StatusFound, "/traffic")
}

func newCar() Car {
	newCar := Car{onRoad: roads[0], position: Pair{0, 0}, speed: 1, crash: false}
	RoadMap.mux.Lock()
	RoadMap.check[newCar.position] = true
	RoadMap.mux.Unlock()
	carsOnRoad.mux.Lock()
	carsOnRoad.num++
	carsOnRoad.mux.Unlock()
	return newCar
}

func (c *Car) carInFront(distance int) bool {
	for i := 0; i < distance; i++ {
		checking := c.onRoad.direction
		checking.x = (checking.x + i) % roadLen
		checking.y = (checking.y + i) % roadLen
		RoadMap.mux.Lock()
		if hasCar, ok := RoadMap.check[checking]; ok && hasCar {
			return true
		}
		RoadMap.mux.Unlock()
	}

	return false
}

func (car *Car) Depart() {
	for !end {
		time.Sleep(200 * time.Millisecond)
		println("(", car.position.x, ", ", car.position.y, ")")
		lastPosition := car.position
		if !car.carInFront(50) {
			car.Speedup()
		} else if !car.carInFront(20) {
			car.Run()
		} else {
			car.Brake()
		}
		RoadMap.mux.Lock()
		if RoadMap.check[car.position] == true {
			RoadMap.mux.Unlock()
			accident(car)
			return
		}
		RoadMap.check[lastPosition] = false
		RoadMap.check[car.position] = true
		RoadMap.mux.Unlock()
	}
}

func accident(c *Car) {
	c.speed = 0
	c.crash = true
	crashCarsOnRoad.mux.Lock()
	crashCarsOnRoad.num++
	crashCarsOnRoad.mux.Unlock()
}

func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "traffic.html", gin.H{
		"name": controllers.GetCookies(c),
	})
}
