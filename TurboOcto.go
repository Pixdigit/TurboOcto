package TurboOcto

import (
    "github.com/Pixdigit/TurboOcto/internal/renderEngine"
    "github.com/Pixdigit/TurboOcto/internal/sharedStructs"
    "github.com/pkg/errors"
)

func Init() {
    renderEngine.Init()
}


func Blit(sprite sharedStructs.Sprite, env Environment) (error){
    posVec := sprite.Body.GetPosition()
    sprite.Rect.X = int32(posVec.X)
    sprite.Rect.Y = int32(posVec.Y)
    err := env.renderEngine.Blit(sprite)
    if err != nil {return errors.Wrap(err, "Error while blitting " + sprite.Name)}
    return nil
}


func Flip(env Environment) (error){
    return env.renderEngine.Flip()
}
