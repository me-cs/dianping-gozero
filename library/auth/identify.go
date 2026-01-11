package auth

type User interface {
	GetUid() int64
}

type user struct {
	Uid int64 `json:"uid"`
}

func (m *user) GetUid() int64 {
	return m.Uid
}

func NewUser(uid int64) User {
	return &user{
		Uid: uid,
	}
}
