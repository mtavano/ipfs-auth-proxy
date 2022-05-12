package database

func (st *Storage) GetAdminByToken(token string) (*Admin, error) {
	value, exist := st.adminTokens.Load(token)
	if !exist {
		return nil, ErrNotFound
	}

	admin := value.(*Admin)
	return admin, nil
}
