INSERT INTO "_identity_recovery_tokens_tmp" (id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, issued_at, nid) SELECT id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, issued_at, nid FROM "identity_recovery_tokens";