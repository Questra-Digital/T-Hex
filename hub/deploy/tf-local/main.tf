terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.20"
    }
  }
}

provider "docker" {}

resource "docker_network" "selenium" {
  name = "selenium-network"
}

resource "docker_network" "web" {
  name = "web-network"
}

resource "docker_volume" "pgdata" {
  name = "pgdata"
}

resource "docker_container" "db" {
  name  = "db"
  image = "postgres:14-alpine"
  restart = "always"
  volumes {
    volume_name    = docker_volume.pgdata.name
    container_path = "/var/lib/postgresql/data"
  }
  env = [
    "POSTGRES_USER=${var.postgres_user}",
    "POSTGRES_PASSWORD=${var.postgres_password}",
    "POSTGRES_DB=${var.postgres_db}",
    "PGUSER=${var.postgres_user}"
  ]
  networks_advanced {
    name = docker_network.web.name
  }
  healthcheck {
    test         = ["CMD-SHELL", "pg_isready -U thex -d thex -h localhost"]
    interval     = "10s"
    timeout      = "5s"
    retries      = 5
    start_period = "80s"
  }
}

resource "docker_container" "am" {
  name  = "am"
  image = "jawadc/t-hex:am-1.0.0"
  //env   = ["AM_DEMO=1"]
  must_run = false
  restart  = "no"
  networks_advanced {
    name = docker_network.web.name
  }
  depends_on = [docker_container.db]
}

resource "docker_container" "nginx" {
  name  = "proxse-proxy"
  image = "nginx:alpine"
  restart = "always"
  ports {
    internal = 4445
    external = 4445
  }
  volumes {
    host_path      = abspath("${path.module}/nginx.conf")
    container_path = "/etc/nginx/nginx.conf"
  }
  networks_advanced {
    name = docker_network.web.name
  }
  depends_on = [docker_container.proxse]
}

resource "docker_container" "proxse" {
  count  = var.scale_count
  name   = "proxse-${count.index}"
  image = "jawadc/t-hex:proxse-1.0.0"
  restart = "always"
  networks_advanced {
    name = docker_network.selenium.name
  }
  networks_advanced {
    name = docker_network.web.name
  }
  depends_on = [docker_container.db, docker_container.am]
}

resource "docker_container" "fe" {
  name  = "fe"
  image = "jawadc/t-hex:fe-1.0.0"
  restart = "always"
  ports {
    internal = 8080
    external = 80
  }
  networks_advanced {
    name = docker_network.web.name
  }
  depends_on = [docker_container.db, docker_container.am]
}

resource "docker_container" "selenium_hub" {
  name  = "selenium-hub"
  image = "selenium/hub:4.27.0"
  restart = "always"
  /*ports {
    internal = 4444
    external = 4444
  }*/
  networks_advanced {
    name = docker_network.selenium.name
  }
}

resource "docker_container" "chrome_node" {
  count = var.scale_count
  name  = "node-chrome-${count.index}"
  image = "selenium/node-chrome:131.0"
  restart = "always"
  env = [
    "SE_EVENT_BUS_HOST=selenium-hub",
    "SE_EVENT_BUS_PUBLISH_PORT=4442",
    "SE_EVENT_BUS_SUBSCRIBE_PORT=4443",
    "HUB_HOST=selenium-hub",
    "HUB_PORT=4444"
  ]
  networks_advanced {
    name = docker_network.selenium.name
  }
  volumes {
    host_path      = "/dev/shm"
    container_path = "/dev/shm"
  }
  depends_on = [docker_container.selenium_hub]
}
