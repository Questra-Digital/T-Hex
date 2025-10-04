// GitHub API service for token and repository validation
import { httpUtils } from './httpUtils';
export interface GitHubTokenValidation {
  isValid: boolean;
  hasRepoAccess: boolean;
  hasWriteAccess: boolean;
  user?: {
    login: string;
    name?: string;
    email?: string;
  };
  error?: string;
}

export interface BranchValidation {
  exists: boolean;
  branchName: string;
  error?: string;
}

export interface RepositoryInfo {
  name: string;
  fullName: string;
  private: boolean;
  defaultBranch: string;
  permissions: {
    admin: boolean;
    push: boolean;
    pull: boolean;
  };
}

class GitHubApiService {
  private baseUrl = 'https://api.github.com';

  /**
   * Validate GitHub token and check repository access
   */
  async validateTokenAndRepo(
    token: string,
    repoPath: string
  ): Promise<GitHubTokenValidation> {
    try {
      // First, validate the token by getting user info
      const userResponse = await httpUtils.get(`${this.baseUrl}/user`, {
        headers: {
          'Authorization': `token ${token}`,
          'Accept': 'application/vnd.github.v3+json',
        },
      });

      const user = userResponse.data;

      // Parse repository path (format: owner/repo)
      const [owner, repo] = repoPath.split('/');
      if (!owner || !repo) {
        return {
          isValid: true,
          hasRepoAccess: false,
          hasWriteAccess: false,
          user: {
            login: user.login,
            name: user.name,
            email: user.email,
          },
          error: 'Invalid repository path format. Use: owner/repository',
        };
      }

      // Check repository access and permissions
      const repoResponse = await httpUtils.get(`${this.baseUrl}/repos/${owner}/${repo}`, {
        headers: {
          'Authorization': `token ${token}`,
          'Accept': 'application/vnd.github.v3+json',
        },
      });

      const repository: RepositoryInfo = repoResponse.data;

      return {
        isValid: true,
        hasRepoAccess: true,
        hasWriteAccess: repository.permissions.push,
        user: {
          login: user.login,
          name: user.name,
          email: user.email,
        },
      };
    } catch (error: any) {
      console.error('GitHub API validation error:', error);
      
      // Handle specific HTTP errors
      if (error.status === 401) {
        return {
          isValid: false,
          hasRepoAccess: false,
          hasWriteAccess: false,
          error: 'Invalid or expired GitHub token',
        };
      }
      
      if (error.status === 404) {
        return {
          isValid: true,
          hasRepoAccess: false,
          hasWriteAccess: false,
          user: error.data?.user,
          error: 'Repository not found or you do not have access to it',
        };
      }
      
      return {
        isValid: false,
        hasRepoAccess: false,
        hasWriteAccess: false,
        error: error.message || 'Network error or GitHub API unavailable',
      };
    }
  }

  /**
   * Validate if a branch exists in the repository
   */
  async validateBranch(
    token: string,
    repoPath: string,
    branchName: string
  ): Promise<BranchValidation> {
    try {
      const [owner, repo] = repoPath.split('/');
      if (!owner || !repo) {
        return {
          exists: false,
          branchName,
          error: 'Invalid repository path format',
        };
      }

      const response = await httpUtils.get(
        `${this.baseUrl}/repos/${owner}/${repo}/branches/${branchName}`,
        {
          headers: {
            'Authorization': `token ${token}`,
            'Accept': 'application/vnd.github.v3+json',
          },
        }
      );

      return {
        exists: true,
        branchName,
      };
    } catch (error: any) {
      console.error('Branch validation error:', error);
      
      if (error.status === 404) {
        return {
          exists: false,
          branchName,
          error: `Branch '${branchName}' does not exist in the repository`,
        };
      }
      
      return {
        exists: false,
        branchName,
        error: error.message || 'Network error while checking branch',
      };
    }
  }

  /**
   * Get repository information
   */
  async getRepositoryInfo(
    token: string,
    repoPath: string
  ): Promise<{ success: boolean; data?: RepositoryInfo; error?: string }> {
    try {
      const [owner, repo] = repoPath.split('/');
      if (!owner || !repo) {
        return {
          success: false,
          error: 'Invalid repository path format',
        };
      }

      const response = await httpUtils.get(`${this.baseUrl}/repos/${owner}/${repo}`, {
        headers: {
          'Authorization': `token ${token}`,
          'Accept': 'application/vnd.github.v3+json',
        },
      });

      const data: RepositoryInfo = response.data;
      return {
        success: true,
        data,
      };
    } catch (error: any) {
      console.error('Repository info error:', error);
      return {
        success: false,
        error: error.message || 'Network error while fetching repository information',
      };
    }
  }
}

export const githubApiService = new GitHubApiService();
