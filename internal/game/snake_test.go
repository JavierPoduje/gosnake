package game

import (
	"reflect"
	"testing"
)

func TestSnake_MoveUp(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{0, 1},
			{0, 2},
			{0, 3},
		},
		Dir: Up,
	}
	expectedBody := []Coord{
		{0, 0},
		{0, 1},
		{0, 2},
	}

	err := snake.Move(Up)
	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if snake.Dir != Up {
		t.Errorf("Expected %v but got %v", Up, snake.Dir)
	}
}

func TestSnake_MoveRight(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{1, 1},
			{1, 2},
			{1, 3},
		},
		Dir: Up,
	}
	expectedBody := []Coord{
		{2, 1},
		{1, 1},
		{1, 2},
	}

	err := snake.Move(Right)
	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if snake.Dir != Right {
		t.Errorf("Expected %v but got %v", Right, snake.Dir)
	}
}

func TestSnake_MoveDown(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		Dir: Up,
	}

	err := snake.Move(Down)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}

func TestSnake_MoveLeft(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{1, 1},
			{1, 2},
			{1, 3},
		},
		Dir: Up,
	}
	expectedBody := []Coord{
		{0, 1},
		{1, 1},
		{1, 2},
	}

	err := snake.Move(Left)
	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if snake.Dir != Left {
		t.Errorf("Expected %v but got %v", Left, snake.Dir)
	}
}

func TestSnake_MoveAfterEatingApple(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{1, 1},
			{1, 2},
			{1, 3},
			{1, 4},
		},
		Dir: Up,
	}
	expectedBody := []Coord{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
	}

	snake.Add()
	err := snake.Move(Up)

	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}

func TestSnake_MoveSmallSnakeAfterEatingApple(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{1, 1},
		},
		Dir: Up,
	}
	expectedBody := []Coord{
		{2, 1},
		{1, 1},
	}

	snake.Add()
	err := snake.Move(Right)

	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}

func TestSnake_EatTwoApplesBackToBack(t *testing.T) {
	snake := Snake{
		Body: []Coord{
			{1, 1},
		},
		Dir: Right,
	}
	expectedBody := []Coord{
		{2, 1},
		{1, 1},
	}

	snake.Add()
	err := snake.Move(Right)

	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	expectedBody = []Coord{
		{3, 1},
		{2, 1},
		{1, 1},
	}

	snake.Add()
	err = snake.Move(Right)

	if !reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}
