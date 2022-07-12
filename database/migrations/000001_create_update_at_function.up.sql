create function update_updated_at() returns trigger as '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' language 'plpgsql';

create extension if not exists "uuid-ossp";
