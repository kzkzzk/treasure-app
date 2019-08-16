-- +goose Up
-- +goose StatementBegin
CREATE TABLE shop_tags (
    shop_id INT NOT NULL,
    tag_id INT NOT NULL,
    PRIMARY KEY(shop_id, tag_id),
    CONSTRAINT fk_shop_tags_shop_id FOREIGN KEY (shop_id) REFERENCES shops (id),
    CONSTRAINT fk_shop_tags_tag_id FOREIGN KEY (tag_id) REFERENCES tags (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shop_tags;
-- +goose StatementEnd
