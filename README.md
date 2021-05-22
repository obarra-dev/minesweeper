<h1 align="center">
  <br>
  <br>
  Minesweeper
  <br>
</h1>
<h4 align="center">Game mineswipeer logic.</h4>
<p align="center">
  <a href="#dependencies">Dependencies</a> •
  <a href="#technologies">Technologies</a> •
  <a href="#how-to-check">How To Check</a> •
  <a href="#documentation">Documentation</a> •
  <a href="#contributing">Contributing</a> 
</p>

## Technologies
* [GoLang 1.16](https://golang.org/)
* [Gofmt](https://golang.org/cmd/gofmt/)
* [Golint](https://github.com/golang/lint)
* [Vet](https://golang.org/cmd/vet/)
* [Github Actions](https://github.com/features/actions)

## Dependencies
* [GoLang 1.16](https://golang.org/)
* [Golint](https://github.com/golang/lint)


## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```
go get -u github.com/obarra-dev/minesweeper@latest
```
## Examples
Let's game it out

```go
func main() {
    //start
    mines := []minesweeper.Mine{{Row: 1, Column: 1}}
    game := minesweeper.New(3, 8, mines)
    
    //play
    gameCopy := game.Play(0, 0, minesweeper.TypeMoveClean)
    
    //show game state
    switch gameCopy.State {
    case minesweeper.StateGameNew:
        fmt.Println("Game Start...")
    case minesweeper.StateGameRunning:
        fmt.Println("Running...")
    case minesweeper.StateGameLost:
        fmt.Println("Game lost...")
    case minesweeper.StateGameWon:
        fmt.Println("Game Won...")
    default:
        fmt.Println("Crash...")
    }
}

```

## How To Check

### Running the lints

Before running the project lints please ensure that all the dependencies are installed in your system. Then execute the following command:

```
make lint
```

### Running the tests

In order to run the project tests you need to execute the following command:

```
make test
```


## Documentation



## Contributing
1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Make your changes
4. Run the tests, adding new ones for your own code if necessary (`junit5`)
5. Commit your changes (`git commit -am 'Added some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new Pull Request

* Questions?, <a href="mailto:barraomar12@gmail.com?Subject=Question about Game Mineswipeer" target="_blank">write here</a>
