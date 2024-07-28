-- +goose Up
-- +goose StatementBegin
create table packages
(
   id_package SERIAL PRIMARY KEY check ( id_package > 0),
   package_cost  SMALLINT check (package_cost > 0)   not null,
   package_name text not null,
   lower_mass smallint check (lower_mass >= 0) not null,
   upper_mass smallint check (upper_mass >  0) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table packages;
-- +goose StatementEnd
