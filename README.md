# mongo-go-example
example repo to help setup with mongo-go drivers

First build the docker image

```bash
docker build -t mongo-image .
```

Then run the docker container

```bash
docker run -d -p 27017:27017 --name mongo mongo-image
```

Once the mongo container is up and running, you can either view the contents using mongo shell

```bash
mongo --host localhost --port 27017
```