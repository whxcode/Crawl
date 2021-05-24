create table ars(
  id int primary key auto_increment comment '文章id',
  title varchar(200) not null comment '文章标题',
  icon varchar(200) comment '作者头像',
  author varchar(30) not null comment '作者',
  content longtext not null comment '内容',
  c_time varchar(20) not null comment '创建时间',
  volume varchar(20) not null default '0' comment '访问量',
  fabulous varchar(20) not null default '0' comment '点赞量',
  comment_quantity varchar(20) not null default '0' comment '评论量'
) charset utf8mb4;
123123