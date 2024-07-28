-- +goose Up
-- +goose StatementBegin
create table order_data
(
    id_order        BIGINT check (id_order > 0)    not null unique,
    id_user         BIGINT check (id_user > 0)     not null,
    id_package SMALLINT check (id_package > 0)  not null,
    foreign key (id_package) references packages (id_package),
    delivered_date  timestamp with time zone not null,
    recieved_date   timestamp with time zone not null,
    dead_line       timestamp with time zone not null,
    refund_date     timestamp with time zone not null,
    item_mass SMALLINT check(item_mass > 0) not null

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table order_data;
-- +goose StatementEnd
