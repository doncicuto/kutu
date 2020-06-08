# Kutu - Kubernetes Useful Tools Updater

When using K8s in Muultipla we have several tools that we use to deploy and build our applications. Most of these tools are binaries that change frequently so we want to check if these binaries have a new version, download the new version and install it easily.

Kutu will check if a new release is available and update it for the following K8s binaries:

- kubectl
- skaffold
- minikube
- kustomize

**_NOTE:_**  Kutu currently only checks and update K8s binaries that exists already in your system, in a future release Kutu may also install them.

In order to download and update new versions just run:

`kutu update`

Kutu will replace your old binary with the new one.

**_WARNING:_**  If you use brew to install K8s binaries, be warned that Kutu will replace your old binary file without brew's intervention so it could break your homebrew's workflow.

**_WARNING:_**  If your K8s binaries are stored in paths requiring higher privileges to replace the binaries, please run kutu with sudo.

If you only want to check if new versions are available for your binaries use the check command.

`kutu check`

If you want to check or update specific binaries use the -b / --binaries option followed by a list of comma-separated binaries:

`kutu update -b minikube`

`kutu check -b skaffold,kubectl`

You can also create a config file if you want to set the binaries that Kutu will check or update:

```yaml
---
binaries: skaffold,kubectl
```

Your YAML configuration can be stored in $HOME/.kutu.yaml or you can pass the file path using -c / --config:

`kutu check --config /tmp/kutu.yaml`

## Acknowledgements

Kutu has been built thanks to these awesome packages:

- [Cobra](https://github.com/spf13/cobra)
- [Color](https://github.com/gookit/color)
- [Grab](https://github.com/cavaliercoder/grab)
