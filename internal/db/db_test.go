package db

import (
	"testing"
)

func TestDB_getRows(t *testing.T) {
	db := NewDB()
	file := db.getRows()

	if 3 != len(file) {
		t.Errorf("Expected %v but got %v", 3, len(file))
	}
}

func TestDB_GetScores(t *testing.T) {
	db := NewDB()
	scores := db.GetScores()

	if 3 != len(scores) {
		t.Errorf("Expected %v but got %v", 3, len(scores))
	}
}

func TestDB_SaveScore(t *testing.T) {
	db := NewDB()
	db.SaveScore(5)

	scores := db.GetScores()

	if 4 != len(scores) {
		t.Errorf("Expected %v but got %v", 4, len(scores))
	}

	if 5 != scores[3] {
		t.Errorf("Expected %v but got %v", 5, scores[3])
	}

	db.deleteLastRow()
}

func TestDB_NewDB(t *testing.T) {
	db := NewDB()

	if DB_TEST_FILE != db.file {
		t.Errorf("Expected %v but got %v", DB_TEST_FILE, db.file)
	}
}
