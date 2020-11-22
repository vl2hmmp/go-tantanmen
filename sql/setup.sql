CREATE DATABASE IF NOT EXISTS `tantanmen` DEFAULT CHARACTER SET utf8mb4;
CREATE USER IF NOT EXISTS `tantanmen`@`localhost` IDENTIFIED WITH mysql_native_password BY 'tantanmen';
GRANT ALL ON `tantanmen`.* TO `tantanmen`@`localhost`;