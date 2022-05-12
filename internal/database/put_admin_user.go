package database

func (st *Storage) PutAdmin(admin *Admin) {
	st.adminUsers.Store(admin.User, admin)
	return
}
