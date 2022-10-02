# Belajar Go Lang RESTful API

## Project: Aplikasi CRUD Sederhana
- Digunakan untuk belajar RESTful API, bukan untuk membuat aplikasi
- Memilik API Authentication untuk semua _endpoint_

**Data**
```go
type Category struct {
  Id    int32   `json:"id"`
  Name  string  `json:"name"`
}

```