数据库操作
show databases;
create database dbname;
use dbname;
select database();//查看当前所在db
drop database dbname;

表操作
show tables;
create table  tbname(字段 类型)
desc tbname;
show create table;//建表信息
alter table tbname add/modify/change/drop/rename to...
drop table tbname;

查表
select * from tbname;

增删改
insert into tablename() values() where;
update tablename set 字段 = value where;
delete from tablename where;