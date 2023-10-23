package time

import (
	"fmt"
	"strings"
	"time"
)

const Layout = "2006-01-02 15:04:05"

// Time is a self-defined Time based on time.Time which
// unify the format to 'YYYY-MM-DD HH:mm:ss' when
// marshaling/unmarshaling using encoding/json package
type Time time.Time

// UnmarshalJSON implements the json.Unmarshaller interface.
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	var ct time.Time
	if ct, err = time.Parse(Layout, strings.Trim(string(b), `"`)); err != nil {
		return
	}
	*t = Time(ct)
	return
}

// MarshalJSON implements the json.Marshaler interface.
func (t *Time) MarshalJSON() (b []byte, err error) {
	b = []byte(t.String())
	return
}

func (t *Time) String() string {
	return fmt.Sprintf("%q", time.Time(*t).Format(Layout))
}
