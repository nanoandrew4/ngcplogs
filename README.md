# ngcplogs

[Dockerhub](https://hub.docker.com/repository/docker/nanoandrew4/ngcplogs/general)

A modified version of the standard `gcplogs` docker logging driver. The standard `gcplogs` driver does not process the
output from the containers, which means JSON logs result in a log like this:

```json
{
  "insertId": "1x3kge4f3if919",
  "jsonPayload": {
    "instance": {
      "id": "5968118946548037465",
      "zone": "us-east1-b",
      "name": "gcp-vm"
    },
    "message": "{\"app\":\"sample-app\",\"level\":\"error\",\"msg\":\"Error authenticating user\",\"time\":\"2024-03-14T13:27:42Z\"}",
    "container": {
      "imageId": "sha256:360b4beb988621daaa87572c42af11142a14ecc7c3a5b4cdf221d5d97b19acdc",
      "id": "32c7e2402ec77cf94121a52c9d284939038d0dff9952696a17b2fa6da74f47bb0",
      "imageName": "nanoandrew4/some-image",
      "name": "/sample-app",
      "created": "2023-09-07T23:01:22.718629265Z"
    }
  },
  "resource": {
    "type": "gce_instance",
    "labels": {
      "instance_id": "12345",
      "zone": "us-east1-b",
      "project_id": "someproject"
    }
  },
  "timestamp": "2024-03-14T13:27:42.033851150Z",
  "logName": "projects/someproject/logs/gcplogs-docker-driver",
  "receiveTimestamp": "2024-03-14T13:27:42.670937226Z"
}
```

This driver behaves similarly to the `gcplogs` one, but if it detects a JSON log, it unmarshals it and sends the unmarshalled map as the payload, resulting in a log like this (with the default options):
```json
{
  "insertId": "yero77f8j919i9",
  "jsonPayload": {
    "msg": "Updating at 2024-03-15 11:21:38.56773049 +0000 UTC m=+901.837891394\n",
    "app": "sample-app",
    "container": {
      "created": "2024-03-15T11:06:28.730214829Z",
      "id": "af2c42f7720c8dec812abc5d7cee903aaadf1cd04d87488f3ab1657b92977bc6",
      "name": "/sample-app",
      "imageId": "sha256:360b4beb988121df8587572c42af15102a14ecc7c3a5d4cdf221d5d67b29acdc",
      "imageName": "nanoandrew4/sample-app"
    },
    "instance": {
      "zone": "us-east1-b",
      "name": "gcp-vm",
      "id": "8319386972505717539"
    },
    "time": "2024-03-15T11:21:40Z"
  },
  "resource": {
    "type": "gce_instance",
    "labels": {
      "zone": "us-east1-b",
      "instance_id": "8390836155502727539",
      "project_id": "someproject"
    }
  },
  "timestamp": "2024-03-15T11:21:40.080322442Z",
  "severity": "INFO",
  "logName": "projects/someproject/logs/ngcplogs-docker-driver",
  "receiveTimestamp": "2024-03-15T11:21:44.099223634Z"
}
```

Non JSON logs will not be processed, and will be sent to GCP as they were received, without being manipulated.

### Installation

```shell
docker plugin install nanoandrew4/ngcplogs:v1.0.0 --grant-all-permissions
```

In your `daemon.json` file, change the `log-driver` to `nanoandrew4/ngcplogs:v1.0.0`

Finally, restart the daemon and docker services:

```shell
sudo systemctl daemon-reload && sudo systemctl restart docker
```

### Upgrading
First stop all containers using the plugin. Once they are all stopped, run the following commands, supposing there was a 
version v0.1.0, and we want to upgrade to use v1.0.0

```shell
docker plugin disable nanoandrew4/ngcplogs:v0.1.0
docker plugin rm nanoandrew4/ngcplogs:v0.1.0
docker plugin install nanoandrew4/ngcplogs:v1.0.0 --grant-all-permissions
```


In your `daemon.json` file, change the `log-driver` to `nanoandrew4/ngcplogs:v1.0.0`

Finally, restart the daemon and docker services:

```shell
sudo systemctl daemon-reload && sudo systemctl restart docker
```

Start all your containers again, and they should be using the new version of the plugin

### Configuration

The following [log-opts](https://docs.docker.com/config/containers/logging/configure/#configure-the-default-logging-driver) are available for configuration:

| log-opt              | default | description                                                                                                                                                                                                                                                               |
|----------------------|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| extract-json-message | true    | Enables unmarshalling JSON messages and sending the jsonPayload as the unmarshalled map. Kind of the whole point of this plugin, but you can disable it so it behaves just like the `gcplogs` plugin if you wish                                                          |
| local-logging        | false   | Enables logging to a local file, so logs can be viewed with the `docker logs` command. If false, the command will show no output                                                                                                                                          |
| extract-severity     | true    | Extracts the severity from JSON logs to set them for the log that will be sent to GCP. It will be removed from the jsonPayload section, since it is set at the root level. Currently the supported severity field names to extract are the following: `severity`, `level` |
| exclude-timestamp    | false   | Excludes timestamp fields from the final jsonPayload, since docker sends its own nanosecond precision timestamp for each log. Currently it can remove fields with the following names: `timestamp`, `time`, `ts`                                                          |

### Building locally

If you want to build the plugin yourself, use the makefile with the following command
```shell
make all
```

See the Makefile for other build targets