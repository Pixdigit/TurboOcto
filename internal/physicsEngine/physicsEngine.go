package physicsEngine

import "github.com/ByteArena/box2d"

func CreateWorld(xGravity,yGravity float64) (box2d.B2World){
    gravity := box2d.MakeB2Vec2(xGravity, yGravity)
    world := box2d.MakeB2World(gravity)
    return world
}
