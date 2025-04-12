package config

type MongoWriteDB struct {
	MONGO_URI           string `mapstructure:"MONGO_URI"                       validate:"required"`
	Database            string `mapstructure:"MONGO_DATABASE"                  validate:"required"`
	User                string `mapstructure:"MONGO_USER"`
	Password            string `mapstructure:"MONGO_PASSWORD"`
	ReplicaSet          string `mapstructure:"MONGO_REPLICA_SET"`
	AuthSource          string `mapstructure:"MONGO_AUTH_SOURCE"`
	SSL                 bool   `mapstructure:"MONGO_SSL"`
	MaxPoolSize         uint64 `mapstructure:"MONGO_MAX_POOL_SIZE"`
	MinPoolSize         uint64 `mapstructure:"MONGO_MIN_POOL_SIZE"`
	MaxConnIdleTimeInMs int    `mapstructure:"MONGO_MAX_CONN_IDLE_TIME_MS"`
}

func (db *MongoWriteDB) GetURI() string              { return db.MONGO_URI }
func (db *MongoWriteDB) GetDatabase() string         { return db.Database }
func (db *MongoWriteDB) GetUser() string             { return db.User }
func (db *MongoWriteDB) GetPassword() string         { return db.Password }
func (db *MongoWriteDB) GetReplicaSet() string       { return db.ReplicaSet }
func (db *MongoWriteDB) GetAuthSource() string       { return db.AuthSource }
func (db *MongoWriteDB) GetSSL() bool                { return db.SSL }
func (db *MongoWriteDB) GetMaxPoolSize() uint64      { return db.MaxPoolSize }
func (db *MongoWriteDB) GetMinPoolSize() uint64      { return db.MinPoolSize }
func (db *MongoWriteDB) GetMaxConnIdleTimeInMs() int { return db.MaxConnIdleTimeInMs }
