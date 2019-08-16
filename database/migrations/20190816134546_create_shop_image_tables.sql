-- +goose Up
-- +goose StatementBegin
CREATE TABLE shop_images (
    id INT NOT NULL AUTO_INCREMENT,
    shop_id INT NOT NULL,
    path VARCHAR(255) NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_shop_images_shop_id FOREIGN KEY (shop_id) REFERENCES shops (id) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shop_images;
-- +goose StatementEnd
