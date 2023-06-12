-- @author chatgpt3.5

-- migration_history
CREATE TABLE migration_history (
  version TEXT NOT NULL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now'))
);

-- user
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  username TEXT NOT NULL UNIQUE,
  role TEXT NOT NULL CHECK (role IN ('HOST', 'ADMIN', 'USER')) DEFAULT 'USER',
  password_hash TEXT NOT NULL
);

-- recipient
CREATE TABLE recipient (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  email TEXT NOT NULL,
  reachable TEXT NOT NULL
);

-- sender
CREATE TABLE sender (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  host TEXT NOT NULL,
  port INTEGER NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL
);

-- template
CREATE TABLE template (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  subject TEXT NOT NULL,
  body TEXT NOT NULL
);

-- resource
CREATE TABLE resource (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  filename TEXT NOT NULL DEFAULT '',
  blob BLOB DEFAULT NULL,
  external_link TEXT NOT NULL DEFAULT '',
  type TEXT NOT NULL DEFAULT '',
  size INTEGER NOT NULL DEFAULT 0,
  internal_path TEXT NOT NULL DEFAULT ''
);

-- template_resource
CREATE TABLE template_resource (
  template_id INTEGER NOT NULL,
  resource_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  UNIQUE(template_id, resource_id)
);
