domain	Entities ビジネスルールのためのデータ
infrastructure	Frameworks & Drivers
interfaces	Interface
usecase	Use cases

DB接続には、外部パッケージを使用しているので、infrastructure層に定義し外側のルールを内側に持ち込まない

interfaces/databaseからのInputをusecase/user_repository.goで、interfaces/controllersへのGatewayをusecase/user_interactor.goで実現


## API実装

### ユーザー

#### ユーザー登録 (POST /users)
`curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name": "test", "email": "test@example.com", "password": "password"}' localhost:8080/users`

#### ユーザー一覧 (GET /users)
`curl -i -H 'Content-Type:application/json' localhost:8080/users`

#### ユーザー詳細 (GET /users/{:id})
`curl -i -H 'Content-Type:application/json' localhost:8080/users/1`


### ショップ
#### ショップ登録 (POST /shops)
`curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name": "test2", "address": "京都", "tel": "090-1234-5678"}' localhost:8080/shops`

#### ショップ一覧 (GET /shops)
`curl -i -H 'Content-Type:application/json' localhost:8080/shops`

#### ショップ詳細 (GET /shops/{:id})
`curl -i -H 'Content-Type:application/json' localhost:8080/shops/1`

### いいね
#### いいね登録 (POST /users/like)
`curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"user_id": 1, "shop_id": 1}' localhost:8080/users/like`

### いいね削除 (DELETE /users/like)
`curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE -d '{"user_id": 1, "shop_id": 1}' localhost:8080/users/like`
