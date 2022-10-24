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

The service uses TLS to secure transport of the API endpoints, a cert/key pair must be provided

```
export FIRESTORE_EMULATOR_HOST=0.0.0.0:8080

./food-linux-amd64 server --store-namespace=foo --tls-cert=server.crt --tls-key=server.key --gcp-project-id=dietdoctor --debug --controller-interval=10s
```

## Spended time 
~ 15 hours