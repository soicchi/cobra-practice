# cobra-practice

# Initialize

```sh
docker compose build
```

```sh
docker compose up
```

# Database create log

```sh
cobra-practice-api-1  | 2024/04/28 13:44:20 /go/pkg/mod/github.com/go-gormigrate/gormigrate/v2@v2.1.2/gormigrate.go:398
cobra-practice-api-1  | [14.378ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'migrations' AND table_type = 'BASE TABLE'
cobra-practice-api-1  |
cobra-practice-api-1  | 2024/04/28 13:44:20 /go/pkg/mod/github.com/go-gormigrate/gormigrate/v2@v2.1.2/gormigrate.go:401
cobra-practice-api-1  | [0.837ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'migrations' AND table_type = 'BASE TABLE'
cobra-practice-api-1  |
cobra-practice-api-1  | 2024/04/28 13:44:20 /go/pkg/mod/github.com/go-gormigrate/gormigrate/v2@v2.1.2/gormigrate.go:401
cobra-practice-api-1  | [5.938ms] [rows:0] CREATE TABLE "migrations" ("id" varchar(255),PRIMARY KEY ("id"))
cobra-practice-api-1  |
cobra-practice-api-1  | 2024/04/28 13:44:20 /go/pkg/mod/github.com/go-gormigrate/gormigrate/v2@v2.1.2/gormigrate.go:409
cobra-practice-api-1  | [1.252ms] [rows:1] SELECT count(*) FROM "migrations" WHERE id = '20240429_add_user_table'
cobra-practice-api-1  |
cobra-practice-api-1  | 2024/04/28 13:44:20 /app/20240429_add_user_table.go:18
cobra-practice-api-1  | [4.296ms] [rows:0] CREATE TABLE "users" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" varchar(255) NOT NULL,PRIMARY KEY ("id"),CONSTRAINT "uni_users_name" UNIQUE ("name"))
cobra-practice-api-1  |
cobra-practice-api-1  | 2024/04/28 13:44:20 /app/20240429_add_user_table.go:18
cobra-practice-api-1  | [2.229ms] [rows:0] CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" ("deleted_at")
cobra-practice-api-1  |
cobra-practice-api-1  | 2024/04/28 13:44:20 /go/pkg/mod/github.com/go-gormigrate/gormigrate/v2@v2.1.2/gormigrate.go:463
cobra-practice-api-1  | [3.777ms] [rows:1] INSERT INTO "migrations" ("id") VALUES ('20240429_add_user_table')
cobra-practice-api-1  | 2024/04/28 13:44:20 Migrated database
cobra-practice-api-1  |
cobra-practice-api-1  |    ____    __
cobra-practice-api-1  |   / __/___/ /  ___
cobra-practice-api-1  |  / _// __/ _ \/ _ \
cobra-practice-api-1  | /___/\__/_//_/\___/ v4.12.0
cobra-practice-api-1  | High performance, minimalist Go web framework
cobra-practice-api-1  | https://echo.labstack.com
cobra-practice-api-1  | ____________________________________O/_______
cobra-practice-api-1  |                                     O\
cobra-practice-api-1  |
cobra-practice-api-1  | ⇨ http server started on [::]:1323
```

# Database roll back

```sh
```
