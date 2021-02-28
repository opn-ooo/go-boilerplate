create function update_updated_at() returns trigger as '
  begin
    new.accesstime := ''now'';
    return new;
  end;
' language 'plpgsql';
