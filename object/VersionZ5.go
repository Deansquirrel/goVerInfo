package object

import (
	"time"
)

type VersionZ5 struct {
	ObjectName    string    `toml:"objectName"`
	ObjectType    string    `toml:"objectType"`
	ObjectVersion string    `toml:"objectVersion"`
	ObjectDate    time.Time `toml:"objectDate"`
}
