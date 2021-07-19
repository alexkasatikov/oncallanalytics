CREATE TABLE "alerts" (
  "id" serial PRIMARY KEY,
  "fingerprint" varchar NOT NULL,
  "status" varchar NOT NULL,
  "startsat" timestamptz NOT NULL,
  "endsat" timestamptz,
  UNIQUE (fingerprint)
);
CREATE TABLE "labels" (
  "id" serial PRIMARY KEY,
  "key" varchar NOT NULL,
  "value" varchar NOT NULL,
  UNIQUE (key, value)
);
CREATE TABLE "alerts_labels" (
  "id" serial PRIMARY KEY,
  "alert_id" int NOT NULL,
  "label_id" int NOT NULL,
  UNIQUE (alert_id, label_id),
  FOREIGN KEY (alert_id) REFERENCES alerts(id),
  FOREIGN KEY (label_id) REFERENCES labels(id)
);