CREATE TABLE IF NOT EXISTS "user" (
  "id" INT PRIMARY KEY,
  "wcaId" TEXT,
  "name" TEXT NOT NULL,
  "email" TEXT UNIQUE NOT NULL,
  "avatarUrl" TEXT,
  "role" TEXT DEFAULT 'user',
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);
