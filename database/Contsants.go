package database

import "encoding/binary"

const Uncompleted = "uncompleted_tasks"
const Completed = "completed_tasks"

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
