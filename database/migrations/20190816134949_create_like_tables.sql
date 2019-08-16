-- +goose Up
-- +goose StatementBegin
CREATE TABLE likes (
    user_id INT NOT NULL,
    shop_id INT NOT NULL,
    PRIMARY KEY(user_id, shop_id),
    CONSTRAINT fk_likes_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_likes_shop_id FOREIGN KEY (shop_id) REFERENCES shops (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE likes;
-- +goose StatementEnd
