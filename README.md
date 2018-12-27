# AppsKube

client-go implementation of `appskube`

At first you have to install **appskube** using `go install` command.

Available commands :

* `appskube` - Welcome from AppsKube
* `appskube --help` 
* `appskube <command> --help`
* `appskube create` - default name="appscode", replicas=1
* `appskube create --name=<name> --replicas=<number-of-replicas>`
* `appskube scale` - default name="appscode", replicas=5
* `appskube scale --name=<name> --replicas=<<desired-number>`
* `appskube expose` - default name="appscode"
* `appskube expose --name=<name>`
* `appskube ingress` - default host-name="software.farm", name="appscode"
* `appskube ingress --host=<host-name> --name=<name-of-service>`
* `appskube delete` - default name="appscode"
* `appskube delete --name=<name>`


