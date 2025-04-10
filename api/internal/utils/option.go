package utils

import (
	"database/sql"

	"github.com/mcucu/mns-ecommerce/config"
)

type Options struct {
	Config config.IConfig
	DB     *sql.DB
}
