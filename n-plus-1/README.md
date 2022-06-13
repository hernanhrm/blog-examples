# N+1 query problem

To execute the different programs, you must have a database running:

1. Create a container of postres
   ```bash
   docker run -p 5432:5432 --name postgres12 -e POSTGRES_PASSWORD=secret postgres:12
   ```
2. Connect to the container
   ```bash
   docker exec -it postgres12 psql -U postgres
   ```
4. Create the database
   ```sql
   CREATE DATABASE restaurant;
   ```
5. Execute the sql inside the `sqlmigration` directory
6. Now you can execute the different programs