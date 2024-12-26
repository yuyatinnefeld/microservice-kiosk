# Grant 'create', 'list', 'read' and 'update' permission to paths prefixed by 'secret/data/kv/' and 'secret/data/kvv2/'

path "kv/*" {
  capabilities = [ "create", "list", "read", "update" ]
}

path "kvv2/*" {
  capabilities = [ "create", "list", "read", "update" ]
}