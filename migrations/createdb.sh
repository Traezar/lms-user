psql -U postgres -c "CREATE USER leave_management WITH PASSWORD 'leave_management_password';"
psql -U postgres  -c "CREATE DATABASE leave_management;"
psql -U postgres  -c "GRANT ALL PRIVILEGES ON DATABASE leave_management TO leave_management;"