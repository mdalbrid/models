package model

import (
	"encoding/json"
	"strings"
	"time"
)

const NullUUID UUID = `00000000-0000-0000-0000-000000000000`

// UUID - тип UUID  объекта
type UUID string

func (u UUID) String() string {
	return string(u)
}

// JSONDate -
type JSONDate time.Time

// UnmarshalJSON Implement Marshaller and Unmarshaler interface
func (j *JSONDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("02.01.2006", s)
	if err != nil {
		return err
	}
	*j = JSONDate(t)
	return nil
}

// MarshalJSON -
func (j JSONDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Format  Maybe a Format function for printing your date
func (j JSONDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
