{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule rec {
  pname = "chx";
  version = "1.0.0";

  src = ./.;

  vendorSha256 = null;

  buildInputs = with pkgs; [];

  buildFlags = [ "-tags=production" ];

  meta = with pkgs.lib; {
    description = "CLI for threat research right in the terminal";
    homepage = "https://github.com/EpykLab/chx";
    license = licenses.mit;
    maintainers = with maintainers; [ DavidHoenisch ];
  };
}
