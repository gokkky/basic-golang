package handler

import "project-prim-aor/primdb"

type Handler struct {
	DB *primdb.DatabaseNaja
}

func InitHandler(db *primdb.DatabaseNaja) *Handler {
	return &Handler{
		db,
	}
}
