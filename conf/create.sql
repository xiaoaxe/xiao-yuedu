drop table books;

create table books
(
  id INTEGER
    constraint id_pk
    primary key,
  name varchar not null,
  format varchar not null,
  content varchar,
  tags varchar,
  category INTEGER,
  pan_url varchar,
  view_count INTEGER,
  download_count INTEGER,
  created_at INTEGER,
  updated_at INTEGER
)
;

create index cate_idx on books (category);

create unique index name_uniq on books (name);

######## category ########
drop table categories;

create table categories
(
  id INTEGER
    constraint id_pk
    primary key,
  category INTEGER not null,
  major varchar,
  sub varchar,
  created_at INTEGER,
  updated_at INTEGER
)
;

create unique index cate_uniq on categories (category);

