package appcontexts

import (
	"database/sql"

	"github.com/mcucu/mns-ecommerce/config"
	"github.com/mcucu/mns-ecommerce/internal/drivers"
)

type AppContext struct {
	config config.IConfig
}

// NewAppContext initiate appcontext object
func NewAppContext(cfg config.IConfig) *AppContext {
	return &AppContext{
		config: cfg,
	}
}

func (a *AppContext) GetSQLiteConn() (*sql.DB, error) {
	return drivers.NewSQLite(
		drivers.SQLiteOption{
			URL:     a.config.GetString("DB_URL"),
			MaxIdle: a.config.GetInt("DB_MAX_IDLE"),
			MaxOpen: a.config.GetInt("DB_MAX_OPEN"),
		})
}
