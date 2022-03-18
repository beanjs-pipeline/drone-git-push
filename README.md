# drone-git-push

## usage

```yaml
kind: pipeline
type: docker
name: default

steps:
  - name: push <xxx>
    image: beanjs/drone-git-push
    settings:
      local_dir: repository
      branch: master
      commit_message: 'commit from ci'
```
