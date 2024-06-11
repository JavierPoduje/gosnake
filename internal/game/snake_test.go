package game

import (
	"reflect"
	"testing"
)

func TestSnake_Move(t *testing.T) {
	// UP
	snake := Snake{
		Body: []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		Dir: Up,
	}
	expectedBody := []Coord{
		{-1, 0},
		{0, 0},
		{1, 0},
	}

	err := snake.Move(Up)
	if reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if snake.Dir != Up {
		t.Errorf("Expected %v but got %v", Up, snake.Dir)
	}

	// RIGHT
	snake = Snake{
		Body: []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		Dir: Up,
	}
	expectedBody = []Coord{
		{0, 1},
		{0, 0},
		{1, 0},
	}

	err = snake.Move(Right)
	if reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if snake.Dir != Right {
		t.Errorf("Expected %v but got %v", Right, snake.Dir)
	}

	// Down
	snake = Snake{
		Body: []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		Dir: Up,
	}

	err = snake.Move(Down)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	// Left
	snake = Snake{
		Body: []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
		},
		Dir: Up,
	}
	expectedBody = []Coord{
		{-1, 0},
		{0, 0},
		{1, 0},
	}

	err = snake.Move(Left)
	if reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if snake.Dir != Left {
		t.Errorf("Expected %v but got %v", Left, snake.Dir)
	}

	// Move after eating apple
	snake = Snake{
		Body: []Coord{
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 0},
		},
		Dir: Up,
	}
	expectedBody = []Coord{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
	}
	err = snake.Move(Up)
	if reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	// Move after smaller snake eats apple
	snake = Snake{
		Body: []Coord{
			{0, 0},
			{0, 0},
		},
		Dir: Up,
	}
	expectedBody = []Coord{
		{-1, 0},
		{0, 0},
	}
	err = snake.Move(Up)
	if reflect.DeepEqual(snake.Body, expectedBody) {
		t.Errorf("Expected %v but got %v", expectedBody, snake.Body)
	}
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}
