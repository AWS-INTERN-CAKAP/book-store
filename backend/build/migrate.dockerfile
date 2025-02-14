FROM migrate/migrate:v4.15.2

WORKDIR /migrations

COPY ./database/migrations /migrations

ENV DATABASE_URL="mysql://root:password@tcp(db:3306)/your_database?parseTime=true"

ENTRYPOINT ["sh", "-c", "migrate -database ${DATABASE_URL} -path /migrations $@"]