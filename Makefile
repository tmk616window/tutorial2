HASURA_PORT := 8080
HASURA_CONSOLE_PORT := 9695
DB_USER=hasura
DB_PASSWORD=secret
DB_NAME=postgres

# .env をロードして上記デフォルト値と一致する環境変数をオーバーライドする
ifneq (,$(wildcard ./.env))
	include .env
	export
	unexport HASURA_GRAPHQL_JWT_SECRET
endif

.PHONY: apply
apply: ## マイグレーションを適用する
	@until curl -s -o /dev/null http://127.0.0.1:${HASURA_PORT}/healthz; do sleep 1; done
	hasura --project hasura deploy --admin-secret secret --skip-update-check

.PHONY: up
up: bridge ## バックグラウンドで起動
	@docker compose up ${SERVICES} -d ${OPTS}

.PHONY: down
down: bridge ## コンテナ停止
	@docker compose down

.PHONY: console
console: ## hasura console
	@hasura --project hasura console --admin-secret secret --console-port ${HASURA_CONSOLE_PORT} --skip-update-check

.PHONY: bridge
bridge: ## docker network を生成する
	@if [ -z "$$(docker network ls -q -f name='^api-network$\')" ]; then docker network create api-network; fi

.PHONY: sqlboiler
sqlboiler: ## SQLBoiler コード生成
	docker compose run --rm sidecar sqlboiler psql

.PHONY: gqlgen
gqlgen: ## sidecar/graph/schema に書いた GraphQL スキーマをもとに Go のコードを生成する
	docker compose run --rm sidecar gqlgen
