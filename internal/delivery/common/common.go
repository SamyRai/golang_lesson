package common

type Delivery interface {
	PlayerData() (string, string)
	BoardData() int
	MoveData() (int, int)
	ShowBoard([][]string)
	Notify(string)
}
