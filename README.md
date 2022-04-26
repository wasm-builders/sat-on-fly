# sat-on-fly

## Create and build the runnable

```bash
subo create runnable sha256 --lang go
cd sha256
subo build .

```

## Build and push the Docker image

```bash 
docker_handle="k33g" # use your own handle
docker_pwd="your password"
docker login -u ${docker_handle} -p ${docker_pwd}

cd sha256
app_name="sha256-demo"
tag="0.0.0"
docker build -t ${app_name} . 
docker tag ${app_name} ${docker_handle}/${app_name}:${tag}
docker image ls
docker push ${docker_handle}/${app_name}:${tag}
```

## Fly.io configuration

You need to create a `fly.toml` file at the root of the Runnable project `sha256` (it's the application configuration file: https://fly.io/docs/reference/configuration/)

```toml
kill_signal = "SIGINT"
kill_timeout = 5
processes = []

[experimental]
  allowed_public_ports = []
  auto_rollback = true

[[services]]
  http_checks = []
  internal_port = 8080
  processes = ["app"]
  protocol = "tcp"
  script_checks = []

  [services.concurrency]
    hard_limit = 25
    soft_limit = 20
    type = "connections"

  [[services.ports]]
    handlers = ["http"]
    port = 80

  [[services.ports]]
    handlers = ["tls", "http"]
    port = 443

  [[services.tcp_checks]]
    grace_period = "1s"
    interval = "15s"
    restart_limit = 0
    timeout = "2s"
```

## Deployment

```bash
# Create the application, only at the first deplyment
flyctl apps create ${app_name} --json

# Deploy
# Don't forget to set FLY_ACCESS_TOKEN
flyctl deploy \
  --app ${app_name} \
  --image ${docker_handle}/${app_name}:${tag} \
  --env SAT_HTTP_PORT=8080 \
  --verbose --json

```

Wait for a moment, and then call:

```bash
http --form POST https://${app_name}.fly.dev --raw "Bob Morane"
curl -d 'Bob Morane' https://${app_name}.fly.dev

curl -d '👋 Hello World 🌍' https://sha256-demo.fly.dev
```