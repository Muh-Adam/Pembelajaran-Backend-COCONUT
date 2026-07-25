Clean Architechture
- Penulisan kode agar rapi
- agar mudah dibaca
- agar mudah dikembangkan

# Structure
- Browser (Frontend)
- API (Application Programming Interface)
- Handler layer
- Service layer
- Repository layer
- Database

## Handler layer (HTTP layer)
- Menerima request
- Memanggil service
- Mengirim response

## Service layer
- Menerima request dari handler
- Melakukan validasi
- Melakukan bisnis logic
- Memanggil repository
- Mengirim response ke handler

## Repository layer
- Menerima request dari service
- Melakukan validasi
- Melakukan bisnis logic
- Memanggil database
- Mengirim response ke service