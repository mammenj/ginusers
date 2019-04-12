Sample Microservices in Go language

======================

Command line instructions

Git global setup
git config --global user.name "John Sakthimangalam Mammen01"
git config --global user.email "john.mammen01@infosys.com"

Create a new repository
git clone http://infygit.ad.infosys.com/John.Mammen01/gousers.git
cd gousers
touch README.md
git add README.md
git commit -m "add README"
git push -u origin master

Existing folder
cd existing_folder
git init
git remote add origin http://infygit.ad.infosys.com/John.Mammen01/gousers.git
git add .
git commit -m "Initial commit"
git push -u origin master

Existing Git repository
cd existing_repo
git remote rename origin old-origin
git remote add origin http://infygit.ad.infosys.com/John.Mammen01/gousers.git
git push -u origin --all
git push -u origin --tags
====

keytool -export -keystore keystore.p12 -alias interns2019 -file interns2019.cer
openssl x509 -inform der -in interns2019.cer -out interns2019.pem
