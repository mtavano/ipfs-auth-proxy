package database

func (s *Storage) PutApiKey(apiKey *ApiKey) {
	s.apiKeys.Store(apiKey.Key, apiKey)
	return
}
