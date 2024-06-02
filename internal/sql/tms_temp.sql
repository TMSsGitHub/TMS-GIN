create table tb_user
(
    id         bigint unsigned primary key comment 'id',
    phone      char(11) unique not null default '' comment '电话号码',
    pwd        varchar(20)     not null default '' comment '密码',
    sex        char(1)         not null default '男' comment '性别 男/女',
    email      varchar(50)     not null default '' comment '邮箱',
    avatar_url varchar(200)    not null default '' comment '头像url',
    created_at  bigint unsigned not null default 0 comment '创建时间',
    updated_at  bigint unsigned not null default 0 comment '更新时间',
    deleted_at  bigint unsigned not null default 0 comment '删除时间'
) comment '用户信息表';
