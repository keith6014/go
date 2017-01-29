#!/usr/bin/env bash
db="books"
dropdb $db
createdb $db 
psql $db < books.sql  
psql $db < func.sql
