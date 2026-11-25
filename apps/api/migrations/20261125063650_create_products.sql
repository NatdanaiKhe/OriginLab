-- +goose Up
CREATE TABLE products (
  id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name           TEXT NOT NULL CHECK (btrim(name) <> ''),
  description    TEXT,
  price   INTEGER NOT NULL CHECK (price >= 0), 
  stock_on_hand  INTEGER NOT NULL DEFAULT 0 CHECK (stock_on_hand >= 0),
  image_url      TEXT,
  is_active      BOOLEAN NOT NULL DEFAULT TRUE,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_products_name_active ON products (name, is_active);


-- +goose Down
DROP TABLE products;
