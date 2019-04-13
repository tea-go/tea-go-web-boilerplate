DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `role` enum('admin','member') NOT NULL DEFAULT 'member',
  `status` enum('disabled','enabled') NOT NULL DEFAULT 'enabled',
  `language` varchar(255) NOT NULL DEFAULT 'zh',
  `isDelete` enum('yes','no') NOT NULL DEFAULT 'no',
  `createdAt` datetime NOT NULL,
  `updatedAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB CHARSET=utf8 COMMENT='用户表';

LOCK TABLES `user` WRITE;

INSERT INTO `user` (`id`, `name`, `avatar`, `email`, `role`, `status`, `language`, `isDelete`, `createdAt`, `updatedAt`)
VALUES
	(1, 'baiyu', NULL, 'baiyu@admaster.com.cn', 'admin', 'enabled', 'zh', 'no', '2019-01-01 00:00:00', '2019-01-01 00:00:01'),
	(2, 'jason', NULL, 'jason@admaster.com.cn', 'member', 'enabled', 'zh', 'no', '2019-01-01 00:00:00', '2019-01-01 00:00:01');
