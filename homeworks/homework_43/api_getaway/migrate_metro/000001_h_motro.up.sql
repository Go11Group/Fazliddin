CREATE TABLE station(
    station_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name varchar,
    deleted_at bigint default 0
);

CREATE TABLE terminal(
    terminal_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    station_id uuid REFERENCES station(station_id),
    deleted_at bigint default 0
);
