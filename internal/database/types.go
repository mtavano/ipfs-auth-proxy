package database

type ApiKey struct {
	Key             string
	Enabled         bool
	Requests        uint64
	BytesTransfered int64
}

type Admin struct {
	User            string
	password        string
	ApiKeys         uint64
	DisabledApiKeys uint64
}
