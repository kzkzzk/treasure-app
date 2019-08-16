package database

import "treasure-app/backend/domain"

type ShopRepository struct {
	SqlHandler
}

func (repo *ShopRepository) Store(s domain.Shop) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO shops (name, address, tel, image) VALUES (?,?,?,?)", s.Name, s.Address, s.Tel, s.Image,
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

func (repo *ShopRepository) FindById(identifier int) (shop domain.Shop, err error) {
	row, err := repo.Query("SELECT id, name, address, tel, image FROM shops WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	var address string
	var tel string
	var image string
	row.Next()
	if err = row.Scan(&id, &name, &address, &tel, &image); err != nil {
		return
	}
	shop.ID = id
	shop.Name = name
	shop.Address = address
	shop.Image = image
	shop.Tel = tel
	return
}

func (repo *ShopRepository) FindAll() (shops domain.Shops, err error) {
	rows, err := repo.Query("SELECT id, name, address, tel, image FROM shops")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		var address string
		var tel string
		var image string
		if err := rows.Scan(&id, &name, &address, &tel, &image); err != nil {
			continue
		}
		shop := domain.Shop{
			ID:      id,
			Name:    name,
			Address: address,
			Tel:     tel,
			Image:   image,
		}
		shops = append(shops, shop)
	}
	return
}
