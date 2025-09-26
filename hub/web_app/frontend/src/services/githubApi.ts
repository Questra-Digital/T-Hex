// GitHub API service for token and repository validation
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
      const userResponse = await fetch(`${this.baseUrl}/user`, {
        headers: {
          'Authorization': `token ${token}`,
          'Accept': 'application/vnd.github.v3+json',
        },
      });

      if (!userResponse.ok) {
        if (userResponse.status === 401) {
          return {
            isValid: false,
            hasRepoAccess: false,
            hasWriteAccess: false,
            error: 'Invalid or expired GitHub token',
          };
        }
        return {
          isValid: false,
          hasRepoAccess: false,
          hasWriteAccess: false,
          error: 'Failed to validate GitHub token',
        };
      }

      const user = await userResponse.json();

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
      const repoResponse = await fetch(`${this.baseUrl}/repos/${owner}/${repo}`, {
        headers: {
          'Authorization': `token ${token}`,
          'Accept': 'application/vnd.github.v3+json',
        },
      });

      if (!repoResponse.ok) {
        if (repoResponse.status === 404) {
          return {
            isValid: true,
            hasRepoAccess: false,
            hasWriteAccess: false,
            user: {
              login: user.login,
              name: user.name,
              email: user.email,
            },
            error: 'Repository not found or you do not have access to it',
          };
        }
        return {
          isValid: true,
          hasRepoAccess: false,
          hasWriteAccess: false,
          user: {
            login: user.login,
            name: user.name,
            email: user.email,
          },
          error: 'Failed to access repository',
        };
      }

      const repository: RepositoryInfo = await repoResponse.json();

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
    } catch (error) {
      console.error('GitHub API validation error:', error);
      return {
        isValid: false,
        hasRepoAccess: false,
        hasWriteAccess: false,
        error: 'Network error or GitHub API unavailable',
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

      const response = await fetch(
        `${this.baseUrl}/repos/${owner}/${repo}/branches/${branchName}`,
        {
          headers: {
            'Authorization': `token ${token}`,
            'Accept': 'application/vnd.github.v3+json',
          },
        }
      );

      if (!response.ok) {
        if (response.status === 404) {
          return {
            exists: false,
            branchName,
            error: `Branch '${branchName}' does not exist in the repository`,
          };
        }
        return {
          exists: false,
          branchName,
          error: 'Failed to check branch existence',
        };
      }

      return {
        exists: true,
        branchName,
      };
    } catch (error) {
      console.error('Branch validation error:', error);
      return {
        exists: false,
        branchName,
        error: 'Network error while checking branch',
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

      const response = await fetch(`${this.baseUrl}/repos/${owner}/${repo}`, {
        headers: {
          'Authorization': `token ${token}`,
          'Accept': 'application/vnd.github.v3+json',
        },
      });

      if (!response.ok) {
        return {
          success: false,
          error: 'Failed to fetch repository information',
        };
      }

      const data: RepositoryInfo = await response.json();
      return {
        success: true,
        data,
      };
    } catch (error) {
      console.error('Repository info error:', error);
      return {
        success: false,
        error: 'Network error while fetching repository information',
      };
    }
  }
}

export const githubApiService = new GitHubApiService();
