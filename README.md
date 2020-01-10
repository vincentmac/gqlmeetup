# gqlgen tutorial

https://www.youtube.com/watch?v=A6lDNao00WQ&list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn&index=1

DB Migrations:
```bash
# Create new migration (example)
λ: migrate create -ext sql -dir postgres/migrations create_users

# Apply migrations (up/down)
λ: source .env # to get $POSTGRESQL_URL
λ: migrate -path "postgres/migrations" -database "$POSTGRESQL_URL" up

# Seed Data
λ: psql -d meetup_dev -a -f postgres/seeds.sql # where meetup_dev is the dbname
```
