{
  description = "A very basic flake";
  inputs = { 
    nixpkgs.url = "github:nixos/nixpkgs?ref=23.11"; 
    surreal.url = "github:surrealdb/surrealdb?ref=v1.4.2";
    flake-utils.url = "github:numtide/flake-utils";
    gotempl.url = "github:a-h/templ";
  };
  outputs = { self, nixpkgs, surreal, flake-utils, gotempl}: 
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {inherit system; config.allowUnfree = true;};
        go = pkgs.go;
        gopls = pkgs.gopls;
        helix = pkgs.helix;
        goland = pkgs.jetbrains.goland;
        node = pkgs.nodejs_18;
        watchexec = pkgs.watchexec;
        delve = pkgs.delve;
        starship = pkgs.starship;
        uutils = pkgs.uutils-coreutils;
        git = pkgs.git;
        fish = pkgs.fish;
        zellij = pkgs.zellij;
        pnpm = pkgs.nodePackages.pnpm;
        envsubst = pkgs.envsubst;
        proto = pkgs.protobuf_23;
        protogo = pkgs.protoc-gen-go;
        runit = pkgs.runit;
        surrealdb = surreal.packages.${system}.default;
        templ = gotempl.packages.${system}.templ; 
        minio = pkgs.minio;
        vite = pkgs.vite;
        air = pkgs.air;
        vscodelsps = pkgs.vscode-langservers-extracted;
        miniouser = "minioroot";
        miniopass = "miniopass";
        minioport = "9000";
        minioaddress = "0.0.0.0";
        minioaccesskey = "jDBQiPx5hW4z3mYMXQqx";
        miniosecretkey = "AbZ2Hj8BeCv46UMv17W9fsVcmuCsHCDH0xF61fVs";
        minioconsoleport = "9001";
        surrealdbuser = "surrealroot";
        surrealdbpass = "surrealpass";
        surrealdbport = "9002";
        surrealdbaddress = "0.0.0.0";
        manifoldport = "8080";
        manifoldadminpass = "manifoldpass";
      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            go
            gopls
            helix
            node
            watchexec
            delve
            starship
            uutils
            git
            fish
            zellij
            pnpm
            envsubst
            proto
            protogo
            runit
            surrealdb
            templ
            minio
            vite
            vscodelsps
            air
            goland
          ];
          RUST_BACKTRACE = "full";
          MINIO_ROOT_USER = miniouser;
          MINIO_ROOT_PASSWORD = miniopass;
          MINIO_CONSOLE_PORT = minioconsoleport;
          MINIO_ACCESS_KEY = minioaccesskey;
          MINIO_SECRECT_KEY = miniosecretkey;
          MINIO_PORT = minioport;
          MINIO_ADDRESS = minioaddress;
          SURREAL_DB_USER = surrealdbuser;
          SURREAL_DB_PASSWORD = surrealdbpass;
          SURREAL_DB_PORT = surrealdbport;
          SURREAL_DB_ADDRESS = surrealdbaddress;
          MANIFOLD_PORT = manifoldport;
          MANIFOLD_ADMIN_PASSWORD_HASH = builtins.hashString "md5" manifoldadminpass;
          shellHook = "
              if ! [ -e flake.nix ]; then
                echo \"Please execute nix develop in the directory where your flake.nix is located.\"
                exit 1
              fi
              export HOME=$PWD
              export XDG_HOME=$PWD
              export PNPM_HOME=$HOME/.local/share/pnpm
              export SVDIR=$HOME/.state/services
              export GOPATH=$HOME/.go
              cat $HOME/.state/services/minio/runsubst | envsubst > $HOME/.state/services/minio/run 
              cat $HOME/.state/services/minio/log/runsubst | envsubst > $HOME/.state/services/minio/log/run 
              cat $HOME/.state/services/surrealdb/runsubst | envsubst > $HOME/.state/services/surrealdb/run 
              cat $HOME/.state/services/surrealdb/log/runsubst | envsubst > $HOME/.state/services/surrealdb/log/run 
              cat $HOME/.config/manifold/configsubst.toml | envsubst > $HOME/.config/manifold/config.toml 
              chmod -R 755 .state
              runsvdir .state/services &
              RUNSVDIRPID=$!
              trap 'sv stop surrealdb && sv stop minio && kill -SIGHUP $RUNSVDIRPID' EXIT
              if ! [ -e $GOPATH/bin/nilaway ]; then
                go install go.uber.org/nilaway/cmd/nilaway@latest
              fi
              if ! [ -e $GOPATH/bin/stringer ]; then
                go install golang.org/x/tools/cmd/stringer@latest
              fi
              if ! [ -e $GOPATH/bin/godoc ]; then
                go install golang.org/x/tools/cmd/godoc@latest
              fi
              export PATH=\"$HOME/.go/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/autoprefixer@10.4.19_postcss@8.4.38/node_modules/autoprefixer/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/eslint@8.57.0/node_modules/eslint/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/eslint-config-prettier@9.1.0_eslint@8.57.0/node_modules/eslint-config-prettier/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/prettier@3.2.5/node_modules/prettier/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/svelte-check@3.6.9_postcss@8.4.38_svelte@4.2.14/node_modules/svelte-check/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/svelte-language-server@0.16.7_postcss@8.4.38/node_modules/svelte-language-server/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/typescript@5.4.5/node_modules/typescript/bin:$PATH\"
              export PATH=\"$HOME/Frontend/manifold/node_modules/pnpm/vite@5.2.9/node_modules/vite/bin:$PATH\"
              zellij
          ";
        };
      }  
    );
  }

