A example Microservices in Go language

### Supported features ###
- Create user
- Modify user
- Delete user
- List user and all users

- Authentication based on JWT tokens
- Authorization based on JWT claims and Casbin declartive authorization
- SSL enabled endpoint

- Uses Gin gonic web framework https://github.com/gin-gonic/gin


### For ssl ####

keytool -export -keystore keystore.p12 -alias interns2019 -file interns2019.cer
openssl x509 -inform der -in interns2019.cer -out interns2019.pem

### postgresql ####
systemctl status postgresql - check status of postgres
systemctl status postgresql - start postgres

$ sudo -u postgres psql

`User table
==========
   Column   |           Type           | Collation | Nullable |              Default              
------------+--------------------------+-----------+----------+-----------------------------------
 id         | integer                  |           | not null | nextval('users_id_seq'::regclass)
 uid        | uuid                     |           |          | 
 username   | text                     |           |          | 
 password   | text                     |           |          | 
 message    | text                     |           |          | 
 created_at | timestamp with time zone |           |          | 
 updated_at | timestamp with time zone |           |          | 

Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
===========

select * from users

id | uid | username | password | message | created_at | updated_at 
 
INSERT INTO users (username, password) VALUES
  ('admin', '12345'),
  ('user1', '12345');

delete from users where id = 2`

==========================