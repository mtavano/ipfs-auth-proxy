package database

func (st *Storage) DisableApiKey(admin *Admin, key string) error {
	value, ok := st.apiKeys.Load(key)
	if !ok {
		return ErrNotFound
	}

	// disable api key
	apiKey := value.(*ApiKey)
	apiKey.Enabled = false
	st.apiKeys.Store(key, apiKey)

	// update admin stats
	admin.DisabledApiKeys += 1
	st.adminUsers.Store(admin.User, admin)

	return nil
}
