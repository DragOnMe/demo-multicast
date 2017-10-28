## Does your cluster supports multicasting? ##

### Prerequites ###

go build env:
   assuming 'go install' builds *.go binary into ~/projects/bin

need to use your own docker.io repository account or just use my drlee001/demo-multicast image


### How to do it? ###

Run docker container by running 'docker run ...' or 'kubectl run ...'
and see if multicast is ok by checking out the container logs

### Example(k8s case) ###

- kubectl run demo-multicast --image=drlee001/demo-multicast
- kubectl scale deployment demo-multicast --replicas=3
- kubectl logs -f demo-multicast-xxxxx-yyyy (for each pods)

