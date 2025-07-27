-- 员工表
CREATE TABLE `employees` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '员工ID',
    `name` varchar(255) NOT NULL COMMENT '员工名称',
    `department` varchar(255) NOT NULL COMMENT '员工部门',
    `salary` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '员工薪水',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='员工表';

-- 书籍表
CREATE TABLE `books` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '书籍ID',
    `title` varchar(255) NOT NULL COMMENT '书籍标题',
    `author` varchar(255) NOT NULL COMMENT '书籍作者',
    `price` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '书籍价格',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='书籍表';