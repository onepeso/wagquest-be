package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// StringArray type represents an array of strings
type StringArray []string

// Value implements the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}
	// Escape double quotes in each element and join them with commas
	escaped := make([]string, len(a))
	for i, v := range a {
		escaped[i] = fmt.Sprintf(`"%s"`, strings.ReplaceAll(v, `"`, `\"`))
	}
	return fmt.Sprintf("{%s}", strings.Join(escaped, ",")), nil
}

// Scan implements the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	stringValue, ok := value.(string)
	if !ok {
		return nil // Return nil if the value is not a string
	}
	// Remove surrounding curly braces and split the string into elements
	stringValue = strings.Trim(stringValue, "{}")
	elements := strings.Split(stringValue, ",")
	// Trim surrounding quotes and unescape double quotes in each element
	for _, elem := range elements {
		*a = append(*a, strings.ReplaceAll(strings.Trim(elem, `"`), `\"`, `"`))
	}
	return nil
}
