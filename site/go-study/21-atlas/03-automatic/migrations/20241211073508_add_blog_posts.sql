-- Create "blog_posts" table
CREATE TABLE `blog_posts` (`id` integer NULL, `title` text NOT NULL, `content` text NOT NULL, `user_id` integer NOT NULL, `created_at` timestamp NOT NULL DEFAULT (current_timestamp), `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
