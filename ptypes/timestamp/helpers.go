package timestamp

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// XXX_WellKnownType allows this non-Google timestamp protobuf to interact properly with jsonpb
// and other libraries as it is was the standard WKT timestamp.
func (*Timestamp) XXX_WellKnownType() string { return "Timestamp" }

// Scan implements the Scanner interface of the database driver
func (m *Timestamp) Scan(value interface{}) error {

	// initialize timestamp if pointer is nil
	if m == nil {
		*m = Timestamp{}
	}

	// convert the interface to a time type
	dbTime, ok := value.(time.Time)

	fmt.Printf("%T\n%+#v\n%v\n%v\n", value, value, dbTime, ok)

	//if ok {
	//	m.Milliseconds = dbTime.UnixNano() / 1000 / 1000
	//	m.IsNotNull = true
	//	return nil
	//}
	//
	//m.Milliseconds = 0
	//m.IsNotNull = false
	return nil
}

// Value implements the db driver Valuer interface
func (m Timestamp) Value() (driver.Value, error) {
	//return time.Unix(0, m.Milliseconds*1000*1000), nil
	return nil, nil
}
