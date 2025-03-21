export interface Project {
  id: number;
  name: string;
  max_file_size: number;
  description: string;
  created_at: string;
}

export interface IProjectFormProps {
  projectName: string;
  description: string;
  allowedFiles: string[];
  maxFileSize: number;
}

export interface IFormValidationErrors {
  projectName?: string;
  maxFileSize?: string;
}
