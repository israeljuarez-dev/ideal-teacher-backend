package routers

import (
	"github.com/gorilla/mux"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/storage/postgres"

)

const apiV1 = "/api/v1"

func InitRouters(router *mux.Router, db *postgres.DB) {
	setUpUser(router, db)
}
