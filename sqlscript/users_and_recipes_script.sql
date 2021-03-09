CREATE DATABASE IF NOT EXISTS user_recipes;
USE user_recipes;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userName` varchar(255) NOT NULL,
  `loginName` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` varchar(255) NOT NULL,
  `userRole` varchar(255) NOT NULL,
  `pictureURL` varchar(255) DEFAULT 'https://thumbs.dreamstime.com/z/unknown-person-22968622.jpg',
  `shortDescription` varchar(255) NOT NULL,
  `activated` tinyint(1) NOT NULL,
  `created` datetime NOT NULL,
  `modified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `userName` (`userName`),
  CONSTRAINT `users_chk_2` CHECK (((`gender` = _utf8mb4'male') or (`gender` = _utf8mb4'female'))),
  CONSTRAINT `users_chk_3` CHECK (((`userRole` = _utf8mb4'user') or (`userRole` = _utf8mb4'admin'))),
  CONSTRAINT `users_chk_4` CHECK ((length(`password`) >= 8))
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `recipes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NOT NULL,
  `recipe_title` varchar(80) NOT NULL,
  `short_description` text NOT NULL,
  `time_to_cook` int NOT NULL,
  `products` text NOT NULL,
  `recipe_pictureURL` varchar(255) DEFAULT 'https://i.insider.com/5d0bc2a0e3ecba03841d82d2?width=960&format=jpeg',
  `detailed_description` text NOT NULL,
  `keywords` text NOT NULL,
  `created` datetime NOT NULL,
  `modified` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `recipes_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci