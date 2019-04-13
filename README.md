Sample Microservices in Go language

### For ssl ####

keytool -export -keystore keystore.p12 -alias interns2019 -file interns2019.cer
openssl x509 -inform der -in interns2019.cer -out interns2019.pem

### postgresql ####
systemctl status postgresql - check status of postgres
systemctl status postgresql - start postgres

$ sudo -u postgres psql

select * from users

id | uid | username | password | message | created_at | updated_at 
 
INSERT INTO users (username, password) VALUES
  ('admin', '12345'),
  ('user1', '12345');

delete from users where id = 2