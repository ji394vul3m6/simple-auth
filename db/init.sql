-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- 主機: db
-- 產生時間： 2018 年 01 月 03 日 18:13
-- 伺服器版本: 5.7.19
-- PHP 版本： 7.0.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

--
-- 資料庫： `auth`
--
CREATE DATABASE IF NOT EXISTS `auth` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `auth`;

-- --------------------------------------------------------

--
-- 資料表結構 `app`
--

CREATE TABLE `apps` (
  `id` bigint(20) NOT NULL,
  `uuid` char(36) NOT NULL,
  `name` char(64) NOT NULL DEFAULT '',
  `start` timestamp NULL DEFAULT NULL,
  `end` timestamp NULL DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL,
  `enterprise` char(36) NOT NULL DEFAULT '',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 資料表新增前先清除舊資料 `app`
--

TRUNCATE TABLE `apps`;
--
-- 資料表的匯出資料 `app`
--

INSERT INTO `apps` (`id`, `uuid`, `name`, `start`, `end`, `count`, `enterprise`, `created_time`, `status`) VALUES
(2, '0f7b4143-f0ae-11e7-bd86-0242ac120003', 'example-bot', NULL, NULL, NULL, 'bb3e3925-f0ad-11e7-bd86-0242ac120003', CURRENT_TIMESTAMP, 1);

-- --------------------------------------------------------

--
-- 資料表結構 `enterprise`
--

CREATE TABLE `enterprises` (
  `id` bigint(20) NOT NULL,
  `uuid` char(36) NOT NULL,
  `name` char(64) NOT NULL DEFAULT '',
  `admin_user` char(32) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 資料表新增前先清除舊資料 `enterprise`
--

TRUNCATE TABLE `enterprises`;
--
-- 資料表的匯出資料 `enterprise`
--

INSERT INTO `enterprises` (`id`, `uuid`, `name`, `admin_user`, `created_time`) VALUES
(1, 'bb3e3925-f0ad-11e7-bd86-0242ac120003', 'emotibot', 'd3e03673-f0ad-11e7-bd86-0242ac12', CURRENT_TIMESTAMP);

-- --------------------------------------------------------

--
-- 資料表結構 `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `uuid` char(36) NOT NULL,
  `display_name` char(64) DEFAULT NULL,
  `email` char(255) NOT NULL DEFAULT '',
  `enterprise` char(36) NOT NULL DEFAULT '',
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT '2',
  `password` char(32) NOT NULL DEFAULT '',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 資料表新增前先清除舊資料 `users`
--

TRUNCATE TABLE `users`;
--
-- 資料表的匯出資料 `users`
--

INSERT INTO `users` (`id`, `uuid`, `display_name`, `email`, `enterprise`, `type`, `password`, `created_time`, `status`) VALUES
(1, 'd3e03673-f0ad-11e7-bd86-0242ac120003', 'emotibot', NULL, 'emotibot@test.com', 'bb3e3925-f0ad-11e7-bd86-0242ac120003', 1, '1a165ac8a11f729ecfcea4cfb58adb74', CURRENT_TIMESTAMP, 1);

--
-- 已匯出資料表的索引
--

--
-- 資料表索引 `app`
--
ALTER TABLE `apps`
  ADD PRIMARY KEY (`id`),
  ADD KEY `enterprise` (`enterprise`);

--
-- 資料表索引 `enterprise`
--
ALTER TABLE `enterprises`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uuid` (`uuid`);

--
-- 資料表索引 `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `enterprise` (`enterprise`);

--
-- 在匯出的資料表使用 AUTO_INCREMENT
--

--
-- 使用資料表 AUTO_INCREMENT `app`
--
ALTER TABLE `apps`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用資料表 AUTO_INCREMENT `enterprise`
--
ALTER TABLE `enterprises`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用資料表 AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 已匯出資料表的限制(Constraint)
--

--
-- 資料表的 Constraints `app`
--
ALTER TABLE `apps`
  ADD CONSTRAINT `app_ibfk_1` FOREIGN KEY (`enterprise`) REFERENCES `enterprises` (`uuid`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- 資料表的 Constraints `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`enterprise`) REFERENCES `enterprises` (`uuid`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;
