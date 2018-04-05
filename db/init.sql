# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.20)
# Database: auth
# Generation Time: 2018-04-05 16:18:29 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table apps
# ------------------------------------------------------------

CREATE TABLE `apps` (
  `id` bigint(20) NOT NULL,
  `uuid` char(36) NOT NULL,
  `name` char(64) NOT NULL DEFAULT '',
  `start` timestamp NULL DEFAULT NULL,
  `end` timestamp NULL DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL,
  `enterprise` char(36) NOT NULL DEFAULT '',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  KEY `enterprise of app` (`enterprise`),
  CONSTRAINT `enterprise of app` FOREIGN KEY (`enterprise`) REFERENCES `enterprises` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `apps` WRITE;
/*!40000 ALTER TABLE `apps` DISABLE KEYS */;

INSERT INTO `apps` (`id`, `uuid`, `name`, `start`, `end`, `count`, `enterprise`, `created_time`, `status`)
VALUES
	(1,'0f7b4143-f0ae-11e7-bd86-0242ac120003','example-bot',NULL,NULL,NULL,'bb3e3925-f0ad-11e7-bd86-0242ac120003','2018-04-05 15:21:02',1);

/*!40000 ALTER TABLE `apps` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table enterprises
# ------------------------------------------------------------

CREATE TABLE `enterprises` (
  `id` bigint(20) NOT NULL,
  `uuid` char(36) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `admin_user` char(32) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  KEY `admin of enterprise` (`admin_user`),
  CONSTRAINT `admin of enterprise` FOREIGN KEY (`admin_user`) REFERENCES `users` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `enterprises` WRITE;
/*!40000 ALTER TABLE `enterprises` DISABLE KEYS */;

INSERT INTO `enterprises` (`id`, `uuid`, `name`, `admin_user`, `created_time`)
VALUES
	(1,'bb3e3925-f0ad-11e7-bd86-0242ac120003','emotibot','d3e03673-f0ad-11e7-bd86-0242ac12','2018-04-05 15:21:02');

/*!40000 ALTER TABLE `enterprises` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_column
# ------------------------------------------------------------

CREATE TABLE `user_column` (
  `id` bigint(20) NOT NULL,
  `column` char(32) NOT NULL DEFAULT '',
  `display_name` varchar(64) NOT NULL DEFAULT '',
  `enterprise` char(36) NOT NULL DEFAULT '',
  `note` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `enterprise of custom column` (`enterprise`),
  CONSTRAINT `enterprise of custom column` FOREIGN KEY (`enterprise`) REFERENCES `enterprises` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user_column` WRITE;
/*!40000 ALTER TABLE `user_column` DISABLE KEYS */;

INSERT INTO `user_column` (`id`, `column`, `display_name`, `enterprise`, `note`)
VALUES
	(0,'custom1','自訂屬性1','bb3e3925-f0ad-11e7-bd86-0242ac120003','示範自訂屬性效果');

/*!40000 ALTER TABLE `user_column` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_info
# ------------------------------------------------------------

CREATE TABLE `user_info` (
  `id` bigint(20) NOT NULL,
  `user_id` char(36) NOT NULL DEFAULT '',
  `column_id` bigint(64) NOT NULL,
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user of info` (`user_id`),
  KEY `column of info` (`column_id`),
  CONSTRAINT `column of info` FOREIGN KEY (`column_id`) REFERENCES `user_column` (`id`),
  CONSTRAINT `user of info` FOREIGN KEY (`user_id`) REFERENCES `users` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user_info` WRITE;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;

INSERT INTO `user_info` (`id`, `user_id`, `column_id`, `value`)
VALUES
	(0,'d3e03673-f0ad-11e7-bd86-0242ac12',0,'custom_value1');

/*!40000 ALTER TABLE `user_info` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `uuid` char(36) NOT NULL,
  `display_name` varchar(64) NOT NULL DEFAULT '',
  `email` char(255) NOT NULL DEFAULT '',
  `enterprise` char(36) NOT NULL DEFAULT '',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '2',
  `password` char(32) NOT NULL DEFAULT '',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `uuid`, `display_name`, `email`, `enterprise`, `type`, `password`, `created_time`, `status`)
VALUES
	(1,'d3e03673-f0ad-11e7-bd86-0242ac12','TestUser','test@test.com','bb3e3925-f0ad-11e7-bd86-0242ac120003',1,'5d9c68c6c50ed3d02a2fcf54f63993b6','2018-04-05 15:21:54',1);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
