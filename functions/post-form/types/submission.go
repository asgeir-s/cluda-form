package types

type Submission struct {
	FormID     string // primary key (part 2)
	FormOrigin string // primary key (part 1)
	Timestamp  int64  // range key
	Data       string
}
