Sample Microservices in Go language

### For ssl ####

keytool -export -keystore keystore.p12 -alias interns2019 -file interns2019.cer
openssl x509 -inform der -in interns2019.cer -out interns2019.pem
