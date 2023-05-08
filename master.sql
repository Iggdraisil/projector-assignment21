create user 'slave1'@'20.20.20.4' identified by 'password';
GRANT REPLICATION SLAVE ON *.* TO 'slave1'@'20.20.20.4'
create user 'slave2'@'20.20.20.3' identified by 'password';
GRANT REPLICATION SLAVE ON *.* TO 'slave2'@'20.20.20.3'
USE mydb;
FLUSH TABLES WITH READ LOCK;
-- Crate dump
UNLOCK TABLES;
