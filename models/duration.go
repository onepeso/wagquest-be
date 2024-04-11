package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Duration time.Duration

func (d Duration) Value() (driver.Value, error) {
    dur := time.Duration(d)
    return fmt.Sprintf("%dh%dm%ds", int(dur.Hours()), int(dur.Minutes())%60, int(dur.Seconds())%60), nil
}

func (d *Duration) Scan(value interface{}) error {
    dur, err := time.ParseDuration(value.(string))
    if err != nil {
        return err
    }
    *d = Duration(dur)
    return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
    dur := time.Duration(d)
    return []byte(fmt.Sprintf(`"%dh%dm%ds"`, int(dur.Hours()), int(dur.Minutes())%60, int(dur.Seconds())%60)), nil
}

func (d *Duration) UnmarshalJSON(data []byte) error {
    s := strings.Trim(string(data), `"`)
    dur, err := time.ParseDuration(s)
    if err != nil {
        return err
    }
    *d = Duration(dur)
    return nil
}