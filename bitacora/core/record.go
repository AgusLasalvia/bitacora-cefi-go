package core

type Record struct {
	ID            int    `sql:"id"`
	Name          string `sql:"name"`
	Lab           string `sql:"lab"`
	Equipment     string `sql:"equipment"`
	StartDateTime int64  `sql:"startDateTime"`
	EndDateTime   int64  `sql:"endDateTime"`
	Received      int    `sql:"received"`
	Returned      int    `sql:"returned"`
	Comments      string `sql:"comments"`
	Timestamp     int64  `sql:"timestamp"`
}

// I use int64 because i want to store the dates+hours in Unix format
// After a Select for excample I convert that number to readeable DateTime format
