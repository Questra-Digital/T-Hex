DELETE FROM identity_verification_tokens
WHERE selfservice_verification_flow_id IS NULL;
