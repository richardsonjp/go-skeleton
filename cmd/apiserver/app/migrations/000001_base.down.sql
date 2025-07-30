-- Drop all tables and objects in reverse order of creation

-- First drop constraints that reference other tables
ALTER TABLE "user" DROP CONSTRAINT IF EXISTS user_role_id_fkey;
ALTER TABLE credential DROP CONSTRAINT IF EXISTS credential_user_id_fkey;
ALTER TABLE context_path DROP CONSTRAINT IF EXISTS context_path_role_id_fkey;

-- Then drop indexes
DROP INDEX IF EXISTS idx_context_path_role_id;
DROP INDEX IF EXISTS idx_context_path_context_tag;
DROP INDEX IF EXISTS idx_backend_path_group_label;

-- Then drop tables
DROP TABLE IF EXISTS context_path;
DROP TABLE IF EXISTS backend_path;
DROP TABLE IF EXISTS frontend_path;
DROP TABLE IF EXISTS credential;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "role";

-- Note: If you had any custom functions, they should be dropped here too
-- DROP FUNCTION IF EXISTS add_role_permission_by_label;
-- DROP FUNCTION IF EXISTS add_role_permission_by_page_group;