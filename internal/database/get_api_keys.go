package database

func (st *Storage) GetApiKeys() ([]*ApiKey, error) {
	apiKeys := make([]*ApiKey, 0)

	st.apiKeys.Range(func(key interface{}, value interface{}) bool {
		apiKey := value.(*ApiKey)
		apiKeys = append(apiKeys, apiKey)
		return true
	})
	return apiKeys, nil
}
