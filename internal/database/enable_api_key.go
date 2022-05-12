package database

func (st *Storage) EnableApiKey(admin *Admin, key string) error {
	value, ok := st.apiKeys.Load(key)
	if !ok {
		return ErrNotFound
	}

	// enable api key
	apiKey := value.(*ApiKey)
	apiKey.Enabled = false
	st.apiKeys.Store(key, apiKey)

	// update admin stats
	admin.DisabledApiKeys -= 1
	st.adminUsers.Store(admin.User, admin)

	return nil
}
