create table device(
	clientId	varchar(15),
    name		varchar(20),
    primary key(clientId)
);

create table deviceInfo(
	clientId	varchar(15),
    info		varchar(50),
    value		integer,
    alert		integer,
    lng			double,
    lat			double,
    timestamp	long,
    foreign key(clientId) references device(clientId)
);

create table user(
	username	varchar(20)		unique,
    password	varchar(20)		not null,
    email		varchar(50)		unique,
    primary key(username)
);

set SQL_SAFE_UPDATES = 0;

update device set name = "空调" where clientId = "device0001";

drop table device;
select * from device;
select * from deviceInfo;
delete from device;

drop table user;
select * from user;

insert into user values("11","11","11",null);
delete from user;

select count(*) from deviceInfo where alert = 1;
select * from deviceInfo  where clientId = "device0001" order by clientId desc limit 10;

