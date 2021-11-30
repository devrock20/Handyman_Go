# Handyman Api's

## **Sample curl requests for USER:**

- **Register User:**
  curl http://localhost:8080/users --include --header "Content-Type: application/json" --request "POST" --data '{"first_name": "Sharon", "last_name": "Williams", "phone_no": 904268124, "email": "sharon12 @gmail.com", "address": "UTN, APT #P", "city": "Rayleigh","state": "North Carolina","password": "sharon13"}'

- **Get all Users:**
  curl http://localhost:8080/users \
    --header "Content-Type: application/json" \
    --request "GET"

- **Update User by id:**
  curl http://localhost:8080/users --include --header "Content-Type: application/json" --request "PUT" --data ‘{“_id":"619c08cf3d567d7611b56729","first_name": "Sharon", "last_name": "Williams", "phone_no": 904268124, "email": "sharon12 @gmail.com", "address": "UTN, APT #P", "city": "Rayleigh","state": "North Carolina","password": "sharon13"}'

- **Authenticate User:**
  curl http://localhost:8080/users/authenticate/sharon12@gmail.com/sharon12 \
    --header "Content-Type: application/json" \
    --request "GET"

- **Delete User by id:**
  curl http://localhost:8080/users/619c08cf3d567d7611b56729 \
    --header "Content-Type: application/json" \
    --request "DELETE"

- **Get User by id:**
  curl http://localhost:8080/users/user/619c2d4eab90fac374b8007c \
    --header "Content-Type: application/json" \
    --request "GET"
    
# Launch Application
--Run **go build** and then double click the exe file that it generates.
--Other command is **go run main.go**

 
  
