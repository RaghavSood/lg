package types

type LogEntry struct {
	Title     string
	Content   string
	Timestamp int64
	Revision  int64
	Locations []Location
}

type Location struct {
	Fragment        string
	ResolvedAddress string
	Latitude        string
	Longitude       string
	Resolved        bool
}
