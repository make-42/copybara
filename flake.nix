{
  description = "Copybara - A Wayland clipboard automation tool";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };

      goDeps = with pkgs; [
        go
        wl-clipboard
      ];
    in {
      devShell = pkgs.mkShell {
        buildInputs = goDeps;
      };

      packages.default = pkgs.buildGoModule {
        pname = "copybara";
        version = "0.1.0";

        src = ./.;

        vendorHash = null;

        nativeBuildInputs = goDeps;
        buildInputs = goDeps;
      };
    });
}
