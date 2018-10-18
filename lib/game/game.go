package game

// Game holds details of a version and configuration of a game
type Game struct {
	Path     string
	GameData string
}

// Executable is the path to the game executable
func (g Game) Executable() string {
	return g.Path
}

// WorkingDir is the location of all of the configuration, mods and json stuff
func (g Game) WorkingDir() string {
	return g.GameData
}
