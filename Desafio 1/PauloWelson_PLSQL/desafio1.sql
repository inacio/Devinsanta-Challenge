-- Paulo Welson
-- Solução para Desafio 1
-- Criado e testado em Oracle 11G

  with repetidor as (select level as rnk from dual connect by level <= 200 )
  select
  rownum || case 
  when (mod(rownum,5)+mod(rownum,6)) = 0 then ' Santarem'
  when (mod(rownum,5)) = 0 then ' Santa'
  when (mod(rownum,6)) = 0 then ' rem'
  end as fulll
  from repetidor;
