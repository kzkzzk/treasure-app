package database

import "treasure-app/backend/domain"

type LikeRepository struct {
	SqlHandler
}

func (repo *LikeRepository) Store(l domain.Like) (rl domain.Like, err error) {
	_, err = repo.Execute(
		"INSERT INTO likes (user_id, shop_id) VALUES (?,?)", l.UserID, l.ShopID,
	)

	if err != nil {
		return
	}
	rl = l
	return
}

func (repo *LikeRepository) Destroy(l domain.Like) (err error) {
	_, err = repo.Execute(
		"DELETE FROM likes WHERE user_id = ? AND shop_id = ?", l.UserID, l.ShopID,
	)
	return err
}
