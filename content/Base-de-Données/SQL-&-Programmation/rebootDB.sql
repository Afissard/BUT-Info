CREATE TYPE TableNamesType AS TABLE OF VARCHAR2(200);

CREATE OR REPLACE PROCEDURE ORDER_66 IS
DECLARE
  Jedis TableNamesType;
BEGIN
    DBMS_OUTPUT.PUT_LINE('Yes my lord');
    -- Jedis := TableNamesType('EMPLOYE', 'SERVICE', 'PROJET', 'TRAVAIL', 'CONCERNE');
    -- FOR jedi IN (SELECT * FROM Jedis)
    -- LOOP
    --     EXECUTE IMMEDIATE 'DROP TABLE'||jedi||'CASCADE CONSTRAINTS PURGE;'
    END LOOP; 
    EXECUTE IMMEDIATE 'DROP TABLE concerne CASCADE CONSTRAINTS PURGE'; 
    EXECUTE IMMEDIATE 'DROP TABLE travail CASCADE CONSTRAINTS PURGE';
    EXECUTE IMMEDIATE 'DROP TABLE projet CASCADE CONSTRAINTS PURGE';
    EXECUTE IMMEDIATE 'DROP TABLE employe CASCADE CONSTRAINTS PURGE';
    EXECUTE IMMEDIATE 'DROP TABLE service CASCADE CONSTRAINTS PURGE';
    DBMS_OUTPUT.PUT_LINE('This is done my lord');
    COMMIT;
END;

-- EXECUTE ORDER_66;
SELECT table_name FROM user_tables;
