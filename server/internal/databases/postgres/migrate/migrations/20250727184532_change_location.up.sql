BEGIN;

ALTER TABLE locations
RENAME COLUMN neighborhood TO community_id;

ALTER TABLE locations
ADD CONSTRAINT fk_locations_community
FOREIGN KEY (community_id) REFERENCES communities(id);

COMMIT;

