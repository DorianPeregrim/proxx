package proxx

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrWrongValue = errors.New("wrong value")
var ErrTooMuchHoles = errors.New("too much holes")

type Application struct {
	isGameOver bool
}

func (a *Application) Start() error {
	bc, err := a.askBoardConfig()
	if err != nil {
		return err
	}

	return a.startGame(bc)
}

func NewApplication() *Application {
	return &Application{
		isGameOver: false,
	}
}

func (a *Application) askBoardConfig() (BoardConfig, error) {
	val := inputPrompt(fmt.Sprintf("Rows Num [1 - %d]:", maxRowsNum))
	rowsNum, _ := strconv.Atoi(val)
	if rowsNum == 0 || rowsNum > maxRowsNum {
		return BoardConfig{}, ErrWrongValue
	}

	val = inputPrompt(fmt.Sprintf("Columns Num [1 - %d]:", maxColsNum))
	colsNum, _ := strconv.Atoi(val)
	if colsNum == 0 || colsNum > maxColsNum {
		return BoardConfig{}, ErrWrongValue
	}

	val = inputPrompt(fmt.Sprintf("Black holes Num [1 - %d]:", rowsNum*colsNum))
	holesNum, _ := strconv.Atoi(val)
	if holesNum == 0 {
		return BoardConfig{}, ErrWrongValue
	} else if holesNum > rowsNum*colsNum {
		return BoardConfig{}, ErrTooMuchHoles
	}

	bc := BoardConfig{
		RowsNum:       rowsNum,
		ColsNum:       colsNum,
		BlackHolesNum: holesNum,
	}

	return bc, nil
}

func (a *Application) startGame(bc BoardConfig) error {
	b := NewBoard(bc)
	br := NewConsoleBoardRenderer()
	isComplete := false

	printBoard(br, b)

	for !a.isGameOver {
		val := inputPrompt("Input row and column indexes through a space:")
		i, j := extractIndexes(val)
		if i < 0 || i >= b.config.RowsNum || j < 0 || j >= b.config.ColsNum {
			return ErrWrongValue
		}

		isHole := b.Click(i, j)
		isComplete = b.IsComplete()
		if !isHole || isComplete {
			a.isGameOver = true
			b.Open()
		}
		printBoard(br, b)
	}

	gameOver(isComplete)

	return nil
}

func printBoard(br *ConsoleBoardRenderer, b *Board) {
	fmt.Println()
	fmt.Println(br.Render(b))
}

func gameOver(isComplete bool) {
	if isComplete {
		fmt.Println("   YOU WIN!!!    ")
	} else {
		fmt.Println("   GAME OVER    ")
	}
	fmt.Println()
}

func extractIndexes(str string) (int, int) {
	indexes := strings.Split(str, " ")
	if len(indexes) < 2 {
		return 0, 0
	}

	i, _ := strconv.Atoi(indexes[0])
	j, _ := strconv.Atoi(indexes[1])
	return i - 1, j - 1
}
