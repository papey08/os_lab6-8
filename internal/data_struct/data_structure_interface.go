package data_struct

type DataStructure interface {

	// GetTime returns number t of seconds since timer at node id started
	GetTime(id int) (int, error)

	// StartTimer starts timer at node id
	StartTimer(id int) error

	// PauseTimer sets timer at node id on pause
	PauseTimer(id int) error

	// InsertNode adds new node with ID into data structure with
	InsertNode(id int) error

	// DeleteNode deletes node from data structure by id
	DeleteNode(id int) error

	// ResetNode stops the timer at node id and sets it's time as 0
	ResetNode(id int) error

	// Length returns amount of nodes in data structure
	Length() int
}
