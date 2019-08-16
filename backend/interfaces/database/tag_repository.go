package database

import "treasure-app/backend/domain"

type TagRepository struct {
	SqlHandler
}

func (repo *TagRepository) Store(t domain.Tag) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO tags (name) VALUES (?)", t.Name,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *TagRepository) FindById(identifier int) (tag domain.Tag, err error) {
	row, err := repo.Query("SELECT id, name FROM tags WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	tag.ID = id
	tag.Name = name
	return
}

func (repo *TagRepository) FindAll() (tags domain.Tags, err error) {
	rows, err := repo.Query("SELECT id, name FROM tags")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		tag := domain.Tag{
			ID:   id,
			Name: name,
		}
		tags = append(tags, tag)
	}
	return
}

func (repo *TagRepository) StoreShopTag(shopId int, tagId int) (err error) {
	_, err = repo.Execute(
		"INSERT INTO shop_tags (shop_id, tag_id) VALUES (?, ?)", shopId, tagId,
	)
	return err
}

func (repo *TagRepository) FindByShopId(shopId int) (tags domain.Tags, err error) {
	rows, err := repo.Query("SELECT id, name FROM tags INNER JOIN shop_tags ON shop_tags.tag_id = tags.id WHERE shop_id = ?", shopId)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		tag := domain.Tag{
			ID:   id,
			Name: name,
		}
		tags = append(tags, tag)
	}
	return
}
