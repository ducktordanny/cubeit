CREATE TABLE IF NOT EXISTS "averageSolve" (
  "id" SERIAL PRIMARY KEY,
  "averageId" INT NOT NULL REFERENCES average(id) ON DELETE CASCADE,
  "solveId" INT NOT NULL REFERENCES solve(id) ON DELETE CASCADE
);
