package utils

import log "github.com/sirupsen/logrus"

var fsm map[int64]string

func InitFSM() {
	fsm = map[int64]string{}
}

func GetFSM(userId int64) string {
	log.WithField("userid", userId).Debug("Get FSM state for user")
	return fsm[userId]
}

func SetFSM(value string, userId int64) {
	log.WithFields(log.Fields{
		"userid":   userId,
		"newstate": value,
	}).Debug("Set new FSM state for user")
	fsm[userId] = value
}
