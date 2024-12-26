# Vault & Secrets Operator

## Deployment
- ArgoCD

## Vault Server Setup
```bash
# Wait until the vault-0 pod is ready
NS_VAULT=vault
kubectl exec -ti vault-0 -n $NS_VAULT -- vault operator init

# Unseal Vault
kubectl config set-context --current --namespace=$NS_VAULT
kubectl port-forward vault-0 8200:8200 -n $NS_VAULT # important

# Login
export VAULT_ADDR='http://127.0.0.1:8200'
vault login

# Upgrade KV engine from v1 to v2
vault secrets enable -version=2 kv

# TEST KV Engine
vault secrets list
vault kv put "kv/test-pullsecret" username=fuj password=san
```

## Create K8S Tech User to Access the External Secret Operator (ESO)
```bash
# Create secret for the secret store to access vault
export VAULT_TECH_USER_NAME=vault-tech-user
export VAULT_TECH_USER_PWD=super-pwd
export TARGET_NS=shop

vault policy write edit-policy vault/edit.hcl
vault auth enable userpass
vault write auth/userpass/users/$VAULT_TECH_USER_NAME password=$VAULT_TECH_USER_PWD policies=edit-policy
vault login -method=userpass username=$VAULT_TECH_USER_NAME password=$VAULT_TECH_USER_PWD
vault token lookup

export VAULT_TECH_USER_TOKEN=hvs.xxxx
export VAULT_TECH_USER_TOKEN_NAME=vault-techuser-token

# Store a secret in the target namespace to exec pull and pull via SecretStore
kubectl create secret -n $TARGET_NS generic $VAULT_TECH_USER_TOKEN_NAME --from-literal=token=$VAULT_TECH_USER_TOKEN
```


## Setup SecretStore, PullSecret, PushSecret

```bash
# Deploy SecretStore
kubectl apply -f vault/manifest/depl-secretStore.yaml
kubectl get ss -n $TARGET_NS

# Deploy pullSecret
kubectl apply -f vault/manifest/depl-pullSecret.yaml
kubectl get externalsecret -n $TARGET_NS
kubectl get secret pulled-secret-from-vault -n $TARGET_NS

# Deploy pushSecret
kubectl apply -f vault/manifest/depl-pushSecret.yaml
kubectl get pushsecret -n $TARGET_NS
vault read kv/data/test-pushsecret
```


