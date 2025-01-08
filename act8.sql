set serveroutput on

declare
    v_ename emp3.ename%type;
begin
    select ename into tmp from emp3 where ename = 'CINDY';
exception
    when no_data_found then 
    dbms_output.put_line('Data tidak ditemukan');
end;
/
