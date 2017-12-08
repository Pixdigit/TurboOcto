package TurboOcto

var DefaultConfiguration = Configuration{0, 0, false, 0, 10, 1 / 60.0, 8, 3}

type Configuration struct {
    ScreenWidth, ScreenHeight int32
    Fullscreen bool
    XGravitation, YGravitation float64
    Timestep float64
    PositionAccuracy, VelocityAccuracy int
}
