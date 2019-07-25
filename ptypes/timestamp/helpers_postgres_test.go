package timestamp

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq" // Postgres driver
)

var env = &envConfig{}

type envConfig struct {
	PostgresUser     string `default:"root" split_words:"true"`
	PostgresPassword string `default:"password" split_words:"true"`
	PostgresHost     string `default:"127.0.0.1" split_words:"true"`
	PostgresPort     string `default:"5432" split_words:"true"`
	PostgresName     string `default:"proto" split_words:"true"`
}

func init() {
	if err := envconfig.Process("", env); err != nil {
		panic(err)
	}
}

var postgresDB *sqlx.DB
var postgresSetup sync.Once

func getPostgresDB() *sqlx.DB {
	postgresSetup.Do(func() {

		dbConfig := postgresConfig(env.PostgresUser, env.PostgresPassword, env.PostgresName, env.PostgresHost, env.PostgresPort)

		postgresDBConn, err := sql.Open("postgres", dbConfig)
		if err != nil {
			panic(err)
		}

		// `Open` does not directly open connection, so ensure connection established
		if err = postgresDBConn.Ping(); err != nil {
			panic(err)
		}

		postgresDB = sqlx.NewDb(postgresDBConn, "postgres")

		postgresDB.MustExec(
			`create table if not exists times (
				id serial not null constraint times_pk primary key,
				ts timestamp not null
			)`)
	})
	return postgresDB
}

// postgresConfig sets up our postgres client configuration
func postgresConfig(user, password, schemaName, host, port string) string {
	psqlInfo := fmt.Sprintf(`
			host=%s
			port=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=disable`,
		host,
		port,
		user,
		password,
		schemaName)

	return psqlInfo
}

type Message struct {
	ID int64      `db:"id"`
	TS *Timestamp `db:"ts"`
}

func TestPostgresScanValue(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	db := getPostgresDB()

	now := time.Now().Truncate(time.Microsecond).UTC()

	message := &Message{
		ID: 3,
		TS: &Timestamp{},
	}
	_, err := message.TS.SetToTime(now)
	if err != nil {
		t.Fatal(err)
	}

	if now.Format(time.RFC3339Nano) != message.TS.RFC3339() {
		t.Fatal(now.Format(time.RFC3339Nano), message.TS.RFC3339())
	}

	t.Log(message, message.TS.RFC3339())

	_, _ = db.Exec(`delete from times where id = $1`, message.ID)
	_, err = db.Exec(`insert into times (id, ts) values ($1, $2)`, message.ID, message.TS)
	if err != nil {
		t.Fatal(err)
	}

	result := &Message{}
	err = db.Get(result, `select * from times where id = $1`, message.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !proto.Equal(message.TS, result.TS) {
		t.Fatal(result, result.TS.RFC3339())
	}

	tt, err := result.TS.Time()
	if err != nil {
		t.Fatal(err)
	}

	if !tt.Equal(now) {
		t.Fatal(tt)
	}
}
