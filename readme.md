# Hexagonal Architecture

### Executing
- start golang package
  - access app container
  - > docker exec -it container_name bash
  - initialize golang packave
  - > go mod init package_name

### Mocks
- generate mocks from interfaces with mockgen
  - access app container
  - > docker exec -it container_name bash
  - > mockgen -destination=application/mocks/application.go -source=application/interfaces/product_interfaces.go interfaces

### Database
- create sqlite3 database
  - access app container
  - > docker exec -it container_name bash
  - at this point, sqlite3 client should be ready to go, as it is on dockerfile
  - create simple database file so sqlite can use it
  - > touch sqlite.db
  - connect with sqlite client and execute the sql examples (available in rawSql.txt file)
  - > sqlite3 sqlite.db