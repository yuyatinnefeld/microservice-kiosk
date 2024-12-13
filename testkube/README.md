# Testkube

## Install CLI
- https://docs.testkube.io/articles/install/cli

## Deploy TestKube

- ArgoCD ApplicationSets

```bash
# verify the testkube 
kubectl get all -n testkube

# check if api-server is deployed
kubectl rollout status deployment/testkube-api-server -n testkube

# verify api-server is working
testkube get testworkflow

# check the templates
testkube get testworkflowtemplate -l team=service-mesh
```

## Deploy Smoke tests
```bash
testkube get testworkflow -l team=service-mesh

# run tests
testkube run testworkflow ${TARGET_TEST_NAME} # = tk run tw ${TARGET_TEST_NAME}

# show overview of test results
testkube get testworkflow

# show test result details
testkube get twe ${Execution_ID}
```

### Example
```bash
Context:  (2.1.38)   Namespace: testkube
----------------------------------------
Test Workflow Execution:
Name:                 tw-smoketest-istio-system-kiali
Execution ID:         675c0df2b127fc399f8da516
Execution name:       tw-smoketest-istio-system-kiali-1
Execution namespace:  
Execution number:     1
Requested at:         2024-12-13 20:35:30.775 +0000 UTC
Disabled webhooks:    false
Status:               passed
Queued at:            2024-12-13 20:35:32 +0000 UTC
Started at:           2024-12-13 20:35:32 +0000 UTC
Finished at:          2024-12-13 20:35:38.358 +0000 UTC
Duration:             6.358s
Getting logs for test workflow execution 675c0df2b127fc399f8da516

• Initializing
 Creating state... done
 Initializing state... done
 Configuring init process... done
 Configuring toolkit... skipped
 Configuring shell... done


• (1/2) check-pod-state
 ###################### TEST RESULT ######################
 POD_ID: kiali-6cd6d59486-b4729
 POD_STATE: Running
 Pod is successfully running
 #########################################################


• passed in 385ms

• (2/2) check-http-status-code
 ###################### TEST RESULT ######################
 URL: kiali.istio-system.svc.cluster.local:20001
 HTTP_CODE: 302
 #########################################################

• passed in 5.45s
```

## Deploy Integration tests
```bash
kubectl apply -f testkube/tests/integrationtests
testkube get testworkflow -l test=integrationtests

# run tests
testkube run testworkflow ${TARGET_TEST_NAME} # = tk run tw ${TARGET_TEST_NAME}

# show overview of test results
testkube get testworkflow

# show test result details
testkube get twe ${Execution_ID}
```

### Example
```bash
Context:  (2.1.38)   Namespace: testkube
----------------------------------------
Test Workflow Execution:
Name:                 tw-integrationtest-istio-testapp-httpbin
Execution ID:         675c12cbb127fc399f8da51a
Execution name:       tw-integrationtest-istio-testapp-httpbin-2
Execution namespace:  
Execution number:     2
Requested at:         2024-12-13 20:56:11.372 +0000 UTC
Disabled webhooks:    false
Status:               passed
Queued at:            2024-12-13 20:56:11.372 +0000 UTC
Started at:           2024-12-13 20:56:11.372 +0000 UTC
Finished at:          2024-12-13 20:56:13.584 +0000 UTC
Duration:             2.212s
Getting logs for test workflow execution 675c12cbb127fc399f8da51a

• Initializing
 Creating state... done
 Initializing state... done
 Configuring init process... done
 Configuring toolkit... skipped
 Configuring shell... done


• (1/2) check-http-status-code
 ###################### TEST RESULT ######################
 URL: httpbin.istio-testapp.svc.cluster.local:7775/status/200
 HTTP_CODE: 200
 #########################################################


• passed in 645ms

• (2/2) check-istio-injection
 ###################### TEST RESULT ######################
 POD_ID: httpbin-7bcb56cf64-sxj6x
 CONTAINER_LIST: httpbin istio-proxy
 istio-proxy container is present in the pod httpbin-7bcb56cf64-sxj6x.
 Istio sidecar is injected. 
 #########################################################

• passed in 1.07s
```