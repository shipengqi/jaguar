package db

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Options struct {
	Host                  string           `json:"host,omitempty"                     mapstructure:"host"`
	Username              string           `json:"username,omitempty"                 mapstructure:"username"`
	Password              string           `json:"-"                                  mapstructure:"password"`
	Database              string           `json:"database"                           mapstructure:"database"`
	MaxIdleConnections    int              `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
	MaxOpenConnections    int              `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
	LogLevel              int              `json:"log-level"                          mapstructure:"log-level"`
	AutoMigrate           bool             `json:"auto-migrate"                       mapstructure:"auto-migrate"`
	Debug                 bool             `json:"debug"                              mapstructure:"debug"`
	MaxConnectionLifeTime time.Duration    `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	Logger                logger.Interface `json:"-"`
}

func NewOptions() *Options {
	return &Options{
		Host:                  "127.0.0.1:3306",
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		LogLevel:              0,
		MaxConnectionLifeTime: 10 * time.Second,
	}
}

// Validate verifies flags passed to MySQLOptions.
func (o *Options) Validate() []error {
	return nil
}

// AddFlags adds flags related to mysql storage for a specific APIServer to the specified FlagSet.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host,
		"MySQL service host address. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&o.Username, "mysql.username", o.Username,
		"Username for access to mysql service.")

	fs.StringVar(&o.Password, "mysql.password", o.Password,
		"Password for access to mysql, should be used pair with password.")

	fs.StringVar(&o.Database, "mysql.database", o.Database,
		"Database name for the server to use.")

	fs.IntVar(&o.MaxIdleConnections, "mysql.max-idle-connections", o.MaxOpenConnections,
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&o.MaxOpenConnections, "mysql.max-open-connections", o.MaxOpenConnections,
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&o.MaxConnectionLifeTime, "mysql.max-connection-life-time", o.MaxConnectionLifeTime,
		"Maximum connection life time allowed to connect to mysql.")

	fs.BoolVar(&o.AutoMigrate, "mysql.auto-migrate", o.AutoMigrate,
		"Enables the auto migration for database models.")

	fs.BoolVar(&o.Debug, "mysql.debug", o.Debug,
		"Gorm debug mode.")

	fs.IntVar(&o.LogLevel, "mysql.log-level", o.LogLevel,
		"Specify gorm log level.")
}

// New creates a MySQL or PostGre db and retry connection when has error.
func New(opts *Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   opts.Logger,
		DisableForeignKeyConstraintWhenMigrating: true, // Whether to disable foreign key constraints
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	// sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	// sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	return db, nil
}
