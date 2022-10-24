# Scheduler Service
Ensures that the orders are processed by the Kitchen service and that the packages are being delivered by the Drone delivery service.
 
## Tech Stack
1. Golang
2. Firestore

## Environment
Environment variables are listed in .env/dev file

## Installation guide
Food service binary can be fetched from https://github.com/dietdoctor/be-test/releases

Start Firestore emulator 

```
gcloud beta emulators firestore start --host-port=0.0.0.0
```

For Firestore UI docker can be used 

```
docker run   --rm   -p=8080:8080   -p=4000:4000   -p=9099:9099   -p=8085:8085   -p=5001:5001   -p=9199:9199   --env "GCP_PROJECT=dietdoctor"   --env "ENABLE_UI=true"   spine3/firebase-emulator
```

The service uses TLS to secure transport of the API endpoints, a cert/key pair must be provided, example is showen in the [Hits](#hits) part

```
export FIRESTORE_EMULATOR_HOST=0.0.0.0:8080

./food-linux-amd64 server --store-namespace=foo --tls-cert=server.crt --tls-key=server.key --gcp-project-id=dietdoctor --debug --controller-interval=10s
```
## Spent time
~10 hours 

## Hits
Creation of certificate example steps:

-  create config.conf file in the directory binary is, add these

```
[req]
distinguished_name = req_distinguished_name
req_extensions = req_ext
prompt = no

[req_distinguished_name]
C   = IN
ST  = Karnataka
L   = Bengaluru
O   = GoLinuxCloud
OU  = R&D
CN  = ban21.example.com

[req_ext]
subjectAltName = @alt_names

[alt_names]
IP.1 = 127.0.0.1
IP.2 = 0.0.0.0
```

If it will not working use it

```
[req]
prompt = no
req_extensions = ext
distinguished_name = req_distinguished_named[ ext ]
subjectAltName = IP:0.0.0.0, IP:127.0.0.1[ req_distinguished_named ]
C=AU
ST=NSW
L=Sydney
O=Arkady Balaba
OU=Engineering
CN=localhost
```

- Create a certificate:

```
openssl genrsa -out server.key
openssl req -new -sha256 -key server.key -out server.csr -config config.conf
openssl x509 -req -days 3650 -in server.csr -out server.crt -signkey server.key -extensions ext -extfile config.conf
```


