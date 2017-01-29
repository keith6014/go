DROP TABLE xmltest;
DROP TABLE author;

create table xmltest(
  id serial,
  data xml NOT null
);

create table author (
name varchar
);


--select xpath('/attendee/bio/name/text()',data) from xmltest;
