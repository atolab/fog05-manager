# fog05 Manager
fog05 Manager is a component of the fog05 Kubernetes Control Plane. It is responsible for consuming Control Plane CRDs for the purposes of deploying fog05 FDUs and Entities via Kubernetes clusters.



It leverages `operator-sdk`
https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/#build-and-run-the-operator

## build

```

$ make

```


## run

```
$ make install
$ make run ENABLE_WEBHOOKS=false

```