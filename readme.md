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