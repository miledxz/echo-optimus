# echo-optimus
Simple architecture echo server


<img width="703" alt="Screenshot 2024-06-28 at 15 55 10" src="https://github.com/zedGGs/echo-optimus/assets/153705375/5abf1356-dd87-43a3-b97c-5760d07d2c12">



First thing is to make sure you do: go mod tidy <br>
Define your environment .env
Second thing is running database with : docker compose up -d <br>
Run server: go run main.go 

available endpoints: 

GET localhost:5008/health <br>
POST localhost:5008/service/user <br>
GET localhost:5008/service/users
