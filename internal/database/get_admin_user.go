package database

func (st *Storage) GetAdmin(adminUser, pass string) (*Admin, error) {
	value, ok := st.adminUsers.Load(adminUser)
	if !ok {
		return nil, ErrNotFound
	}

	admin := value.(*Admin)
	if admin.password != pass {
		return nil, ErrNotFound
	}
	return admin, nil
}
