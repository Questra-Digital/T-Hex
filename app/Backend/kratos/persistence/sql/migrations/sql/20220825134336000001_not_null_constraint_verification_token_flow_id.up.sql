ALTER TABLE identity_verification_tokens
ALTER selfservice_verification_flow_id SET NOT NULL;

DROP INDEX identity_verification_tokens_token_nid_used_idx;
CREATE INDEX identity_verification_tokens_token_nid_used_flow_id_idx ON identity_verification_tokens (nid, token, used, selfservice_verification_flow_id);
