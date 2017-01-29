DROP TABLE xmltest;
DROP TABLE author;

create table xmltest(
  id serial,
  data xml NOT null,
  ts timestamp without time zone default (now() at time zone 'utc')
);

create table author (
name varchar
);


--select xpath('/attendee/bio/name/text()',data) from xmltest;
