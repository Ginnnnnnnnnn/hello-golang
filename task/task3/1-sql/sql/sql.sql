-- 学生表
CREATE TABLE `students` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '学生ID',
    `name` varchar(255) NOT NULL COMMENT '学生姓名',
    `age` int(11) NOT NULL COMMENT '学生年龄',
    `grade` varchar(255) NOT NULL COMMENT '学生年级',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生表';

-- 账户表
CREATE TABLE `accounts` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `balance` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账户表';

-- 转账记录表
CREATE TABLE `transactions` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `from_account_id` bigint NOT NULL COMMENT '转出账户ID',
    `to_account_id` bigint NOT NULL COMMENT '转入账户ID',
    `amount` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '转账金额',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='转账记录表';