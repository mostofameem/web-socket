package config

type MongoReadDB struct {
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

func (db *MongoReadDB) GetURI() string              { return db.MONGO_URI }
func (db *MongoReadDB) GetDatabase() string         { return db.Database }
func (db *MongoReadDB) GetUser() string             { return db.User }
func (db *MongoReadDB) GetPassword() string         { return db.Password }
func (db *MongoReadDB) GetReplicaSet() string       { return db.ReplicaSet }
func (db *MongoReadDB) GetAuthSource() string       { return db.AuthSource }
func (db *MongoReadDB) GetSSL() bool                { return db.SSL }
func (db *MongoReadDB) GetMaxPoolSize() uint64      { return db.MaxPoolSize }
func (db *MongoReadDB) GetMinPoolSize() uint64      { return db.MinPoolSize }
func (db *MongoReadDB) GetMaxConnIdleTimeInMs() int { return db.MaxConnIdleTimeInMs }
