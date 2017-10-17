package physicsEngine

import "github.com/ByteArena/box2d"

var Env box2d.B2World

func CreateWorld(xGravity,yGravity float64) (box2d.B2World){
    gravity := box2d.MakeB2Vec2(xGravity, yGravity)
    Env = box2d.MakeB2World(gravity)
    return Env
}
