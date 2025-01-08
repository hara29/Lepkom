set serveroutput on
-- No 1
declare
    cursor cur_emp(id_param number) is 
        select ename, sal 
        from emp 
        where sal < id_param ;
    v_ename emp.ename%type;
    v_sal emp.sal%type;
    e_invalid_empno exception;
begin
    open cur_emp(30);
    fetch cur_emp into v_ename, v_sal;
    if cur_emp%notfound then
        raise e_invalid_empno;
    end if;
    loop
        fetch cur_emp into v_ename, v_sal;
        exit when cur_emp%notfound;
        dbms_output.put_line(v_ename||' memiliki gaji '||v_sal);
    end loop;
    dbms_output.put_line('Total pegawai: '||cur_emp%rowcount||' orang');
    close cur_emp;
exception
    when e_invalid_empno then
    dbms_output.put_line('Data tidak ditemukan');
end;
/
-----------------------------------------------------------------------------
-- No 2
-- create table emp3 as select * from emp;
-- declare
--     cursor cur_emp(id_param number) is
--     select * from emp3 where empno = id_param;
-- begin
--     for peg in cur_emp(7349) loop -- sudah open dan fetch cur_emp ke peg
--         delete from emp3 where empno = peg.empno;
--         dbms_output.put_line('Data '||peg.ename||' berhasil dihapus');
--     end loop; -- otomatis close cur_emp
--     commit;
--     for peg in (select * from emp3) loop
--         dbms_output.put_line(peg.empno||', '||peg.ename||', '||peg.sal);
--     end loop;
-- end;
-- /
-------------------------------------------------------------------------------
-- No 3
-- declare
--     cursor cur_emp(id_param number) is
--     select * from emp3 where empno = id_param;
--     e_invalid_empno exception; --deklarasi exception
-- begin
--     for peg in cur_emp(794) loop -- sudah open dan fetch cur_emp ke peg
--         update emp3 set sal = sal * 2 where empno = peg.empno;
--         dbms_output.put_line('Data '||peg.ename||' berhasil diupdate');
--     end loop; -- otomatis close cur_emp
--     commit;
--     for peg in (select * from emp3) loop
--         dbms_output.put_line(peg.empno||', '||peg.ename||', '||peg.sal);
--     end loop;
-- end;
-- /

