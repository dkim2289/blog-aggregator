# Blog Aggregator

A CLI RSS feed aggregator (multi-user) built using GO and PostgreSQL.

## Prerequisites

- GO
- PostgreSQL

## Installation

```bash
go install github.com/dkim2289/blog-aggregator
```

## Configuration

Create a json config file at your home dir (or ~ dir).

```json
{
  "db_url": "postgres://[username]:@localhost:5432/blogaggregator?sslmode=disable"
}
```

* replace [username] with your PostgreSQL username (likely your local computer name)

## Database

Migration done using Goose.

```bash
goose -dir (path/to/migrations) postgres "postgres://[username]:@localhost:5432/blogaggregator?sslmode=disable" up
```

* replace [username] with your PostgreSQL username (likely your local computer name)

## Commands

- register <name> - registers a new user
- login <name> - log in as an existing user
- users - List all users
- addfeed <name> <url> - adds a new feed
- feeds - List all feeds
- follow <feed> - follow a feed
- following - List all followed feeds
- unfollow <feed> - unfollow a feed
- agg <duration> - aggregate feeds
- browse <post limit> - browse aggregated feeds (default limit: 2)
- reset - reset the database

## Brutal
arg. it was brutal, but it's done.
