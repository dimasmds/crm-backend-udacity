# CRM Backend Udacity

### Customers Feature

#### Get All Customers

- URL
    - `/customers`
- Method
    - `GET`
- Response
    ```json
    [
      {
        "id": "081e798e-23ca-4db2-83bb-ddc6a1ac5727",
        "name": "Dimas Saputra",
        "role": "Software Engineer",
        "email": "dimas@dicoding.com",
        "phone": 621324,
        "contacted": false
      },
      {
        "id": "b19e6af6-725b-4574-8aad-ab79490466aa",
        "name": "Nina Pratiwi",
        "role": "Artist",
        "email": "nina@dicoding.com",
        "phone": 621321,
        "contacted": true
      }
    ]
   ```

#### Get Single Customer

- URL
    - `/customers/{id}`
- Method
    - `GET`
- Response
    ```json
    {
      "id": "081e798e-23ca-4db2-83bb-ddc6a1ac5727",
      "name": "Dimas Saputra",
      "role": "Software Engineer",
      "email": "dimas@dicoding.com",
      "phone": 621324,
      "contacted": false
   }
    ```

#### Add New Customer

- URL
    - `/customers`
- Method
    - `POST`

#### Update Existing Customer

- URL
    - `/customers/{id}`
- Method
    - `PUT`

#### Delete Existing Customer

- URL
    - `/customers/{id}`
- Method
    - `DELETE`
