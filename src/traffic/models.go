package traffic

const roadLen = 200

type Pair struct {
	x int
	y int
}

type Road struct {
	startAt   Pair
	direction Pair
}

type Car struct {
	onRoad   Road
	position Pair
	speed    int
	crash    bool
}

func (c *Car) Run() {
	c.position.x = ((c.position.x + (c.onRoad.direction.x * c.speed)) % roadLen)
	c.position.y = ((c.position.y + (c.onRoad.direction.y * c.speed)) % roadLen)
}

func (c *Car) Brake() {
	c.speed = c.speed - 1
	if c.speed < 0 {
		c.speed = 0
	}
}

func (c *Car) Speedup() {
	c.speed = c.speed + 1
	if c.speed > 10 {
		c.speed = 10
	}
	c.Run()
}
