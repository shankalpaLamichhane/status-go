CREATE TABLE whisper_keys (
  chat_id TEXT PRIMARY KEY ON CONFLICT IGNORE,
  key BLOB NOT NULL
) WITHOUT ROWID;
