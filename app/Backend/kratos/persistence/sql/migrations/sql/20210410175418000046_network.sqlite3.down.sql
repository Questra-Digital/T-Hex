INSERT INTO "_selfservice_login_flows_tmp" (id, request_url, issued_at, expires_at, active_method, csrf_token, created_at, updated_at, forced, type, ui) SELECT id, request_url, issued_at, expires_at, active_method, csrf_token, created_at, updated_at, forced, type, ui FROM "selfservice_login_flows";