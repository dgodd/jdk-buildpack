# Experimental JDK Buildpack

### Buildpack User Documentation

This buildpack can be used to supply the JDK to other buildpacks.

Example:

```
cf push [NAME] --no-start
cf v3-push [NAME] -p fixtures/simple -b https://github.com/dgodd/jdk-buildpack -b binary_buildpack
```

To confirm jdk is installed

```
cf ssh [NAME]
/tmp/lifecycle/shell
java -version
```

### Manual testing using docker

```
docker run -v $PWD:/bpdir:ro -it cloudfoundry/cflinuxfs2 bash
```

Once in docker, you can run the following to test any changes

```
export CF_STACK=cflinuxfs2 ; rm -rf /app /tmp/cache /tmp/deps ; cp -r /bpdir/fixtures/simple /app ; mkdir -p /tmp/deps/0 ; mkdir -p /tmp/cache; /bpdir/bin/supply /app /tmp/cache /tmp/deps 0
```

## Disclaimer

This buildpack is experimental and not yet intended for production use.
