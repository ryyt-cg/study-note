-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "blog_posts" table
DROP TABLE `blog_posts`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
