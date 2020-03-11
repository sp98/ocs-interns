//The program displays Hello World by calling functions from a local package
package main

import (
	"fmt"
	"ocs-interns/yug-gupta/golang-docker-application/localpackages" // local package with hello() and world() functions
)

func main() {
	hello := localpackages.Hello()
	world := localpackages.World()
	fmt.Println(hello, world)
}
