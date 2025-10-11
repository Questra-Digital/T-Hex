// GitHub API service for token and repository validation
import { APIResponse } from '@/types/API';
import { httpUtils, HttpRequestOptions } from '@/services/httpUtils';

interface RepositoryInfo {
  permissions: {
    admin: boolean;
  };
}

const baseUrl = 'https://api.github.com';

/*
  Default options for GitHub API requests
*/
const defaultOptions = (token: string) => ({
  headers: {
    'Authorization': `token ${token}`,
    'Accept': 'application/vnd.github.v3+json',
  },
}) as HttpRequestOptions;

function evaluateStatus(statusCode: number | undefined): APIResponse {
  if (!statusCode) {
    return {
      success: false,
      message: 'Unexpected error occurred',
      error: 'Unexpected error occurred',
    };
  }
  switch (statusCode) {
    case 401:
      return {
        success: false,
        message: 'Invalid or expired token',
        error: 'Invalid or expired token',
      };
    case 403:
      return {
        success: false,
        message: 'Token does not have access to repository',
        error: 'Token does not have access to repository',
      };
    case 404:
      return {
        success: false,
        message: 'Repository not found or token does not have access',
        error: 'Repository not found or token does not have access',
      };

    default:
      return {
        success: false,
        message: 'Unexpected error occurred',
        error: 'Unexpected error occurred',
      };
  }
}
/**
 * Validate GitHub token and check repository access
 */
export const validateTokenAndRepo = async (
  token: string,
  repoPath: string
): Promise<APIResponse> => {

  // First, validate the token by getting user info

  const userResponse = await httpUtils.get(`${baseUrl}/user`, defaultOptions(token));

  if (!userResponse.success) {
    return evaluateStatus(userResponse.statusCode);
  }

  // Parse repository path (format: owner/repo)
  const [owner, repo] = repoPath.split('/');
  if (!owner || !repo) {
    return {
      success: false,
      message: 'Invalid repository path format. Use: owner/repository',
      error: 'Invalid repository path format. Use: owner/repository',
    };
  }

  // Check repository access and permissions
  const repoResponse = await httpUtils.get<APIResponse<RepositoryInfo>>(`${baseUrl}/repos/${owner}/${repo}`, defaultOptions(token));


  if (!repoResponse.success) {
    return evaluateStatus(repoResponse.statusCode);
  }

  const repository: RepositoryInfo = repoResponse.data as RepositoryInfo;

  if (!repository.permissions.admin) {
    return evaluateStatus(repoResponse.statusCode);
  }

  return {
    success: true,
    message: 'Token and repository validation successful',
  };

}

/**
 * Validate if a branch exists in the repository
 */
export const validateBranch = async (
  token: string,
  repoPath: string,
  branchName: string
): Promise<APIResponse> => {


  const [owner, repo] = repoPath.split('/');
  if (!owner || !repo) {
    return {
      success: false,
      message: 'Invalid repository path format. Use: owner/repository',
      error: 'Invalid repository path format. Use: owner/repository',
    };
  }

  const response = await httpUtils.get(
    `${baseUrl}/repos/${owner}/${repo}/branches/${branchName}`,
    defaultOptions(token)
  );

  if (!response.success) {
    return evaluateStatus(response.statusCode);
  }

  return {
    success: true,
    message: `Branch '${branchName}' exists in repository '${repoPath}'`,
  };

}