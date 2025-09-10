package core

type Record struct {
	ID            int    `sql:"id"`
	Name          string `sql:"name"`
	Lab           string `sql:"lab"`
	Equipment     string `sql:"equipment"`
	StartDateTime string `sql:"startDateTime"`
	EndDateTime   string `sql:"endDateTime"`
	Received      string `sql:"received"`
	Returned      string `sql:"returned"`
	Comments      string `sql:"comments"`
	Timestamp     int    `sql:"timestamp"`
}
