-- Create table for car data
create table car
(
    id       uuid primary key,
    make     varchar,
    model    varchar,
    year     int,
    price    int,
    owner    varchar
);

SELECT
    tablename,
    indexname,
    indexdef
FROM
    pg_indexes
WHERE
    schemaname = 'public'
ORDER BY
    tablename,
    indexname;

-- Create single column B-tree index
create index car_id_btree_idx on car using btree (id);

-- Create single column Hash index
create index car_id_hash_idx on car using hash (id);


-- Create multi-column B-tree index
create index car_make_model_btree_idx on car using btree (make, model);

create index car_make_model_ID_btree_idx on car using btree (id, make, model);



-- Drop
drop index car_id_btree_idx;


explain (analyse)
select * from car where make = 'BMW' and model = 'David' and id = '38695b16-dc2b-4150-88f6-2bc240ca41e2';



 Planning Time: 40.45 ms
 Execution Time: 31.17 ms