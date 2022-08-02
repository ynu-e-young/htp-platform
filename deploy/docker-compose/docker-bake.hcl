variable "REPO" {
  default = "hominsu"
}

variable "AUTHOR_NAME" {
  default = "hominsu"
}

variable "AUTHOR_EMAIL" {
  default = "hominsu@foxmail.com"
}

variable "VERSION" {
  default = ""
}

group "default" {
  targets = [
    "htp-platform-machine-service",
    "htp-platform-machine-robot",
    "htp-platform-capture-service",
    "htp-platform-user-service",
    "htp-platform-htpp-interface",
  ]
}

target "htp-platform-machine-service" {
  context    = "."
  dockerfile = "app/machine/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "machine/service"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/htp-platform-machine-service:${VERSION}" : "",
    "${REPO}/htp-platform-machine-service:latest",
  ]
  platforms = ["linux/amd64"]
}

target "htp-platform-machine-robot" {
  context    = "."
  dockerfile = "app/machine/robot/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "machine/robot"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/htp-platform-machine-robot:${VERSION}" : "",
    "${REPO}/htp-platform-machine-robot:latest",
  ]
  platforms = ["linux/amd64"]
}

target "htp-platform-capture-service" {
  context    = "."
  dockerfile = "app/capture/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "capture/service"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/htp-platform-capture-service:${VERSION}" : "",
    "${REPO}/htp-platform-capture-service:latest",
  ]
  platforms = ["linux/amd64"]
}

target "htp-platform-user-service" {
  context    = "."
  dockerfile = "app/user/service/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "user/service"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/htp-platform-user-service:${VERSION}" : "",
    "${REPO}/htp-platform-user-service:latest",
  ]
  platforms = ["linux/amd64"]
}

target "htp-platform-htpp-interface" {
  context    = "."
  dockerfile = "app/htpp/interface/Dockerfile"
  args       = {
    AUTHOR_NAME       = "${AUTHOR_NAME}"
    AUTHOR_EMAIL      = "${AUTHOR_EMAIL}"
    APP_RELATIVE_PATH = "htpp/interface"
  }
  tags = [
    notequal("", VERSION) ? "${REPO}/htp-platform-htpp-interface:${VERSION}" : "",
    "${REPO}/htp-platform-htpp-interface:latest",
  ]
  platforms = ["linux/amd64"]
}