CREATE TABLE IF NOT EXISTS communities (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('community'),
    name VARCHAR(255) NOT NULL UNIQUE,
    censo_id integer NOT NULL
);

INSERT INTO communities (name, censo_id) 
VALUES
  ('Conjunto Esperança', 1),
  ('Vila Do João', 2),
  ('Conjunto Pinheiros', 3),
  ('Vila Dos Pinheiros', 4),
  ('Novo Pinheiros (Salsa E Merengue)', 5),
  ('Conjunto Bento Ribeiro Dantas (Fogo Cruzado)', 6),
  ('Morro Do Timbau', 7),
  ('Baixa Do Sapateiro', 8),
  ('Nova Maré', 9),
  ('Parque Maré', 10),
  ('Nova Holanda', 11),
  ('Parque Rubens Vaz', 12),
  ('Parque União', 13),
  ('Parque Roquete Pinto', 14),
  ('Praia De Ramos', 15),
  ('Marcílio Dias', 16)
ON CONFLICT (name) DO NOTHING;
