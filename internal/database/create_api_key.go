package database

func (st *Storage) CreateApiKey(admin *Admin, apiKey *ApiKey) {
	st.adminUsers.Store(admin.User, &Admin{
		User:            admin.User,
		password:        admin.password,
		DisabledApiKeys: admin.DisabledApiKeys,
		ApiKeys:         admin.ApiKeys + 1,
	})

	st.apiKeys.Store(apiKey.Key, apiKey)
	return
}
