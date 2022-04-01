{ platform, ... }:

platform.buildGo.package {
  name = "github.com/mutable/base32";

  srcs = [
    ./base32.go
  ];
}
