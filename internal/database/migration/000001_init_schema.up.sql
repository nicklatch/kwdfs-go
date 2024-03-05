create type user_position as enum ('fleet_rep', 'service_manager', 'service_director', 'dealer_principal', 'store_manager', 'other');

alter type user_position owner to postgres;

create type fleet_status as enum ('Regional', 'National');

alter type fleet_status owner to postgres;

create table dealers
(
    id                  uuid default gen_random_uuid() not null
        primary key,
    dealer              varchar(100)                   not null,
    street_address      varchar(255)                   not null,
    state               char(2)                        not null,
    zip                 varchar(20),
    general_phone       varchar(20)                    not null,
    director_of_service varchar(100),
    logo_url            varchar(255)
);

comment on column dealers.street_address is 'Full address';

comment on column dealers.state is 'Two letter state abbr';

comment on column dealers.general_phone is 'XXX-XXX-XXXX format';

comment on column dealers.director_of_service is 'First and last name';

comment on column dealers.logo_url is 'url for s3';

alter table dealers
    owner to postgres;

grant delete, insert, references, select, trigger, truncate, update on dealers to anon;

grant delete, insert, references, select, trigger, truncate, update on dealers to authenticated;

grant delete, insert, references, select, trigger, truncate, update on dealers to service_role;

create table locations
(
    id                      char(4)     not null
        primary key,
    dealer                  uuid        not null
        references dealers,
    branch_name             text        not null,
    state                   char(2)     not null,
    city_county             varchar(50) not null,
    fleet_support_rep       text        not null,
    fleet_support_rep_email text,
    general_phone           text
);

comment on column locations.state is 'Two letter state abbr';

comment on column locations.fleet_support_rep is 'The name of the fleet support rep';

comment on column locations.fleet_support_rep_email is 'Email of the fleet support representitive';

comment on column locations.general_phone is 'main phone number for a location';

alter table locations
    owner to postgres;

grant delete, insert, references, select, trigger, truncate, update on locations to anon;

grant delete, insert, references, select, trigger, truncate, update on locations to authenticated;

grant delete, insert, references, select, trigger, truncate, update on locations to service_role;

create table customers
(
    id                      uuid         default gen_random_uuid() not null
        primary key,
    dealer                  uuid                                   not null
        references dealers,
    name                    text                                   not null,
    state                   char(2)                                not null,
    pfleet_acct_id          text,
    status                  fleet_status default 'national'::fleet_status,
    truck_qty               smallint                               not null,
    fleet_support_rep       text                                   not null,
    field_service_rep       text,
    field_service_rep_email text
);

alter table customers
    owner to postgres;

grant delete, insert, references, select, trigger, truncate, update on customers to anon;

grant delete, insert, references, select, trigger, truncate, update on customers to authenticated;

grant delete, insert, references, select, trigger, truncate, update on customers to service_role;

