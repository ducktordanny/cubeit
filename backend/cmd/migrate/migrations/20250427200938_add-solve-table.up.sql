CREATE TABLE IF NOT EXISTS "solve" (
  "id" SERIAL PRIMARY KEY,
  "userId" INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  "categoryId" INT NOT NULL REFERENCES "puzzleCategory"(id),
  "penaltyId" INT REFERENCES penalty(id),
  "result" INT NOT NULL,
  "isPR" BOOLEAN NOT NULL DEFAULT FALSE,
  "scramble" TEXT NOT NULL,
  "solution" TEXT,
  "note" TEXT,
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);
