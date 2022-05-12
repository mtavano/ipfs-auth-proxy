package database

func (st *Storage) GetApiKey(key string) (*ApiKey, error) {
	value, exist := st.apiKeys.Load(key)
	if !exist {
		return nil, ErrNotFound
	}

	apiKey := value.(*ApiKey)
	return apiKey, nil
}
