create function update_updated_at() returns opaque as '
  begin
    new.accesstime := ''now'';
    return new;
  end;
' language 'plpgsql';
