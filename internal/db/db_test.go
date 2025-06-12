package db

import (
	"log"
	"os"
	"testing"
)

type testContext struct{}

const DB_TEST_FILE = "scores_test.txt"

func (c *testContext) beforeEach() {
	f, err := os.Create(DB_TEST_FILE)
	if err != nil {
		log.Fatalf("Test file couldn't be created: %v", err)
	}

	defer f.Close()

	// write the rows 10, 9 and 6 in the file
	if _, err := f.WriteString("10\n9\n6\n"); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func (c *testContext) afterEach() {
	// remove the test file
	if err := os.Remove(DB_TEST_FILE); err != nil {
		log.Fatalf("Error removing test file: %v", err)
	}
}

func testCase(test func(t *testing.T, c *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach()
		defer context.afterEach()
		test(t, context)
	}
}

func TestDB_getRows(t *testing.T) {
	t.Run("rows are properly getted", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		file := db.getRows()

		if len(file) != 3 {
			t.Errorf("Expected %v but got %v", 3, len(file))
		}
	}))
}

func TestDB_GetScores(t *testing.T) {
	t.Run("should return 3 scores as integers", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		scores := db.GetScores()

		if len(scores) != 3 {
			t.Errorf("Expected %v but got %v", 3, len(scores))
		}

		if scores[0] != 10 {
			t.Errorf("scores[0] should be %v; got %v", 10, scores[0])
		}
		if scores[1] != 9 {
			t.Errorf("scores[1] should be %v; got %v", 9, scores[1])
		}
		if scores[2] != 6 {
			t.Errorf("scores[2] should be %v; got %v", 6, scores[2])
		}
	}))
}

func TestDB_SaveScore(t *testing.T) {
	t.Run("score is saved correctly", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		db.SaveScore(5)
		scores := db.GetScores()

		if len(scores) != 4 {
			t.Errorf("Expected %v but got %v", 4, len(scores))
		}

		if scores[3] != 5 {
			t.Errorf("Expected %v but got %v", 5, scores[3])
		}
	}))

}

func TestDB_NewDB(t *testing.T) {
	t.Run("test file is used when testing", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		if DB_TEST_FILE != db.file {
			t.Errorf("Expected %v but got %v", DB_TEST_FILE, db.file)
		}
	}))
}

func TestDB_cleanScores(t *testing.T) {
	t.Run("scores are sorted after adding a score", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		db.SaveScore(8)
		db.SaveScore(5)
		scores := db.GetScores()

		if scores[2] != 8 {
			t.Errorf("Expected %v but got %v", 8, scores[2])
		}

		if scores[4] != 5 {
			t.Errorf("Expected %v but got %v", 5, scores[4])
		}
	}))

	t.Run("We should only keep up to 10 scores", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		for range 10 {
			db.SaveScore(1)
		}
		scores := db.GetScores()

		if len(scores) != 10 {
			t.Errorf("Expected %v but got %v", 10, len(scores))
		}
	}))
}
