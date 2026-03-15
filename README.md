# gator

A CLI RSS feed aggregator built in Go. Register an account, subscribe to feeds, and browse the latest posts — all from your terminal.

> Built as part of [Boot.dev's "Build a Blog Aggregator" course](https://www.boot.dev).

---

## Prerequisites

You'll need the following installed before running `gator`:

- **[Go](https://go.dev/dl/)** (1.21 or later)
- **[PostgreSQL](https://www.postgresql.org/download/)** (14 or later)

Once Postgres is running, create a database for the aggregator:

```sql
CREATE DATABASE gator;
ALTER USER postgres PASSWORD 'postgres';
```

---

## Installation

Because Go produces statically compiled binaries, you can install `gator` directly with `go install` — no need to keep the source around after that:

```bash
go install github.com/amarquezmazzeo/gator@latest
```

This compiles the binary and places it in your `$GOPATH/bin` directory (typically `~/go/bin`). Make sure that directory is on your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

You can add that line to your `~/.bashrc`, `~/.zshrc`, or equivalent shell config to make it permanent.

---

## Configuration

`gator` reads its config from a JSON file at `~/.gatorconfig.json`. Create it before running the program for the first time:

```bash
touch ~/.gatorconfig.json
```

Populate it with your Postgres connection string:

```json
{
  "db_url": "postgres://YOUR_USER:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `YOUR_USER` and `YOUR_PASSWORD` with your Postgres credentials. The `current_user_name` field is managed automatically by `gator` when you log in.

---

## Usage

The general form of every command is:

```
gator <command> [arguments]
```

### Account management

**Register a new user:**
```bash
gator register <username>
```
Creates an account and automatically logs you in as that user.

**Log in as an existing user:**
```bash
gator login <username>
```
Sets `current_user_name` in your config file.

**List all registered users:**
```bash
gator users
```

---

### Feed management

**Add a new RSS feed:**
```bash
gator addfeed "Blog Name" https://example.com/rss.xml
```
Adds the feed to the database and automatically follows it for your current user.

**List all feeds in the database:**
```bash
gator feeds
```

**Follow an existing feed:**
```bash
gator follow https://example.com/rss.xml
```

**List feeds you're currently following:**
```bash
gator following
```

**Unfollow a feed:**
```bash
gator unfollow https://example.com/rss.xml
```

---

### Aggregation & browsing

**Start the feed aggregator:**
```bash
gator agg 30s
```
Begins polling all feeds on a recurring interval (e.g. `30s`, `1m`, `5m`). Leave this running in a separate terminal — it fetches new posts in the background and writes them to the database.

**Browse the latest posts:**
```bash
gator browse
```

Browse with a custom post limit:
```bash
gator browse 10
```

---

### Utilities

**Reset the database** (deletes all users and data — useful for development):
```bash
gator reset
```

---

## Typical workflow

```bash
# 1. Register and log in
gator register alice

# 2. Subscribe to some feeds
gator addfeed "Go Blog"   https://go.dev/blog/feed.atom
gator addfeed "Hacker News" https://news.ycombinator.com/rss

# 3. Start the aggregator in one terminal
gator agg 1m

# 4. Browse posts in another terminal
gator browse 20
```
