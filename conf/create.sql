create table books_dg_tmp
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
  dl_count INTEGER,
  created_at datetime default strftime('%Y-%m-%d %H:%M:%f','now','localtime'),
  updated_at datetime default strftime('%Y-%m-%d %H:%M:%f','now','localtime')
)
;

create index cate_idx
  on books (category)
;

create unique index name_uniq
  on books (name)
;