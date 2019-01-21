# About this repo
This repo is a collection for small gadgets.

## Dependence
1. This repo is written in [Go](https://golang.org/)
1. Use `https://github.com/veandco/go-sdl2`(SDL2 binding for Go) to provide GUI. `SDL2` is needed, [detail intallation guidence](https://github.com/veandco/go-sdl2#installation)


## Conway's Life Game

### Description
> Conway's Game of Life, also known as the Game of Life or simply Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970. It is the best-known example of a cellular automaton.    
The "game" is actually a zero-player game, meaning that its evolution is determined by its initial state, needing no input from human players. One interacts with the Game of Life by creating an initial configuration and observing how it evolves.    
-- [Wikipedia](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)


## Ball Collision

### Description
Stimulate a rigid ball given an origin speed falling, it will bounce once it collides with any boundary around.    
Each time it make collision, the speed will be decrease and turn to an opposite direction. 