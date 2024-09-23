package models

type StatusOrder string

const (
	ORDER_NEW     StatusOrder = "N"
	ORDER_FINISH  StatusOrder = "F"
	ORDER_PROCESS StatusOrder = "P"
)

type StatusNotif string

const (
	NOTIF_NEW    StatusNotif = "N"
	NOTIF_CLOSED StatusNotif = "C"
)

type TipeNotif string

const (
	NOTIF_INFO  TipeNotif = "I"
	NOTIF_ORDER TipeNotif = "O"
)
