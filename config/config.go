package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Debug    bool   `json:"debug"`
	db       *gorm.DB
	lock     sync.Mutex
}

type Config struct {
	App   *App   `json:"app"`
	MySQL *MySQL `json:"mysql"`
	Log   *Log   `json:"log"`
}

// stringger
func (c *Config) String() string {
	v, _ := json.Marshal(c)
	return string(v)
}

// APP method
func (a *App) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

// DB method
func (m *MySQL) DB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.Database,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to connect database", err)
			os.Exit(1)
		}

		if m.Debug {
			db = db.Debug()
		}
		m.db = db
	}
	return m.db
}
