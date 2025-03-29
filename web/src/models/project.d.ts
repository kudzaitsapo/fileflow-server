import { User } from "./user";

export interface Project {
  id: number;
  name: string;
  max_upload_size: number;
  description: string;
  created_at: string;
  project_key: string;
  allowed_file_types: string[];
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

export interface ProjectUser {
  id: number;
  project_id: number;
  user_info: User;
}
