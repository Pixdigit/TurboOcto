package GameEngine

import "github.com/ByteArena/box2d"

var World box2d.B2World

func Test(){
    gravity := box2d.MakeB2Vec2(0.0, -10.0)
    World = box2d.MakeB2World(gravity)
}
