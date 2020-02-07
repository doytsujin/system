apiVersion: v1
kind: ConfigMap
metadata:
  name: pulsar-gateway
data:
  gatewayImage: bsideup/liiklus:0.9.2
  provisionerImage: {{ gcloud container images describe gcr.io/projectriff/pulsar-provisioner/provisioner:0.5.0 --format="value(image_summary.fully_qualified_digest)" }}
