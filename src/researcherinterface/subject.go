package main

import (
	"strings"
)

// Subject represnts the subject being captured by a dataset. A dataset can have
// many of these
type Subject struct {
	ID        int64
	DatasetID int64
	Target    string
	Dummies   []string
}

func NewSubject(datasetID int64, target string, dummies []string) (s *Subject, err error) {
	var id int64

	d := "{" + strings.Join(dummies, ",") + "}"

	err = DB.QueryRow("INSERT INTO subjects (ds_id, target, dummies) VALUES ($1, $2, $3) RETURNING id", datasetID, target, d).Scan(&id)
	if err != nil {
		return
	}

	s = &Subject{
		ID:        id,
		DatasetID: datasetID,
		Target:    target,
		Dummies:   dummies,
	}

	return
}
