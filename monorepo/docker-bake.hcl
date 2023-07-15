variable "GITHUB_SHA" {
  default = ""
}

target "api" {
    dockerfile = "./apps/api/Dockerfile"
    context = "./"
    tags = [
      "api"
    ]
}
