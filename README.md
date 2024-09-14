> :warning: ** Under development. For my personal use while I'm learning. Use it on your own risk.**

## Helm charts development facts

* each Helm chart is stored inside `source` folder in its own folder
* each Helm chart is created out of starter Helm chart stored in `source/starter`
* helm charts are published in this repo, as [releases](https://github.com/curuvija/charts/releases) 
* [here](https://curuvija.github.io/charts/) in GitHub pages you can find instructions how to use it
* documentation is generated using [helm-docs](https://github.com/norwoodj/helm-docs)
* full process of publishing is delegated to [nyx](https://github.com/mooltiverse/nyx) with tasks defined inside `Taskfile.yaml`
* for now publishing process is still manual since I didn't find time to automate it and there are a few issues with the
  setup
* use `task release` definition to understand publishing process (and run tasks manually one by one when publishing)