-- https://stackoverflow.com/questions/12343984/insert-trigger-to-update-another-table-using-postgresql

CREATE OR REPLACE FUNCTION function_copy() RETURNS TRIGGER AS
$BODY$
BEGIN
   INSERT INTO author (name) values ((select unnest(xpath('/attendee/bio/name/text()',data)) from xmltest));

           RETURN new;
END;
$BODY$
language plpgsql;

CREATE TRIGGER trig_copy
     AFTER INSERT ON xmltest
     FOR EACH ROW
     EXECUTE PROCEDURE function_copy();


INSERT INTO xmltest (data, id) VALUES ('                                                                                                                                                       
<attendee>                                                                                                                                                                                     
 <bio>                                                                                                                                                                                         
 <name>John Doe</name>                                                                                                                                                                         
 <birthYear>1986</birthYear>                                                                                                                                                                   
 </bio>                                                                                                                                                                                        
 <languages>                                                                                                                                                                                   
 <lang level="5">php</lang>                                                                                                                                                                    
 <lang level="4">python</lang>                                                                                                                                                                 
 <lang level="2">java</lang>                                                                                                                                                                   
 </languages>                                                                                                                                                                                  
</attendee>', 1);    

