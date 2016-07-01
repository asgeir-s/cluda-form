package database_test

import (
	"testing"
	"time"

	"github.com/cluda/cluda-form/functions/post-form/database"
	"github.com/cluda/cluda-form/functions/post-form/types"
)

const submissionTable = "test-submission-table"
const formOrigin = "test.com"
const formID = "test-id"
const data = "data2=test&data2=test2"

var timestamp = time.Now().UnixNano()

func TestAddFormSubmission(t *testing.T) {

	sumbmission := types.Submission{
		FormID:     formID,     // primary key (part 2)
		FormOrigin: formOrigin, // primary key (part 1)
		Timestamp:  timestamp,  // range key
		Data:       data,
	}

	err := database.AddFormSubmission(dynamo, submissionTable, sumbmission)
	if err != nil {
		t.Fatal("Could not add sumbmission to database. Error:", err)
	}

	err = database.AddFormSubmission(dynamo, submissionTable, sumbmission)
	if err == nil {
		t.Fatal("should not be possible to add the same sumbmission again")
	}
}
