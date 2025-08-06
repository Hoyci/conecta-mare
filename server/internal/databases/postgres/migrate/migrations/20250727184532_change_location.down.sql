BEGIN;

ALTER TABLE locations
DROP CONSTRAINT IF EXISTS fk_locations_community;

ALTER TABLE locations
RENAME COLUMN community_id TO neighborhood;

COMMIT;
