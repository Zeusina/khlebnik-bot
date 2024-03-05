package utils

var fsm map[int64]string

func InitFSM() {
	fsm = map[int64]string{}
}

func GetFSM(userId int64) string {
	return fsm[userId]
}

func SetFSM(value string, userId int64) {
	fsm[userId] = value
}
