package db

import (
	"log"
	"os"
	"testing"
)

type testContext struct {
	//test    *func(t *testing.T, c *testContext)
	//context int64
}

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

		if 3 != len(file) {
			t.Errorf("Expected %v but got %v", 3, len(file))
		}
	}))
}

func TestDB_GetScores(t *testing.T) {
	t.Run("should return 3 scores as integers", testCase(func(t *testing.T, c *testContext) {
		db := NewDB()
		db.file = DB_TEST_FILE

		scores := db.GetScores()

		if 3 != len(scores) {
			t.Errorf("Expected %v but got %v", 3, len(scores))
		}

		if 10 != scores[0] {
			t.Errorf("scores[0] should be %v; got %v", 10, scores[0])
		}
		if 9 != scores[1] {
			t.Errorf("scores[1] should be %v; got %v", 9, scores[1])
		}
		if 6 != scores[2] {
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

		if 4 != len(scores) {
			t.Errorf("Expected %v but got %v", 4, len(scores))
		}

		if 5 != scores[3] {
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
