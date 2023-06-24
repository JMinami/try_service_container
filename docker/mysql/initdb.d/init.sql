-- データ読み込みはデフォルトで無効化されているので有効にする
SET GLOBAL local_infile=on;
DROP DATABASE IF EXISTS features;
CREATE DATABASE test DEFAULT CHARACTER SET utf8mb4;

USE test;
CREATE TABLE `test_tables` (
    `id` VARCHAR(40) NOT NULL,
    `name` VARCHAR(40) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;