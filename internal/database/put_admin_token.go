package database

func (st *Storage) PutAdminToken(token string, admin *Admin) {
	st.adminTokens.Store(token, admin)
	return
}
