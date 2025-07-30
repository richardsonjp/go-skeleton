-- Role table
CREATE TABLE "role" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- User table
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL,
    role_id INTEGER NOT NULL REFERENCES "role"(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'active',
    last_login_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Credentials table
CREATE TABLE credential (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    secret VARCHAR(255) NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert default roles
INSERT INTO role
("name", status)
VALUES
    ('master', 'active'),
    ('admin', 'active');

-- Insert default users
INSERT INTO "user"
("name", "password", email, phone, role_id, status)
VALUES
    ('IT', '$2a$12$qKgf5x9ff.TanOy/Dj1rCu9fWRJwX5X/WT2NrwUfqVB7PvsAU1fXS', 'it@kbs.com', '-', 1, 'active'),
    ('admin', '$2a$12$scAwpp3WThykbbGgEDag1OC3VBLhVTwPk8d/spf2HDtUTWzgCQ76C', 'admin@kbs.com', '-', 2, 'active');

-- Frontend paths - keeping page_group for organizational purposes only
CREATE TABLE frontend_path (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, -- e.g., "/dashboard", "/user/:id"
    page_group VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert frontend paths
INSERT INTO frontend_path
("name", page_group)
VALUES
    ('/dashboard', 'dashboard'),
    ('/user', 'user'),
    ('/item-list', 'item');

-- Backend APIs - keeping label for organizational purposes only
CREATE TABLE backend_path (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,   -- e.g., "/dashboard/user", "/dashboard/purchase-request/:code"
    method VARCHAR(20) NOT NULL,  -- GET, POST, PUT, DELETE, etc.
    label VARCHAR(20),            -- e.g., read, create, update - organizational, not for permissions
    page_group VARCHAR(255),      -- organizational, not for permissions
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert backend paths
INSERT INTO backend_path
("name", method, label, page_group)
VALUES
    ('/auth/signin', 'POST', 'login', 'auth'),
    ('/auth/logout', 'POST', 'logout', 'auth'),
    ('/auth/frontend-path', 'GET', 'read', 'auth'),
    ('/auth/home-profile', 'GET', 'read', 'auth'),

    ('/dashboard/user', 'GET', 'read', 'user'),
    ('/dashboard/user', 'POST', 'create', 'user'),
    ('/dashboard/user/roles', 'GET', 'read', 'roles'),
    ('/dashboard/user/roles', 'POST', 'create', 'roles'),
    ('/dashboard/user/roles', 'PUT', 'update', 'roles'),
    ('/dashboard/user/roles/rbac', 'POST', 'read', 'roles');

-- Simplified context_path table - only for direct path associations
CREATE TABLE context_path (
    id SERIAL PRIMARY KEY,
    context_tag VARCHAR(255) NOT NULL, -- Only 'frontend_path' or 'backend_path'
    role_id INTEGER NOT NULL REFERENCES "role"(id) ON DELETE CASCADE,
    path_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (context_tag, role_id, path_id)
);

-- Create index for performance
CREATE INDEX idx_context_path_role_id ON context_path(role_id);
CREATE INDEX idx_context_path_context_tag ON context_path(context_tag);
CREATE INDEX idx_backend_path_group_label ON backend_path(page_group, label);

-- Grant master role access to all frontend paths
INSERT INTO context_path (context_tag, role_id, path_id)
SELECT 'frontend_path', 1, id FROM frontend_path;

-- Grant master role access to all backend paths
INSERT INTO context_path (context_tag, role_id, path_id)
SELECT 'backend_path', 1, id FROM backend_path;

--
-- -- Example: Grant admin role access to all frontend paths
-- INSERT INTO context_path (context_tag, role_id, path_id)
-- SELECT 'frontend_path', 2, id FROM frontend_path;
--
-- -- Example: Grant admin role access to only read operations on certain page groups
-- INSERT INTO context_path (context_tag, role_id, path_id)
-- SELECT 'backend_path', 2, id
-- FROM backend_path
-- WHERE label = 'read'
--   AND page_group IN ('user', 'purchase-request', 'quotation');
--
-- -- Example utility function to find and add all APIs with a specific label
-- CREATE OR REPLACE FUNCTION add_role_permission_by_label(role_id_param INTEGER, label_param VARCHAR)
-- RETURNS void AS $$
-- BEGIN
-- INSERT INTO context_path (context_tag, role_id, path_id)
-- SELECT 'backend_path', role_id_param, id
-- FROM backend_path
-- WHERE label = label_param
--     ON CONFLICT DO NOTHING;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- -- Example utility function to find and add all APIs for a specific page group
-- CREATE OR REPLACE FUNCTION add_role_permission_by_page_group(role_id_param INTEGER, page_group_param VARCHAR)
-- RETURNS void AS $$
-- BEGIN
-- INSERT INTO context_path (context_tag, role_id, path_id)
-- SELECT 'backend_path', role_id_param, id
-- FROM backend_path
-- WHERE page_group = page_group_param
--     ON CONFLICT DO NOTHING;
-- END;
-- $$ LANGUAGE plpgsql;