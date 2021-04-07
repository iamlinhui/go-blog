CREATE DATABASE db_blog CHARACTER SET utf8mb4;

DROP TABLE IF EXISTS db_blog.t_article;
CREATE TABLE db_blog.t_article
(
    id             bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    content        longtext        NOT NULL COMMENT '文章内容',
    title          text            NOT NULL COMMENT '文章标题',
    excerpt        text            NOT NULL COMMENT '文章简介',
    url            varchar(255)    NOT NULL DEFAULT '' COMMENT '文章访问固定链接',
    user_id        bigint          NOT NULL DEFAULT 0 COMMENT '作者ID',
    create_time    datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    modify_time    datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    status         tinyint         NOT NULL DEFAULT 0 COMMENT '文章状态 0:未发布,1:已发布',
    comment_status tinyint         NOT NULL DEFAULT 0 COMMENT '评论开关 0:不开启,1:开启',
    PRIMARY KEY (id),
    UNIQUE INDEX unique_idx_url (url),
    FULLTEXT INDEX full_idx_article (title, excerpt, content) WITH PARSER ngram
) ENGINE = InnoDB
  CHARSET = utf8mb4 COMMENT '文章表';

DROP TABLE IF EXISTS db_blog.t_category;
CREATE TABLE db_blog.t_category
(
    id     bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    name   varchar(255)    NOT NULL DEFAULT '' COMMENT '分类目录中文名(展示)',
    prefix varchar(255)    NOT NULL DEFAULT '' COMMENT '分类目录访问前缀(url)',
    sort   tinyint         NOT NULL DEFAULT 0  COMMENT '标签排序',
    PRIMARY KEY (id),
    UNIQUE INDEX unique_idx_url (prefix)
) ENGINE = InnoDB
  CHARSET = utf8mb4 COMMENT '分类目录表';

DROP TABLE IF EXISTS db_blog.t_dictionary;
CREATE TABLE db_blog.t_dictionary
(
    id        bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    name      varchar(255)    NOT NULL DEFAULT '' COMMENT '配置key',
    value     varchar(255)    NOT NULL DEFAULT '' COMMENT '配置value',
    introduce varchar(255)    NOT NULL DEFAULT '' COMMENT '配置说明',
    PRIMARY KEY (id),
    UNIQUE INDEX unique_idx_name (name)
) ENGINE = InnoDB
  CHARSET = utf8mb4 COMMENT '配置字典表';

DROP TABLE IF EXISTS db_blog.t_user;
CREATE TABLE db_blog.t_user
(
    id          bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    login       varchar(255)    NOT NULL DEFAULT '' COMMENT '登录的名称 不可修改 唯一',
    email       varchar(255)    NOT NULL DEFAULT '' COMMENT '邮箱 唯一 可修改',
    salt        varchar(255)    NOT NULL DEFAULT '' COMMENT '盐',
    password    varchar(255)    NOT NULL DEFAULT '' COMMENT '密码',
    name        varchar(255)    NOT NULL DEFAULT '' COMMENT '中文名',
    create_time datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    login_time  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近登录日期',
    status      tinyint         NOT NULL DEFAULT 0 COMMENT '状态 0冻结 1可用',
    role        tinyint         NOT NULL DEFAULT 0 COMMENT '0管理,1会员',
    PRIMARY KEY (id),
    UNIQUE INDEX unique_idx_login (login),
    UNIQUE INDEX unique_idx_email (email)
) ENGINE = InnoDB
  CHARSET = utf8mb4 COMMENT '用户表';

DROP TABLE IF EXISTS db_blog.t_relationships;
CREATE TABLE db_blog.t_relationships
(
    article_id  bigint UNSIGNED NOT NULL DEFAULT 0,
    category_id bigint UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (article_id, category_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '文章-目录关联表';
