package db

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	DB_FILE                  = "scores.txt"
	NUMBER_OF_SCORES_TO_KEEP = 10
)

type DB struct {
	file string
}

func NewDB() DB {
	return DB{file: DB_FILE}
}

func (db DB) GetScores() []int {
	rows := db.getRows()
	scores := make([]int, 0, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		score, err := strconv.Atoi(row)
		if err != nil {
			log.Fatalf("Error converting row to int: %s", err)
		}

		scores = append(scores, score)
	}

	return scores
}

func (db DB) SaveScore(score int) {
	f, err := os.OpenFile(db.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer f.Close()

	if _, err := f.WriteString(strconv.Itoa(score) + "\n"); err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	db.cleanScores()
}

func (db DB) cleanScores() {
	scores := db.GetScores()
	sort.Ints(scores)

	var scoresToKeep []int
	for i := len(scores); i > 0; i-- {
		scoresToKeep = append(scoresToKeep, scores[i-1])
		if len(scoresToKeep) == NUMBER_OF_SCORES_TO_KEEP {
			break
		}
	}

	// write the cleaned scores to the file
	f, err := os.OpenFile(db.file, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer f.Close()

	for _, score := range scoresToKeep {
		if _, err := f.WriteString(strconv.Itoa(score) + "\n"); err != nil {
			log.Fatalf("Error writing to file: %s", err)
		}
	}
}

func (db DB) getRows() []string {
	dat, err := os.ReadFile(db.file)
	if err != nil {
		// if the file doesn't exist, create it
		_, err = os.Create(db.file)
		if err != nil {
			log.Fatalf("Test file couldn't be created: %v", err)
		}
		dat = []byte("")
	}
	rows := strings.Split(string(dat), "\n")
	trimmedRows := make([]string, 0, len(rows))
	for _, row := range rows {
		trimmedRow := strings.TrimSpace(row)
		if trimmedRow == "" {
			continue
		}
		trimmedRows = append(trimmedRows, trimmedRow)
	}

	return trimmedRows

}
