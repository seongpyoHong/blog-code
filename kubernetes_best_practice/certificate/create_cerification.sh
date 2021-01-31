#!/bin/bash

csr_name="client-csr"
name="user"
csr="client-csr.csr"

cat <<EOF | kubectl create -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
    name: ${csr_name}
spec:
    groups:
    - system:authenticated
    request: $(cat ${csr} | base64 | tr -d '\n')
    signerName: kubernetes.io/kube-apiserver-client
    usages:
    - digital signature
    - key encipherment
    - client auth
EOF

echo "Approving Sigining Request"
kubectl certificate approve ${csr_name}

echo "Downloading Certificate."
kubectl get csr ${csr_name} -o jsonpath='{.status.certificate}' \
        | base 64 --decode > $(basename ${csr} .csr).crt

echo "Cleaning Up"
kubectl delete csr ${csr_name}

echo "Add Configuration for User to kubeconfig"
kubectl config set-credentials ${name} --client-key=${PWD}/$(basename ${csr} .csr)-key.pem -client-certificate=${PWD}/$(basename ${csr} .csr).crt --embed-certs=true

