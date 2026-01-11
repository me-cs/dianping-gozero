package dao

import "context"

func (d *Dao) GetUserInfo(c context.Context, id int) error {
	r := d.mysql.QueryRowContext(c, "SELECT * FROM user_info WHERE id=?", id)
	return r.Err()
}
