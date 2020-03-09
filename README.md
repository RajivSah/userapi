# userapi

### Create table
```
aws dynamodb create-table --table-name User --attribute-definitions AttributeName=Uid,AttributeType=S --key-schema AttributeName=Uid,KeyType=HASH --endpoint-url http://localhost:8000 --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```

#### Create user
1. Get all users
Request
```
GET /users
```
Response
```
200:
[
  {
    "Age": 67,
    "Gender": "Male",
    "Uid": "b8a475ca-71e9-4f99-be8a-55cd27e683de",
    "Name": "Ramesh Snow"
  },
  {
    "Age": 20,
    "Gender": "Female",
    "Uid": "81327b65-d7c4-42d8-9166-45ae9474b994",
    "Name": "Arya Kumari Bahadur"
  }
]
```

2. Add a User
Request
```
POST /users
Body:
  {
    "Age": 67,
    "Gender": "Male",
    "Name": "Ramesh Snow"
  }
```
Response
```
201
```

3. Update a user
Request
```
PUT /users
Body:
  {
    "Age": 20,
    "Gender": "Female",
    "Uid": "81327b65-d7c4-42d8-9166-45ae9474b994",
    "Name": "Arya Kumari Bahadur"
  }
```
Response
```
200
```


