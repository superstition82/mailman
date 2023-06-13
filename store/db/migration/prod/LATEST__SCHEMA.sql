-- Create user table
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  username TEXT NOT NULL UNIQUE,
  role TEXT NOT NULL CHECK (role IN ('HOST', 'ADMIN', 'USER')) DEFAULT 'USER',
  password_hash TEXT NOT NULL
);

-- Create recipient table
CREATE TABLE recipient (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  email TEXT NOT NULL,
  reachable TEXT NOT NULL
);

-- Create sender table
CREATE TABLE sender (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  host TEXT NOT NULL,
  port INTEGER NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL
);

-- Create template table
CREATE TABLE template (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  subject TEXT NOT NULL,
  body TEXT NOT NULL
);

-- Create resource table
CREATE TABLE resource (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  filename TEXT NOT NULL DEFAULT '',
  blob BLOB,
  external_link TEXT NOT NULL DEFAULT '',
  type TEXT NOT NULL DEFAULT '',
  size INTEGER NOT NULL DEFAULT 0,
  internal_path TEXT NOT NULL DEFAULT '',
  CONSTRAINT fk_resource_template
    FOREIGN KEY (id) REFERENCES template_resource(resource_id)
    ON DELETE CASCADE
);

-- Create template_resource table
CREATE TABLE template_resource (
  template_id INTEGER NOT NULL,
  resource_id INTEGER NOT NULL,
  created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  PRIMARY KEY (template_id, resource_id),
  CONSTRAINT fk_template_resource_template
    FOREIGN KEY (template_id) REFERENCES template(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_template_resource_resource
    FOREIGN KEY (resource_id) REFERENCES resource(id)
    ON DELETE CASCADE
);
