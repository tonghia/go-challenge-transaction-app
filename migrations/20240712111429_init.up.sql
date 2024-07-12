# generate by sqlize

CREATE TABLE `users` (
 `id`         bigint(20) AUTO_INCREMENT PRIMARY KEY,
 `created_at` datetime DEFAULT CURRENT_TIMESTAMP(),
 `updated_at` datetime DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()
);

CREATE TABLE `accounts` (
 `id`         bigint(20) AUTO_INCREMENT PRIMARY KEY,
 `user_id`    bigint(20) COMMENT 'user id',
 `name`       varchar(255) COMMENT 'name',
 `bank`       varchar(10) COMMENT 'bank',
 `created_at` datetime DEFAULT CURRENT_TIMESTAMP(),
 `updated_at` datetime DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()
);
CREATE INDEX `idx_user_id` ON `accounts`(`user_id`);

CREATE TABLE `account_transactions` (
 `id`               bigint(20) AUTO_INCREMENT PRIMARY KEY,
 `user_id`          bigint(20) COMMENT 'user id',
 `account_id`       bigint(20) COMMENT 'account id',
 `bank`             varchar(20) COMMENT 'bank',
 `amount`           bigint(20) COMMENT 'amount',
 `transaction_type` varchar(20) COMMENT 'transaction type',
 `delete_at`        datetime COMMENT 'delete at',
 `created_at`       datetime DEFAULT CURRENT_TIMESTAMP(),
 `updated_at`       datetime DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()
);
CREATE INDEX `idx_user_id_account_id` ON `account_transactions`(`user_id`, `account_id`);