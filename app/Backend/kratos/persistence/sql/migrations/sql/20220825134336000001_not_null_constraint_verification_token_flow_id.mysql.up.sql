ALTER TABLE identity_verification_tokens
MODIFY selfservice_verification_flow_id CHAR(36) NOT NULL;

DROP INDEX identity_verification_tokens_token_nid_used_idx ON identity_verification_tokens;
CREATE INDEX identity_verification_tokens_token_nid_used_flow_id_idx ON identity_verification_tokens (nid, token, used, selfservice_verification_flow_id);
