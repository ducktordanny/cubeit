CREATE TABLE IF NOT EXISTS "user" (
  "id" INT PRIMARY KEY,
  "wcaId" TEXT,
  "name" TEXT NOT NULL,
  "email" TEXT UNIQUE NOT NULL,
  "gender" TEXT NOT NULL,
  "bio" TEXT DEFAULT '',
  "countryISO" TEXT NOT NULL,
  "avatarURL" TEXT,
  "role" TEXT DEFAULT 'user',
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);
