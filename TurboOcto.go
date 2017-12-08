package TurboOcto

import (
    "github.com/Pixdigit/TurboOcto/internal/renderEngine"
    "github.com/Pixdigit/TurboOcto/internal/sharedStructs"
)

func Init() {
    renderEngine.Init()
}

func Blit(sprite sharedStructs.Sprite, env Environment) {
    posVec := sprite.Body.GetPosition()
    sprite.Rect.X = int32(posVec.X)
    sprite.Rect.Y = int32(posVec.Y)
    env.renderEngine.Blit(sprite)
}

func Flip(env Environment) {
    env.renderEngine.Flip()
}
