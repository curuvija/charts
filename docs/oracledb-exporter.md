
To be able to test oracledb-exporter you'll have to run oracledb instance in Kubernetes cluster.

Package Helm chart (oracledb https://github.com/oracle/docker-images/blob/main/OracleDatabase/SingleInstance/helm-charts/oracle-db/README.md.):

```bash
git clone https://github.com/oracle/docker-images.git
cd ../docker-images/OracleDatabase/SingleInstance/helm-charts/
helm package oracle-db/ --destination ~/Desktop/museum-charts/
```

Command to run local helm charts repo:

```bash
chartmuseum --debug --port=8080 \
  --storage="local" \
  --storage-local-rootdir="/home/milos/Desktop/museum-charts"
```

Now update repo:

```bash
helm repo add chartmuseum http://localhost:8080
```

And you should see:

```bash
milos@milos-desktop:~/Downloads$ helm search repo chartmuseum
NAME                 	CHART VERSION	APP VERSION	DESCRIPTION           
chartmuseum/oracle-db	1.0.0        	           	Oracle Database Server
```

Choose docker image from repo https://container-registry.oracle.com/.

Pull docker image (the image is large and the download speed is slow so Kubernetes will fail to download it in timely manner):

```bash
docker pull container-registry.oracle.com/database/express:21.3.0-xe
```

Now you should be able to install it (docker-desktop example):

```bash
helm upgrade --install oracledb chartmuseum/oracle-db --version 1.0.0 --set persistence.storageClass=hostpath --set image=container-registry.oracle.com/database/express:21.3.0-xe
```

Check the output:

```bash
 
Release "oracledb" has been upgraded. Happy Helming!
NAME: oracledb
LAST DEPLOYED: Tue Jun 25 06:01:19 2024
NAMESPACE: default
STATUS: deployed
REVISION: 2
NOTES:
#
# Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
# Licensed under the Universal Permissive License v 1.0 as shown at http://oss.oracle.com/licenses/upl.
#

# ===========================================================================
# == Add below entries to your tnsnames.ora to access this database server ==  
# ====================== from external host =================================  
ORCLCDB=(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=<ip-address>)(PORT=<port>))
    (CONNECT_DATA=(SERVER=DEDICATED)(SERVICE_NAME=<ORACLE_SID>)))
ORCLPDB1=(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=<ip-address>)(PORT=<port>))
    (CONNECT_DATA=(SERVER=DEDICATED)(SERVICE_NAME=<ORACLE_PDB>)))
#                                                                              
#ip-address : IP address of any of the Kubernetes nodes
#port       : Service Port that is mapped to the port 1521 of the container.
#

Application details
====================
IP and port can be found using the following:
export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services oracledb-oracle-db)
export NODE_XDB_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[1].nodePort}" services oracledb-oracle-db)
export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
echo listener at $NODE_IP:$NODE_PORT
echo XDB at $NODE_IP:$NODE_XDB_PORT

Oracle Databases SID, PDB name can be figured out by :

ORACLE_SID=$(kubectl get  -o jsonpath="{.spec.template.spec.containers[?(.name == 'oracle-db')].env[?(.name == 'ORACLE_SID')].value }" deploy oracledb-oracle-db)
ORACLE_PDB=$(kubectl get  -o jsonpath="{.spec.template.spec.containers[?(.name == 'oracle-db')].env[?(.name == 'ORACLE_PDB')].value }" deploy oracledb-oracle-db)

```

Docs for chart museum are here https://github.com/helm/chartmuseum.