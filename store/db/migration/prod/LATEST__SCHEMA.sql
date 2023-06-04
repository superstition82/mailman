-- migration_history
CREATE TABLE migration_history (
  version TEXT NOT NULL PRIMARY KEY,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now'))
);

CREATE TABLE recepient (
  id      INTEGER PRIMARY KEY AUTOINCREMENT,
  email   TEXT NOT NULL,
  status  TEXT NOT NULL
);

CREATE TABLE sender (
  id       INTEGER PRIMARY KEY AUTOINCREMENT,
  email    TEXT NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE template (
  id          INTEGER PRIMARY KEY AUTOINCREMENT,
  subject     TEXT NOT NULL,
  body        TEXT NOT NULL
);
