// Pipeline service for API interactions
import { httpUtils } from './httpUtils';

export interface CreatePipelineRequest {
  pipeline_id: string;
  repo_path: string;
  access_token: string;
}

export interface CreatePipelineResponse {
  success: boolean;
  message: string;
  error?: string;
}

class PipelineService {
  private baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:3000/api';

  /**
   * Create a new pipeline (simulated for now)
   */
  async createPipeline(request: CreatePipelineRequest): Promise<CreatePipelineResponse> {
    try {
      // Simulate API call for now
      console.log('Creating pipeline with:', request);
      
      // Simulate network delay
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Simulate success/failure based on some conditions
      const shouldSucceed = Math.random() > 0.2; // 80% success rate
      
      if (shouldSucceed) {
        return {
          success: true,
          message: 'Pipeline created successfully',
        };
      } else {
        return {
          success: false,
          message: 'Failed to create pipeline',
          error: 'Simulated server error - repository access denied',
        };
      }
    } catch (error: any) {
      console.error('Pipeline creation error:', error);
      return {
        success: false,
        message: 'Failed to create pipeline',
        error: error.message || 'Network error occurred',
      };
    }
  }
}

export const pipelineService = new PipelineService();
